package ooo

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func RunProg(argus ...string) (int, string) {
	var cmd *exec.Cmd
	oldstdout := os.Stdout

	if len(argus) == 1 {
		cmd = exec.Command(argus[0])
	}
	if len(argus) > 1 {
		x := argus[0]   // get the 0 index element from slice
		aa := argus[1:] // remove the 0 index element from slice

		cmd = exec.Command(x, aa...)

	}

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)

	err := cmd.Run()
	if err != nil {
		fmt.Printf("caller err:\n%s\n", stderrBuf.String())
		return -1, err.Error()
	}

	outStr, _ := stdoutBuf.String(), stderrBuf.String()
	fmt.Println("--->", outStr)
	os.Stdout = oldstdout
	return 0, ""
}
