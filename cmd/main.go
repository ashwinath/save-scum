package main

import (
	"flag"
	"log"
	"sync"

	"github.com/ashwinath/save-scum/pkg/config"
	"github.com/ashwinath/save-scum/pkg/shell"
)

func main() {
	configPath := flag.String("config", "", "config location")
	flag.Parse()
	log.Print(*configPath)
	if configPath == nil {
		log.Fatalf("config not provided.")
	}

	c, err := config.New(*configPath)
	if err != nil {
		log.Fatalf("error parsing config: %s", err)
	}

	// Rsync files
	var filewg sync.WaitGroup
	for _, file := range c.Files {
		filewg.Add(1)
		go func(f config.FileConfig) {
			defer filewg.Done()
			o, err := shell.Rsync(f.Flags, f.From, f.To)
			logOutput(o, err)
			if f.Chown.Enabled {
				o, err := shell.ChownRecursive(f.To, f.Chown.User, f.Chown.Group)
				logOutput(o, err)
			}
		}(file)
	}

	filewg.Wait()
}

func logOutput(o *string, err error) {
	if o != nil {
		log.Printf("output: %s", *o)
	}
	if err != nil {
		log.Printf("stderr: %s", err.Error())
	}
}
