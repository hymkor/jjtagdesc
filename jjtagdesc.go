//go:build run

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func run(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func quote(args []string, f func(string) error) error {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	r, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	defer r.Close()
	cmd.Start()

	sc := bufio.NewScanner(r)
	for sc.Scan() {
		//println(sc.Text())
		if err := f(sc.Text()); err != nil {
			io.Copy(io.Discard, r)
			return err
		}
	}
	return nil
}

func mains() error {
	if dir, err := os.Stat(".git"); err == nil && dir.IsDir() {
		return run("git", "describe", "--tags")
	}
	nlines := 0
	ncommits := 0
	var current string
	result := "v0.0.0"
	err := quote([]string{"jj", "log", "--no-graph", "-r", "latest(tags()):: ~ description(exact:\"\")"},
		func(line string) error {
			nlines++
			if nlines%2 != 0 {
				fields := strings.Fields(line)
				if ncommits == 0 {
					current = fields[0]
				}
				if len(fields) == 7 {
					if ncommits == 0 {
						result = fields[5]
					} else {
						result = fmt.Sprintf("%s-%d-%s", fields[5], ncommits, current)
					}
				}
				ncommits++
			}
			return nil
		},
	)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}

func main() {
	if err := mains(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
