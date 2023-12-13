package shell

import (
	"fmt"
	"os/exec"
)

func runShell(binary string, args ...string) (*string, error) {
	cmd := exec.Command(binary, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("could not run command: %v, binary message: %s", err, string(out))
	}
	output := string(out)
	return &output, nil
}

func Rsync(flags []string, from string, to string) (*string, error) {
	var args []string
	for _, flag := range flags {
		args = append(args, flag)
	}
	args = append(args, from)
	args = append(args, to)
	return runShell("rsync", args...)
}

func ChownRecursive(path string, user string, group string) (*string, error) {
	return runShell("chown", "-R", fmt.Sprintf("%s:%s", user, group), path)
}
