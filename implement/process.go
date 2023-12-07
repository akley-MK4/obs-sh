package implement

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func startProcess(kwArgs map[string]interface{}) (*exec.Cmd, string, error) {
	var arg []string
	for k, v := range kwArgs {
		arg = append(arg, fmt.Sprintf("-%s=%v", k, v))
	}

	cmd := exec.Command(outExecFilePath, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	//_, outputErr := cmd.Output()
	//if outputErr != nil {
	//	logger.WarnFmt("startSubprocess, %v", outputErr)
	//}

	argsStr := fmt.Sprintf("%s %s", outExecFilePath, strings.Join(arg, " "))
	if err := cmd.Start(); err != nil {
		return nil, argsStr, err
	}

	return cmd, argsStr, nil
}
