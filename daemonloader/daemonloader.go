package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

const (
	PROGRAM = "prog"
	DAEMON  = "daemon"
	FOREVER = "forever"
)

func StripSlice(slice []string, element string) []string {
	for i := 0; i < len(slice); {
		if slice[i] == element && i != len(slice)-1 {
			slice = append(slice[:i], slice[i+1:]...)
			break
		} else if slice[i] == element && i == len(slice)-1 {
			slice = slice[:i]
			break
		} else {
			i++
		}
	}
	return slice
}

func SubProcess(args []string, shell bool) *exec.Cmd {
	var cmd *exec.Cmd
	if shell {
		switch runtime.GOOS {
		case "darwin":
			cmd = exec.Command(os.Getenv("SHELL"), "-c", args[0])
		case "linux":
			cmd = exec.Command(os.Getenv("SHELL"), "-c", args[0])
		case "windows":
			cmd = exec.Command("cmd", "/C", args[0])
		default:
			os.Exit(1)
		}
	} else {
		cmd = exec.Command(args[0], args[1:]...)
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[-] Error: %s\n", err)
	}
	return cmd
}

func main() {
	var cmd *exec.Cmd
	program := flag.String(PROGRAM, "", "run program")
	daemon := flag.Bool(DAEMON, false, "run in daemon")
	forever := flag.Bool(FOREVER, false, "run forever")
	flag.Parse()
	fmt.Printf("[*] PID: %d PPID: %d ARG: %s PROG:\"%s\"\n", os.Getpid(), os.Getppid(), os.Args, *program)
	if *program == "" {
		flag.Usage()
		os.Exit(1)
	}
	if *daemon {
		fmt.Printf("[*] Daemon running in PID: %d PPID: %d\n", os.Getpid(), os.Getppid())
		SubProcess(StripSlice(os.Args, "-"+DAEMON), false)
		os.Exit(0)
	} else if *forever {
		args := os.Args
		for {
			fmt.Printf("[*] Forever running in PID: %d PPID: %d\n", os.Getpid(), os.Getppid())
			cmd = SubProcess(StripSlice(args, "-"+FOREVER), false)
			cmd.Wait()
		}
	} else {
		fmt.Printf("[*] Service running in PID: %d PPID: %d\n", os.Getpid(), os.Getppid())
		cmd = SubProcess([]string{*program}, true)
		cmd.Wait()
	}
}
