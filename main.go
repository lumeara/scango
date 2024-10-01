package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

// TODO: get **Scanner** class and methods out of main package
type Scanner struct {
	targetIP    string
	targetPorts []int
}

// TODO: make this run in parallel
func (s Scanner) RunScan(t string) {
	// TODO: abstract this away to multiple scan types
	var targetURL string

	for _, port := range s.targetPorts {
		targetURL = s.targetIP + ":" + string(port)

		conn, err := net.Dial("tcp", targetURL)
		// if err != nil {
		// 	log.Fatal(err)
		// }
	}
}

func handleArgs(args []string) (*Scanner, string, error) {
	if len(args) == 1 {
		return &Scanner{}, "", fmt.Errorf("Usage: scango <target-ip> port1,port1,port3")
	}

	scanner := new(Scanner)

	// TODO: validate targetIP argument
	scanner.targetIP = args[1]

	// TODO: validate targetPorts argument
	scanner.targetPorts = []int{21, 80}

	// TODO: validate scanType argument
	scanType := "syn"

	return scanner, scanType, nil
}

func main() {
	args := os.Args

	scanner, scanType, err := handleArgs(args)
	if err != nil {
		log.Fatal(err)
	}

	scanner.RunScan(scanType)
}
