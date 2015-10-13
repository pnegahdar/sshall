package main

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"github.com/codegangsta/cli"
)

var DefaultUsers = cli.StringSlice{"ubuntu"}
var DefaultPort = "22"

func NewMachineFromString(ipString string, possibleUsers ...string) (*Machine, error) {
	machine := &Machine{}
	ipString = strings.TrimSpace(ipString)
	splitUser := strings.Split(ipString, "@")
	machine.HostIP = splitUser[0]
	machine.PotentialUsers = possibleUsers
	if len(splitUser) == 2 {
		machine.HostIP = splitUser[1]
		machine.PotentialUsers = []string{splitUser[0]}
	}
	splitColon := strings.Split(machine.HostIP, ":")
	machine.HostIP = splitColon[0]
	machine.Port = DefaultPort
	if len(splitColon) == 2 {
		machine.Port = splitColon[1]
	}
	if machine.HostIP == "" || machine.Port == "" || len(machine.PotentialUsers) == 0 {
		return nil, errors.New("Unable to parse this ip")
	}
	return machine, nil
}

type Machine struct {
	sync.Mutex
	HostIP         string
	Port           string
	PotentialUsers []string
}

func (m *Machine) String() string {
	return m.DialAddr()
}

func (m *Machine) DialAddr() string {
	return fmt.Sprintf("%v:%v", m.HostIP, m.Port)
}

func (m *Machine) ExecCmd(cmd string) error {
	m.Lock()
	defer m.Unlock()
	executeCmd(cmd, m)
	return nil
}
