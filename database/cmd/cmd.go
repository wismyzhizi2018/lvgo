package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// cmd := exec.Command("notepad")
	// cmd.Run()
	cmd := exec.Command("go")
	out, _ := cmd.CombinedOutput()
	fmt.Printf("%s", out)
}
