package main

import (
	"bytes"
	"os/exec"
)

func run(program string, args []string) (outStr string, errStr string, err error) {
	cmd := exec.Command(program, args...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		return "", "", err
	}
	outStr, errStr = Convert(41, stdout.Bytes()), Convert(41, stderr.Bytes())
	return outStr, errStr, nil
}
