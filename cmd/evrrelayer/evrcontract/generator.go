package evrcontract

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
	// BaseABIBINGEnCmd is the base command for evrcontract compilation to ABI and BIN
	BaseABIBINGenCmd = strings.Join([]string{"solc ",
		fmt.Sprintf("--%s ./evr_contracts/contracts/%s%s.sol ", SolcCmdText, DirectoryText, ContractText),
		fmt.Sprintf("-o ./cmd/evrrelayer/evrcontract/generated/%s/%s ", SolcCmdText, ContractText),
		"--overwrite ",
		"--allow-paths *,"},
		"")
	// BaseBindingGenCmd is the base command for evrcontract binding generation
	BaseBindingGenCmd = strings.Join([]string{"abigen ",
		fmt.Sprintf("--bin ./cmd/evrrelayer/evrcontract/generated/bin/%s/%s.bin ", ContractText, ContractText),
		fmt.Sprintf("--abi ./cmd/evrrelayer/evrcontract/generated/abi/%s/%s.abi ", ContractText, ContractText),
		fmt.Sprintf("--pkg %s ", ContractText),
		fmt.Sprintf("--type %s ", ContractText),
		fmt.Sprintf("--out ./cmd/evrrelayer/evrcontract/generated/bindings/%s/%s.go", ContractText, ContractText)},
		"")
)

// CompileContracts compiles contracts to BIN and ABI files
func CompileContracts(contracts BridgeContracts) error {
	for _, evrcontract := range contracts {
		// Construct generic BIN/ABI generation cmd with evrcontract's directory path and name
		baseDirectory := ""
		if evrcontract.String() == BridgeBank.String() {
			baseDirectory = evrcontract.String() + "/"
		}
		dirABIBINGenCmd := strings.Replace(BaseABIBINGenCmd, DirectoryText, baseDirectory, -1)
		contractABIBINGenCmd := strings.Replace(dirABIBINGenCmd, ContractText, evrcontract.String(), -1)

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

// GenerateBindings generates bindings for each evrcontract
func GenerateBindings(contracts BridgeContracts) error {
	for _, evrcontract := range contracts {
		genBindingCmd := strings.Replace(BaseBindingGenCmd, ContractText, evrcontract.String(), -1)
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
	fmt.Println("CMD: "+cmd)
	//mycmd := exec.Command("sh","-c",cmd)
	//var out bytes.Buffer
	//mycmd.Stdout = &out
	//err := mycmd.Run()
	//if err != nil {
	//	fmt.Println(err)
	//	return err
	//}
	//fmt.Println("Result: " + out.String())
	return err
}
