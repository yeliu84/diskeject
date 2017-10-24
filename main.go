package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/DHowett/go-plist"
)

type diskList struct {
	WholeDisks []string `plist:"WholeDisks"`
}

func errprint(msg string) {
	fmt.Fprintf(os.Stderr, "%s", msg)
}

func errexit(msg string) {
	errprint(msg)
	os.Exit(1)
}

func getDisks() ([]string, error) {
	out, err := exec.Command("diskutil", "list", "-plist", "external").CombinedOutput()
	if err != nil {
		return nil, err
	}
	var disks diskList
	err = plist.NewDecoder(bytes.NewReader(out)).Decode(&disks)
	if err != nil {
		return nil, err
	}
	return disks.WholeDisks, nil
}

func eject(disk string) {
	out, err := exec.Command("diskutil", "eject", disk).CombinedOutput()
	if err != nil {
		errexit(string(out))
	}
}

func main() {
	disks, err := getDisks()
	if err != nil {
		errexit("get disks error")
	}
	for i := len(disks) - 1; i >= 0; i-- {
		eject(disks[i])
	}
}
