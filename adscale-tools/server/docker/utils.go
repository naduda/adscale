package docker

import (
	"bufio"
	"fmt"
	"os/exec"
)

func executeCommand(command string) (string, error) {
	//out, err := exec.Command("/bin/sh", "-c", "docker ps --format \"{{.Names}}\"").Output()
	out, err := exec.Command("/bin/sh", "-c", command).Output()
	if err != nil {
		return "", err
	}

	return string(out[:]), nil
}

func RunCommand(command string, dir string) error {
	cmd := exec.Command("/bin/sh", "-c", command)
	cmd.Dir = dir

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}

	return cmd.Wait()
}
