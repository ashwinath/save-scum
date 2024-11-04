package shell

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
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
	// Check if save scum file is there
	if checkIfTransferInProgress(to) {
		return nil, fmt.Errorf("transfer in progress, skipping")
	}

	timeNow := time.Now().Format("2006-01-02-15-04-05")
	statusFileName := fmt.Sprintf("%s/save-scum-%s", to, timeNow)
	err := os.WriteFile(statusFileName, []byte(""), 0755)
	if err != nil {
		fmt.Printf("unable to write status file: %s", err)
	}

	defer os.Remove(statusFileName)

	var args []string
	args = append(args, flags...)
	args = append(args, from)
	args = append(args, to)
	return runShell("rsync", args...)
}

func checkIfTransferInProgress(to string) bool {
	files, _ := os.ReadDir(to)

	for _, file := range files {
		if strings.HasPrefix(file.Name(), "save-scum-") {
			return true
		}
	}
	return false
}

func ChownRecursive(path string, user string, group string) (*string, error) {
	return runShell("chown", "-R", fmt.Sprintf("%s:%s", user, group), path)
}
