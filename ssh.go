package main

import (
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"os"
	"net"
	"log"
	"strings"
)

func executeCmd(cmd string, machine *Machine) {
	for i, user := range machine.PotentialUsers {
		sock, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
		if err != nil {
			log.Fatal(err)
		}
		agent := agent.NewClient(sock)
		signers, err := agent.Signers()
		if err != nil {
			log.Fatal(err)
		}
		auths := []ssh.AuthMethod{ssh.PublicKeys(signers...)}
		cfg := &ssh.ClientConfig{
			User: user,
			Auth: auths,
		}
		cfg.SetDefaults()
		client, errConnect := ssh.Dial("tcp", machine.DialAddr(), cfg)
		if errConnect != nil {
			authError := strings.Contains(errConnect.Error(), "unable to authenticate")
			if authError && i <= (len(machine.PotentialUsers) - 1) {
				continue
			}
			printState(machine.String(), errConnect.Error())
			return
		}
		session, errSession := client.NewSession()
		if errSession != nil {
			printState(machine.String(), errSession.Error())
			return
		}
		errRun := session.Run(cmd)
		if errRun != nil {
			printState(machine.String(), errRun.Error())
			return
		}
	}
	printState(machine.String(), "SUCCESS")
	return
}
