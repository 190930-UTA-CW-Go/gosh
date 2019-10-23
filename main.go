package main

import (
	"bufio"
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("ssh", "user1@192.168.56.102")
	reader, _ := cmd.StdoutPipe()
	scanner := bufio.NewScanner(reader)
	go func() {
		for scanner.Scan() {
			fmt.Println("user1> ", scanner.Text())
		}
	}()
	cmd.Start()
	cmd.Wait()
}
