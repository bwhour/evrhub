package ethContract

import (
	"fmt"
	"os/exec"
	"strings"
)

const (
	SolcCmdText   = "[SOLC_CMD]"
	DirectoryText = "[DiRECTORY]"
	ContractText  = "[CONTRACT]"
)

var (
	// BaseABIBINGEnCmd is the base command for ethContract compilation to ABI and BIN
	BaseABIBINGenCmd = strings.Join([]string{"solc ",
		fmt.Sprintf("--%s ./eth_contracts/contracts/%s%s.sol ", SolcCmdText, DirectoryText, ContractText),
		fmt.Sprintf("-o ./cmd/evrrelayer/ethContract/generated/%s/%s ", SolcCmdText, ContractText),
		"--overwrite ",
		"--allow-paths *,"},
		"")
	// BaseBindingGenCmd is the base command for ethContract binding generation
	BaseBindingGenCmd = strings.Join([]string{"abigen ",
		fmt.Sprintf("--bin ./cmd/evrrelayer/ethContract/generated/bin/%s/%s.bin ", ContractText, ContractText),
		fmt.Sprintf("--abi ./cmd/evrrelayer/ethContract/generated/abi/%s/%s.abi ", ContractText, ContractText),
		fmt.Sprintf("--pkg %s ", ContractText),
		fmt.Sprintf("--type %s ", ContractText),
		fmt.Sprintf("--out ./cmd/evrrelayer/ethContract/generated/bindings/%s/%s.go", ContractText, ContractText)},
		"")
)

// CompileContracts compiles contracts to BIN and ABI files
func CompileContracts(contracts BridgeContracts) error {
	for _, ethContract := range contracts {
		// Construct generic BIN/ABI generation cmd with ethContract's directory path and name
		baseDirectory := ""
		if ethContract.String() == BridgeBank.String() {
			baseDirectory = ethContract.String() + "/"
		}
		dirABIBINGenCmd := strings.Replace(BaseABIBINGenCmd, DirectoryText, baseDirectory, -1)
		contractABIBINGenCmd := strings.Replace(dirABIBINGenCmd, ContractText, ethContract.String(), -1)

		// Segment BIN and ABI generation cmds
		contractBINGenCmd := strings.Replace(contractABIBINGenCmd, SolcCmdText, "bin", -1)
		err := execCmd(contractBINGenCmd)
		if err != nil {
			return err
		}

		contractABIGenCmd := strings.Replace(contractABIBINGenCmd, SolcCmdText, "abi", -1)
		err = execCmd(contractABIGenCmd)
		if err != nil {
			return err
		}
	}
	return nil
}

// GenerateBindings generates bindings for each ethContract
func GenerateBindings(contracts BridgeContracts) error {
	for _, ethContract := range contracts {
		genBindingCmd := strings.Replace(BaseBindingGenCmd, ContractText, ethContract.String(), -1)
		err := execCmd(genBindingCmd)
		if err != nil {
			return err
		}
	}
	return nil
}

// execCmd executes a bash cmd
func execCmd(cmd string) error {
	_, err := exec.Command("sh", "-c", cmd).Output()
	return err
}
