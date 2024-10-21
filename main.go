package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    for {
        input, err := reader.ReadString('\n')
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
        }

        if err = execInput(input); err != nil {
            fmt.Fprintln(os.Stderr, err)
        }

        fmt.Print(input)
    }
}

func execInput(input string) error {
    input = strings.TrimSuffix(input, "\n")
    cmd := exec.Command(input)
    cmd.Stderr = os.Stderr
    cmd.Stdout = os.Stdout
    return cmd.Run()
}
