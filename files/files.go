package main

import (
	"io/ioutil"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls")
	output, err := cmd.Output()
	writestring := []byte(output)
	err = ioutil.WriteFile("temp.txt", writestring, 777)
	if err != nil {
		panic(err)
	}
}
