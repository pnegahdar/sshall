package main

import (
	"bufio"
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"sync"
	"strings"
)

var flagConcurrency = cli.IntFlag{
	Name:  "concurrency",
	Value : 2,
	Usage: "The number of concurrent ssh requests"}

var flagCmd = cli.StringFlag{
	Name:  "cmd",
	Usage: "The command to run"}


var flagUsers = cli.StringSliceFlag{
	Name:  "try-user",
	Value: &DefaultUsers,
	Usage: "The command to run"}

func streamMachines(wg *sync.WaitGroup, streamTo chan string) {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			text := scanner.Text()
			if strings.TrimSpace(text) == "" {
				continue
			}
			wg.Add(1)
			streamTo <- text
		}
	} else {
		fmt.Println("Please pipe in ips list to stdin.")
		os.Exit(1)
	}
}

func printState(header string, parts ...string) {
	fmt.Printf("%v\t%v\n", header, strings.Join(parts, ", "))
}

func run(c *cli.Context) {
	cmd := c.String(flagCmd.Name)
	if cmd == "" {
		fmt.Println("A cmd is required.")
		os.Exit(1)
	}
	concurrency := c.Int(flagConcurrency.Name)
	if concurrency <= 1 {
		concurrency = 1
	}
	potentialUsers := c.StringSlice(flagUsers.Name)
	mCh := make(chan string)
	wg := &sync.WaitGroup{}
	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()
			for machineString := range mCh {
				machine, err := NewMachineFromString(machineString, potentialUsers...)
				if err != nil {
					printState(machineString, err.Error())
					continue
				}
				err = machine.ExecCmd(cmd)
				if err != nil {
					printState(machineString, err.Error())
				}
			}
		}()
	}
	streamMachines(wg, mCh)
	close(mCh)
	wg.Wait()

}

func main() {
	app := cli.NewApp()
	app.Name = "sshall"
	app.Usage = "Run ssh commands on a cluster."
	app.Version = "0.1"
	app.Author = "Parham Negahdar <pnegahdar@gmail.com>"
	app.Flags = []cli.Flag{flagConcurrency, flagCmd, flagUsers}
	app.Action = run
	app.Run(os.Args)
}
