package main

import (
	"bufio"
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"sync"
)

var FlagIdentityFile = cli.StringSliceFlag{
	Name:  "identity-file",
	Usage: "The identity file to use."}

var FlagConcurrency = cli.IntFlag{
	Name:  "concurrency",
	Usage: "The number of concurrent ssh requests"}

var FlagCmd = cli.IntFlag{
	Name:  "cmd",
	Usage: "The command to run"}

func streamMachines(streamTo chan string) {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			streamTo <- scanner.Text()
		}
	} else {
		fmt.Println("Please pipe in ips list to stdin.")
		os.Exit(1)
	}
}

func PrintState(ipString string, state string) {
	fmt.Printf("%v -- %v\n", ipString, state)
}

func run(c *cli.Context) {
	identityFiles := c.StringSlice(FlagIdentityFile.Name)
	if len(identityFiles) == 0 {
		fmt.Println("Atleast one identity file is required to ssh in")
		os.Exit(1)
	}
	cmd := c.String(FlagCmd.Name)
	if cmd == "" {
		fmt.Println("A cmd is required.")
		os.Exit(1)
	}
	concurrency := c.Int(FlagIdentityFile.Name)
	if concurrency == 0 {
		concurrency = 1
	}
	mCh := make(chan string)
	wg := &sync.WaitGroup{}
	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()
			for machineString := range mCh {
				machine, err := NewMachineFromString(machineString)
				if err != nil {
					PrintState(machineString, err.Error())
					continue
				}
				err = machine.ExecCmd(cmd, identityFiles...)
				if err != nil {
					PrintState(machineString, err.Error())
				}
			}
		}()
	}
	streamMachines(mCh)
	close(mCh)
	wg.Wait()

}

func main() {
	app := cli.NewApp()
	app.Name = "sshall"
	app.Usage = "Run ssh commands on a cluster."
	app.Version = "0.1"
	app.Author = "Parham Negahdar <pnegahdar@gmail.com>"
	app.Flags = []cli.Flag{FlagIdentityFile, FlagConcurrency}
	app.Action = run
	app.Run(os.Args)
}
