package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	flag "github.com/spf13/pflag"
)

func main() {
	var (
		helpMode    bool
		versionMode bool
		timeout     int
		cross       bool
		exclude     string
		from        string
		version     = "1.0.1"
	)

	flag.BoolVarP(&versionMode, "version", "v", false, "Print help message and quit.")
	flag.BoolVarP(&helpMode, "help", "h", false, "Print version and quit.")
	flag.IntVarP(&timeout, "timeout", "t", 30, "Timeout for ncdu.")
	flag.BoolVarP(&cross, "cross", "x", false, "Do not cross filesystem boundaries, i.e. only count files and directories on the same filesystem as the directory being scanned.")
	flag.StringVarP(&exclude, "exclude", "e", "", "Exclude files that match PATTERN.")
	flag.StringVarP(&from, "exclude-from", "X", "", "Exclude files that match any pattern in FILE. Patterns should be separated by a newline.")

	flag.CommandLine.SortFlags = false
	flag.Parse()

	if helpMode {
		fmt.Println("Usage: ncdu-wrapper dir [options]")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if versionMode {
		fmt.Printf("Version: %s\n", version)
		os.Exit(0)
	}

	// TODO: Automatic installation
	if !isExist() {
		fmt.Println("ncdu not exsit.")
		os.Exit(1)
	}

	// common command
	cmd := "ncdu -o-"

	// specify directory
	if len(os.Args) == 2 && os.Args[1] != "" {
		cmd = fmt.Sprintf("ncdu %s -o-", os.Args[1])
	}

	// support cross
	if cross {
		cmd = cmd + "-x"
	}

	// support exclude
	if exclude != "" {
		cmd = cmd + fmt.Sprintf("--exclude %s", exclude)
	}

	// support read exclude from file
	if from != "" {
		cmd = cmd + fmt.Sprintf("--exclude-from %s", from)
	}

	doneSignal := make(chan struct{})

	go func(ch chan struct{}) {
		progressSigns := []string{"|", "/", "-", "\\", "|"}
		for {
			for _, p := range progressSigns {
				fmt.Print("\rProgress: ", p)
				os.Stdout.Sync()
				select {
				case <-ch:
					return
				case <-time.After(time.Millisecond * 50):
					continue
				}
			}
		}
	}(doneSignal)

	Wrapper(cmd, timeout)
	doneSignal <- struct{}{}
	fmt.Print("\rProgress: 100%")
	os.Stdout.Sync()
	fmt.Println()
	result.OutPut()
}

func isExist() bool {
	cmd := exec.Command("which", "ncdu")
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}
