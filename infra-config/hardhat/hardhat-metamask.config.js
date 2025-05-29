/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
    solidity: "0.8.20",
    networks: {
        hardhat: {
            gas: 0,
            accounts: {
                mnemonic: process.env.SEED_PHRASE,
            },
            chainId: 1337,
        },
    },
};
