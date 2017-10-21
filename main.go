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

type diskInfo struct {
	BusProtocol string `plist:"BusProtocol"`
}

func errprint(msg string) {
	fmt.Fprintf(os.Stderr, "%s", msg)
}

func errexit(msg string) {
	errprint(msg)
	os.Exit(1)
}

func getDisks(diskType string) ([]string, error) {
	out, err := exec.Command("diskutil", "list", "-plist", "external", diskType).CombinedOutput()
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

func isDiskImage(disk string) (bool, error) {
	out, err := exec.Command("diskutil", "info", "-plist", disk).CombinedOutput()
	if err != nil {
		return false, err
	}
	var info diskInfo
	err = plist.NewDecoder(bytes.NewReader(out)).Decode(&info)
	if err != nil {
		return false, err
	}
	return info.BusProtocol == "Disk Image", nil
}

func eject(disk string) {
	out, err := exec.Command("diskutil", "eject", disk).CombinedOutput()
	if err != nil {
		errexit(string(out))
	}
}

func ejectDiskImages() {
	disks, err := getDisks("virtual")
	if err != nil {
		errexit("get disks error")
	}
	for _, disk := range disks {
		isImage, err := isDiskImage(disk)
		if err != nil {
			errexit("check disk image error")
		}
		if isImage {
			eject(disk)
		}
	}
}

func ejectPhysicalDisks() {
	disks, err := getDisks("physical")
	if err != nil {
		errexit("get disks error")
	}
	for _, disk := range disks {
		eject(disk)
	}
}

func main() {
	ejectDiskImages()
	ejectPhysicalDisks()
}
