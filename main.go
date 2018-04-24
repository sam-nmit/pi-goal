package main

import (
	"flag"
	"log"

	"github.com/sam-nmit/pi-goal/pidns"
	"github.com/sam-nmit/pi-goal/rest"
	"github.com/sam-nmit/pi-goal/utils"
)

var (
	ConfigName = flag.String("config", "config.json", "Filename for config")
	NewConfig  = flag.Bool("resetconfig", false, "Creates a new config file")

	config *utils.Config
)

func main() {
	flag.Parse()
	log.Println("Pi-goal started...")
	log.Printf("Loading config \"%s\"\n", *ConfigName)

	if *NewConfig {
		config = utils.ResetConfig(*ConfigName)
	} else {
		config = utils.ConfigFromFile(*ConfigName)
		if config == nil {
			log.Printf("Invalid conifig file \"%s\". Use -resetconfig to create a new config.", *ConfigName)
		}
	}

	dnsServer := pidns.DefaultServer(config.Server.Address)
	dnsServer.LoadRules(config.RuleFolder, config.Rules)

	if config.Rest != nil {
		r := rest.NewServer(dnsServer, config.Rest.Address)
		r.Start()
	}

	if err := dnsServer.Start(); err != nil {
		log.Fatalf("pi-goal failed to start udp listener.")
	} else {
		log.Println("pi-hole closing...")
	}
}
