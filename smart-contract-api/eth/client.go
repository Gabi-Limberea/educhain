package eth

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
	"smart-contract-api/contract"
	"smart-contract-api/ipfs"
	"smart-contract-api/models"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	conn         *ethclient.Client
	masterKey    *ecdsa.PrivateKey
	masterWallet *common.Address
}

func NewClient() *Client {
	netUrl := os.Getenv("NET_URL")
	if netUrl == "" {
		log.Fatal("NET_URL environment variable not set")
	}

	ethClient, err := ethclient.Dial(netUrl)
	if err != nil {
		log.Fatalf("could not connect to ethclient: %v", err)
	}

	var masterWallet common.Address
	var masterKey *ecdsa.PrivateKey

	masterKeyFile := os.Getenv("MASTER_KEY")
	if masterKeyFile == "" {
		masterKey, err = crypto.GenerateKey()
		if err != nil {
			log.Fatalf("could not generate master private key: %v", err)
		}

		slog.Info("generated master private key")
	} else {
		masterKeyPasswd := os.Getenv("MASTER_KEY_PASSWD")
		if masterKeyPasswd == "" {
			log.Fatal("MASTER_KEY_PASSWD environment variable not set")
		}

		f, err := os.ReadFile(masterKeyFile)
		if err != nil {
			log.Fatalf("could not read master private key file: %v", err)
		}

		key, err := keystore.DecryptKey(f, masterKeyPasswd)
		if err != nil {
			log.Fatalf("could not decrypt master private key: %v", err)
		}

		masterKey = key.PrivateKey
		slog.Info("master private key provided through file")
	}

	pubKeyECDSA, ok := masterKey.Public().(*ecdsa.PublicKey)
	if !ok {
		log.Fatalf("could not assert public key: %v", err)
	}

	masterWallet = crypto.PubkeyToAddress(*pubKeyECDSA)
	slog.Info("imported master wallet", "address", masterWallet.Hex())

	return &Client{
		conn:         ethClient,
		masterKey:    masterKey,
		masterWallet: &masterWallet,
	}
}

func (c *Client) Close() {
	c.conn.Close()
}

// NewContract is a function to generate and deploy a diploma smart-contract
func (c *Client) NewContract() (string, error) {
	chainID, err := c.conn.ChainID(context.Background())
	if err != nil {
		return "", fmt.Errorf("could not get chain ID: %w", err)
	}

	auth := bind.NewKeyedTransactor(c.masterKey, chainID)
	deployer := bind.DefaultDeployer(auth, c.conn)
	deployParams := &bind.DeploymentParams{
		Contracts: []*bind.MetaData{&contract.DiplomaMetaData},
	}

	deployRes, err := bind.LinkAndDeploy(deployParams, deployer)
	if err != nil {
		return "", fmt.Errorf("could not submit contract: %w", err)
	}

	address, tx := deployRes.Addresses[contract.DiplomaMetaData.ID], deployRes.Txs[contract.DiplomaMetaData.ID]
	slog.Info("contract pending deploy", "address", address.Hex())
	slog.Info("transaction pending", "hash", tx.Hash().Hex())

	if address, err = bind.WaitDeployed(context.Background(), c.conn, tx.Hash()); err != nil {
		return "", fmt.Errorf("could not deploy contract: %w", err)
	}

	slog.Info("master contract deployed", "address", address.Hex())
	return address.Hex(), nil
}

// GenerateNFTMetadataFile generates metadata for the given NFT based on the provider information
func GenerateNFTMetadataFile(file *ipfs.File, provider models.Provider) (io.ReadCloser, error) {
	metadata := map[string]string{
		"name":        file.Name,
		"image":       fmt.Sprintf("ipfs://%s", file.Hash),
		"description": fmt.Sprintf("Diploma provided by %s", provider.OrganizationInfo.Name),
	}

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(metadata); err != nil {
		return nil, fmt.Errorf("failed to encode metadata: %w", err)
	}

	return io.NopCloser(bytes.NewReader(buf.Bytes())), nil
}

// MintNFT mints a new NFT with the given tokenURI on the contractAddr to the destAddr
func (c *Client) MintNFT(tokenURI, contractAddr, destAddr string) (*int64, error) {
	diplomaContract := contract.NewDiploma()
	instance := diplomaContract.Instance(c.conn, common.HexToAddress(contractAddr))
	chainID, err := c.conn.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("could not get chain ID: %w", err)
	}

	auth := bind.NewKeyedTransactor(c.masterKey, chainID)
	tx, err := bind.Transact(
		instance, auth, diplomaContract.PackSafeMint(common.HexToAddress(destAddr), tokenURI),
	)
	if err != nil {
		return nil, fmt.Errorf("could not transact: %w", err)
	}

	if _, err = bind.WaitMined(context.Background(), c.conn, tx.Hash()); err != nil {
		return nil, fmt.Errorf("could not deploy contract: %w", err)
	}

	res, err := bind.Call(
		instance, nil, diplomaContract.PackTokenCounter(), diplomaContract.UnpackTokenCounter,
	)
	if err != nil {
		return nil, fmt.Errorf("could not unpack token counter: %w", err)
	}

	tokenID := res.Int64()
	return &tokenID, nil
}
