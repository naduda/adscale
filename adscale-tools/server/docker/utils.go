package docker

import (
	"bufio"
	"fmt"
	"os/exec"
	"runtime"
)

func executeCommand(command string) (string, error) {
	name := "/bin/sh"
	arg := "-c"
	if runtime.GOOS == "windows" {
		name = "cmd"
		arg = "/C"
	}

	out, err := exec.Command(name, arg, command).Output()
	if err != nil {
		return "", err
	}

	return string(out[:]), nil
}

func RunCommand(command string, dir string) error {
	name := "/bin/sh"
	arg := "-c"
	if runtime.GOOS == "windows" {
		name = "cmd"
		arg = "/C"
	}
	cmd := exec.Command(name, arg, command)
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
