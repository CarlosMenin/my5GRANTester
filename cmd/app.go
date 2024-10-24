package main

import (
	"my5G-RANTester/config"
	"my5G-RANTester/internal/templates"

	// "fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

const version = "0.1"

func init() {
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
	spew.Config.Indent = "\t"

	log.Info("my5G-RANTester version " + version)

}

func main() {

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "ue",
				Aliases: []string{"ue"},
				Usage:   "Testing an ue attached with configuration",
				Action: func(c *cli.Context) error {
					name := "Testing an ue attached with configuration"
					cfg := config.Data

					log.Info("---------------------------------------")
					log.Info("[TESTER] Starting test function: ", name)
					log.Info("[TESTER][UE] Number of UEs: ", 1)
					log.Info("[TESTER][GNB] Control interface IP/Port: ", cfg.GNodeB.ControlIF.Ip, "/", cfg.GNodeB.ControlIF.Port)
					log.Info("[TESTER][GNB] Data interface IP/Port: ", cfg.GNodeB.DataIF.Ip, "/", cfg.GNodeB.DataIF.Port)
					log.Info("[TESTER][AMF] AMF IP/Port: ", cfg.AMF.Ip, "/", cfg.AMF.Port)
					log.Info("---------------------------------------")
					templates.TestAttachUeWithConfiguration()
					return nil
				},
			},
			{
				Name:    "gnb",
				Aliases: []string{"gnb"},
				Usage:   "Testing an gnb attached with configuration",
				Action: func(c *cli.Context) error {
					name := "Testing an gnb attached with configuration"
					cfg := config.Data

					log.Info("---------------------------------------")
					log.Info("[TESTER] Starting test function: ", name)
					log.Info("[TESTER][GNB] Number of GNBs: ", 1)
					log.Info("[TESTER][GNB] Control interface IP/Port: ", cfg.GNodeB.ControlIF.Ip, "/", cfg.GNodeB.ControlIF.Port)
					log.Info("[TESTER][GNB] Data interface IP/Port: ", cfg.GNodeB.DataIF.Ip, "/", cfg.GNodeB.DataIF.Port)
					log.Info("[TESTER][AMF] AMF IP/Port: ", cfg.AMF.Ip, "/", cfg.AMF.Port)
					log.Info("---------------------------------------")
					templates.TestAttachGnbWithConfiguration()
					return nil
				},
			},
			{
				Name:    "load-test",
				Aliases: []string{"load-test"},
				Usage: "\nLoad endurance stress tests.\n" +
					"Example for testing multiple UEs: load-test -n 5 \n",
				Flags: []cli.Flag{
					&cli.IntFlag{Name: "number-of-ues", Value: 1, Aliases: []string{"n"}},
				},
				Action: func(c *cli.Context) error {
					var numUes int
					name := "Testing registration of multiple UEs"
					cfg := config.Data

					if c.IsSet("number-of-ues") {
						numUes = c.Int("number-of-ues")
					} else {
						log.Info(c.Command.Usage)
						return nil
					}

					log.Info("---------------------------------------")
					log.Info("[TESTER] Starting test function: ", name)
					log.Info("[TESTER][UE] Number of UEs: ", numUes)
					log.Info("[TESTER][GNB] gNodeB control interface IP/Port: ", cfg.GNodeB.ControlIF.Ip, "/", cfg.GNodeB.ControlIF.Port)
					log.Info("[TESTER][GNB] gNodeB data interface IP/Port: ", cfg.GNodeB.DataIF.Ip, "/", cfg.GNodeB.DataIF.Port)
					log.Info("[TESTER][AMF] AMF IP/Port: ", cfg.AMF.Ip, "/", cfg.AMF.Port)
					log.Info("---------------------------------------")
					templates.TestMultiUesInQueue(numUes)

					return nil
				},
			},
			{
				Name:    "load-test-parallel",
				Aliases: []string{"load-test-parallel"},
				Usage: "\nLoad endurance stress tests.\n" +
					"Example for testing multiple UEs: load-test-parallel -n 1000 -d 30 -t 30 -a\n",
				Flags: []cli.Flag{
					&cli.IntFlag{Name: "number-of-ues", Value: 1, Aliases: []string{"n"}},
					&cli.IntFlag{Name: "delay-per-ue", Value: 1, Aliases: []string{"d"}},
					&cli.IntFlag{Name: "startup-delay", Value: 1, Aliases: []string{"t"}},
					&cli.BoolFlag{Name: "enable-analytics", Aliases: []string{"a"}},
				},
				Action: func(c *cli.Context) error {
					var numUes int
					var delayUes int
					var delayStart int
					var showAnalytics bool
					name := "Testing registration of multiple UEs in parallel"
					cfg := config.Data

					if c.IsSet("number-of-ues") && c.IsSet("delay-per-ue") && c.IsSet("startup-delay") {
						numUes = c.Int("number-of-ues")
						delayUes = c.Int("delay-per-ue")
						delayStart = c.Int("startup-delay")
						showAnalytics = c.Bool("enable-analytics")
					} else {
						log.Info(c.Command.Usage)
						return nil
					}

					log.Info("---------------------------------------")
					log.Info("[TESTER] Starting test function: ", name)
					log.Info("[TESTER][UE] Number of UEs: ", numUes)
					log.Info("[TESTER][GNB] gNodeB control interface IP/Port: ", cfg.GNodeB.ControlIF.Ip, "/", cfg.GNodeB.ControlIF.Port)
					log.Info("[TESTER][GNB] gNodeB data interface IP/Port: ", cfg.GNodeB.DataIF.Ip, "/", cfg.GNodeB.DataIF.Port)
					log.Info("[TESTER][AMF] AMF IP/Port: ", cfg.AMF.Ip, "/", cfg.AMF.Port)
					log.Info("---------------------------------------")
					templates.TestMultiUesInParallel(numUes, delayUes, delayStart, showAnalytics)

					return nil
				},
			},
			{
				Name:    "load-test-division",
				Aliases: []string{"load-test-division"},
				Usage: "\nLoad endurance stress tests.\n" +
					"Example for testing multiple UEs: load-test-division -n 100 -d 30 -t 30 -u 5 -i 10 -a\n",
				Flags: []cli.Flag{
					&cli.IntFlag{Name: "number-of-ues", Value: 1, Aliases: []string{"n"}},
					&cli.IntFlag{Name: "initial-delay", Value: 1, Aliases: []string{"d"}},
					&cli.IntFlag{Name: "startup-delay", Value: 1, Aliases: []string{"t"}},
					&cli.IntFlag{Name: "num-divisor", Value: 1, Aliases: []string{"u"}},
					&cli.IntFlag{Name: "interval-division", Value: 1, Aliases: []string{"i"}},
					&cli.BoolFlag{Name: "enable-analytics", Aliases: []string{"a"}},
				},
				Action: func(c *cli.Context) error {
					var numUes int
					var delayUes int
					var delayStart int
					var numDivisor int
					var intervalDivision int
					var showAnalytics bool
					name := "Testing registration of multiple UEs in parallel and dividing delay"
					cfg := config.Data

					if c.IsSet("number-of-ues") && c.IsSet("initial-delay") && c.IsSet("startup-delay") && c.IsSet("num-divisor") && c.IsSet("interval-division") {
						numUes = c.Int("number-of-ues")
						delayUes = c.Int("initial-delay")
						delayStart = c.Int("startup-delay")
						numDivisor = c.Int("num-divisor")
						intervalDivision = c.Int("interval-division")
						showAnalytics = c.Bool("enable-analytics")
					} else {
						log.Info(c.Command.Usage)
						return nil
					}

					log.Info("---------------------------------------")
					log.Info("[TESTER] Starting test function: ", name)
					log.Info("[TESTER][UE] Number of UEs: ", numUes)
					log.Info("[TESTER][GNB] gNodeB control interface IP/Port: ", cfg.GNodeB.ControlIF.Ip, "/", cfg.GNodeB.ControlIF.Port)
					log.Info("[TESTER][GNB] gNodeB data interface IP/Port: ", cfg.GNodeB.DataIF.Ip, "/", cfg.GNodeB.DataIF.Port)
					log.Info("[TESTER][AMF] AMF IP/Port: ", cfg.AMF.Ip, "/", cfg.AMF.Port)
					log.Info("---------------------------------------")
					templates.TestMultiUesDivision(numUes, delayUes, delayStart, numDivisor, intervalDivision, showAnalytics)

					return nil
				},
			},
			{
				Name:    "load-test-decrement",
				Aliases: []string{"load-test-decrement"},
				Usage: "\nLoad endurance stress tests.\n" +
					"Example for testing multiple UEs: load-test-decrement -n 100 -d 30 -t 30 -u 5 -i 10 -a\n",
				Flags: []cli.Flag{
					&cli.IntFlag{Name: "number-of-ues", Value: 1, Aliases: []string{"n"}},
					&cli.IntFlag{Name: "initial-delay", Value: 1, Aliases: []string{"d"}},
					&cli.IntFlag{Name: "startup-delay", Value: 1, Aliases: []string{"t"}},
					&cli.IntFlag{Name: "num-decrement", Value: 1, Aliases: []string{"u"}},
					&cli.IntFlag{Name: "interval-decrement", Value: 1, Aliases: []string{"i"}},
					&cli.BoolFlag{Name: "enable-analytics", Aliases: []string{"a"}},
				},
				Action: func(c *cli.Context) error {
					var numUes int
					var delayUes int
					var delayStart int
					var numDecrement int
					var intervalDecrement int
					var showAnalytics bool
					name := "Testing registration of multiple UEs in parallel and decrement delay"
					cfg := config.Data

					if c.IsSet("number-of-ues") && c.IsSet("initial-delay") && c.IsSet("startup-delay") && c.IsSet("num-decrement") && c.IsSet("interval-decrement") {
						numUes = c.Int("number-of-ues")
						delayUes = c.Int("initial-delay")
						delayStart = c.Int("startup-delay")
						numDecrement = c.Int("num-decrement")
						intervalDecrement = c.Int("interval-decrement")
						showAnalytics = c.Bool("enable-analytics")
					} else {
						log.Info(c.Command.Usage)
						return nil
					}

					log.Info("---------------------------------------")
					log.Info("[TESTER] Starting test function: ", name)
					log.Info("[TESTER][UE] Number of UEs: ", numUes)
					log.Info("[TESTER][GNB] gNodeB control interface IP/Port: ", cfg.GNodeB.ControlIF.Ip, "/", cfg.GNodeB.ControlIF.Port)
					log.Info("[TESTER][GNB] gNodeB data interface IP/Port: ", cfg.GNodeB.DataIF.Ip, "/", cfg.GNodeB.DataIF.Port)
					log.Info("[TESTER][AMF] AMF IP/Port: ", cfg.AMF.Ip, "/", cfg.AMF.Port)
					log.Info("---------------------------------------")
					templates.TestMultiUesDecrease(numUes, delayUes, delayStart, numDecrement, intervalDecrement, showAnalytics)

					return nil
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
