package pip

import (
	"os/exec"
	"bytes"
)

func RunPipCommand(args ...string) (string, error) {
	cmd := exec.Command("pip", args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return stderr.String(), err
	}
	return out.String(), nil
}