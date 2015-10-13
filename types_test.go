package main

import "testing"
import "github.com/stretchr/testify/assert"

var expectedResults = map[string]*Machine{
	"10.0.0.1":           &Machine{HostIP: "10.0.0.1", Port: "22", PotentialUsers: []string{"root"}},
	"10.0.0.1:23":        &Machine{HostIP: "10.0.0.1", Port: "23", PotentialUsers: []string{"root"}},
	"ubuntu@10.0.0.1:23": &Machine{HostIP: "10.0.0.1", Port: "23", PotentialUsers: []string{"ubuntu"}},
	"ubuntu@10.0.0.1":    &Machine{HostIP: "10.0.0.1", Port: "22", PotentialUsers: []string{"ubuntu"}},
}

func TestNameToMachine(t *testing.T) {
	for stringToTest, expectedResult := range expectedResults {
		m, err := NewMachineFromString(stringToTest)
		assert.Nil(t, err)
		assert.Equal(t, m, expectedResult)
	}
}
