CONTRACT_PKG:=..
CONTRACT_DIR:=$(CONTRACT_PKG)/solidity
ARTIFACTS_DIR:=$(CONTRACT_DIR)/artifacts

contract-gen:
	solc --combined-json abi,bin $(CONTRACT_DIR)/diploma.sol > $(ARTIFACTS_DIR)/diploma.abi
	abigen --v2 --combined-json=$(ARTIFACTS_DIR)/diploma.abi --pkg=contract --type=diploma --out=$(CONTRACT_PKG)/diploma.go