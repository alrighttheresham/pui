package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/fatih/color"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

var Commands = []cli.Command{
	commandCheck,
	commandInit,
}

var commandCheck = cli.Command{
	Name:  "check",
	Usage: "Run a set of tests to confirm the host is suitable for PSM",
	Description: `
`,
	Action: doCheck,
}
var commandInit = cli.Command{
	Name:  "init",
	Usage: "Provide a summary of the current PSM installation",
	Description: `
`,
	Action: doInit,
}

func testHasPassed(message string, passed bool) {
	if passed {
		fmt.Printf(message+" ... %s\n", color.GreenString("Pass"))
	} else {
		fmt.Printf(message+" ... %s\n", color.RedString("Fail"))
	}
}

func doCheck(c *cli.Context) {
	fmt.Printf("Running Checks ...\n")
	host, _ := host.HostInfo()
	if host.PlatformFamily != "rhel" {
		testHasPassed("Confirm running on Redhat Linux host", false)
		// return
	} else {
		testHasPassed("Confirm running on Redhat Linux host", true)
	}
	ram, _ := mem.VirtualMemory()
	if ram.Total/1073741824 >= 8 {
		testHasPassed("Minimum amount of Memory", true)
	} else {
		testHasPassed("Minimum amount of Memory", false)
	}
	counts, _ := cpu.CPUCounts(true)
	if counts >= 4 {
		testHasPassed("Minimum amount of CPUs", true)
	} else {
		testHasPassed("Minimum amount of CPUs", false)
	}

}

func doInit(c *cli.Context) {

}
