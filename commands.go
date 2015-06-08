package main
 
import (
	"log"
	"os"

	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandProvision,
	commandModify,
	commandSla,
	commandCir,
	commandCvlan,
	commandJitter,
	commandFrameloss,
	commandLatency,
	commandUtilisation,
	commandDiscards,
}

var commandProvision = cli.Command{
	Name:  "provision",
	Usage: "Provision a EVPLINE Service",
	Description: `This operation does blah ...
`,
	Action: doProvision,
}

var commandModify = cli.Command{
	Name:  "modify",
	Usage: "Modify the Bandwidth for a Service",
	Description: `
`,
	Action: doModify,
}

var commandSla = cli.Command{
	Name:  "sla",
	Usage: "Setup SLA monitoring for a Provisioned Service",
	Description: `
`,
	Action: doSla,
}

var commandCir = cli.Command{
	Name:  "cir",
	Usage: "Retrieve the provisioned CIR on a UNI in a Service",
	Description: `
`,
	Action: doCir,
}

var commandCvlan = cli.Command{
	Name:  "cvlan",
	Usage: "Confirm C-VLAN in use",
	Description: `
`,
	Action: doCvlan,
}

var commandJitter = cli.Command{
	Name:  "jitter",
	Usage: "Retrieve current PM data for Avg SLA Jitter",
	Description: `
`,
	Action: doJitter,
}

var commandFrameloss = cli.Command{
	Name:  "frameloss",
	Usage: "Retrieve current PM data for Frame loss",
	Description: `
`,
	Action: doFrameloss,
}

var commandLatency = cli.Command{
	Name:  "latency",
	Usage: "Retrieve current PM data for Latency",
	Description: `
`,
	Action: doLatency,
}

var commandUtilisation = cli.Command{
	Name:  "utilisation",
	Usage: "Retrieve current PM data for Service Utilization",
	Description: `
`,
	Action: doUtilisation,
}

var commandDiscards = cli.Command{
	Name:  "discards",
	Usage: "Retrieve current PM data for Service discards",
	Description: `
`,
	Action: doDiscards,
}

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func doProvision(c *cli.Context) {
}

func doModify(c *cli.Context) {
}

func doSla(c *cli.Context) {
}

func doCir(c *cli.Context) {
}

func doCvlan(c *cli.Context) {
}

func doJitter(c *cli.Context) {
}

func doFrameloss(c *cli.Context) {
}

func doLatency(c *cli.Context) {
}

func doUtilisation(c *cli.Context) {
}

func doDiscards(c *cli.Context) {
}
