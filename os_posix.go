// +build !windows

package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

const defaultSocketType = "unix"

var unixStopSignals = []os.Signal{syscall.SIGTERM}

// Full path of the current executable
func get_executable_filename() string {
	// try readlink first
	path, err := os.Readlink("/proc/self/exe")
	if err == nil {
		return path
	}
	// use argv[0]
	path = os.Args[0]
	if !filepath.IsAbs(path) {
		cwd, _ := os.Getwd()
		path = filepath.Join(cwd, path)
	}
	if fileExists(path) {
		return path
	}
	// Fallback : use "gocode" and assume we are in the PATH...
	path, err = exec.LookPath("gocode")
	if err == nil {
		return path
	}
	return ""
}

// Additional OS-specific signals for process termination
func getProcessStopSignals() []os.Signal {
	return unixStopSignals
}
