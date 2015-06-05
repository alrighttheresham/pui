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
	Usage: "",
	Description: `Provision a EVPLINE Service
`,
	Action: doProvision,
}

var commandModify = cli.Command{
	Name:  "modify",
	Usage: "",
	Description: `
`,
	Action: doModify,
}

var commandSla = cli.Command{
	Name:  "sla",
	Usage: "",
	Description: `
`,
	Action: doSla,
}

var commandCir = cli.Command{
	Name:  "cir",
	Usage: "",
	Description: `
`,
	Action: doCir,
}

var commandCvlan = cli.Command{
	Name:  "cvlan",
	Usage: "",
	Description: `
`,
	Action: doCvlan,
}

var commandJitter = cli.Command{
	Name:  "jitter",
	Usage: "",
	Description: `
`,
	Action: doJitter,
}

var commandFrameloss = cli.Command{
	Name:  "frameloss",
	Usage: "",
	Description: `
`,
	Action: doFrameloss,
}

var commandLatency = cli.Command{
	Name:  "latency",
	Usage: "",
	Description: `
`,
	Action: doLatency,
}

var commandUtilisation = cli.Command{
	Name:  "utilisation",
	Usage: "",
	Description: `
`,
	Action: doUtilisation,
}

var commandDiscards = cli.Command{
	Name:  "discards",
	Usage: "",
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
