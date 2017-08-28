// package never-fail is a cli command wrapper to run a command and ignore exit code
package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) == 2 && (os.Args[1] == "--help" || os.Args[1] == "-h") {
		showHelp()
		return
	} else if len(os.Args) < 2 {
		showHelp()
		return
	}

	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println()
		fmt.Println("An error has occured: ", err)
	}
}

func showHelp() {
	fmt.Println("never-fail - head")
	fmt.Println("")
	fmt.Println("USAGE")
	fmt.Println(" never-fail <bin> <args>")
	fmt.Println(" never-fail --help|-h")
}
