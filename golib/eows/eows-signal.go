// +build !windows

// Package eows is used to Execute commands Over WebSocket
package eows

import (
	"fmt"
	"os"
	"syscall"
)

// Signal sends a signal to the running command / process
func (e *ExecOverWS) Signal(signal string) error {
	var sig os.Signal
	switch signal {
	case "quit", "SIGQUIT":
		sig = syscall.SIGQUIT
	case "terminated", "SIGTERM":
		sig = syscall.SIGTERM
	case "interrupt", "SIGINT":
		sig = syscall.SIGINT
	case "aborted", "SIGABRT":
		sig = syscall.SIGABRT
	case "continued", "SIGCONT":
		sig = syscall.SIGCONT
	case "hangup", "SIGHUP":
		sig = syscall.SIGHUP
	case "killed", "SIGKILL":
		sig = syscall.SIGKILL
	case "stopped (signal)", "SIGSTOP":
		sig = syscall.SIGSTOP
	case "stopped", "SIGTSTP":
		sig = syscall.SIGTSTP
	case "user defined signal 1", "SIGUSR1":
		sig = syscall.SIGUSR1
	case "user defined signal 2", "SIGUSR2":
		sig = syscall.SIGUSR2
	default:
		return fmt.Errorf("Unsupported signal")
	}

	if e.proc == nil {
		return fmt.Errorf("Cannot retrieve process")
	}

	e.logDebug("SEND signal %v to proc %v", sig, e.proc.Pid)
	return e.proc.Signal(sig)
}
