package main

import (
	"fmt"
	"log"
	"os"

	flags "github.com/jessevdk/go-flags"
	"github.com/kgigitdev/gonarr"
)

var opts struct {
	Config     string `short:"c" long:"config" description:"config file" default:"gonarr.json"`
	Search     string `short:"s" long:"search" description:"Search series"`
	ListSeries bool   `short:"l" long:"list-series" description:"List all series in your collection"`
	Info       int    `short:"i" long:"info" description:"Show info about one series in your collection"`
	Season     int    `long:"season" description:"Season (for modifying stuff)"`

	ToggleMonitor bool `short:"m" long:"toggle-monitor" description:"Season (for modifying stuff)"`
	Status        bool `long:"status" description:"Get system status"`

	Full bool `long:"full" description:"List full JSON"`
}

func main() {
	parser := flags.NewParser(&opts, flags.Default)
	args, err := parser.ParseArgs(os.Args)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	log.Println("Args: ", args)

	g := gonarr.NewGonarrFromConfigFile(opts.Config)

	log.Println("Gonarr is: ", g)

	if opts.Status {
		g.GetSystemStatus()
	} else if opts.ListSeries {
		series := g.GetAllSeries()
		fmt.Println(series)
	} else if opts.Info > 0 {
		series := g.GetOneSeries(opts.Info)
		fmt.Println(series)
		if opts.ToggleMonitor && opts.Season > 1 {
			log.Println("Toggling montoring flag ...")
		}
	} else if opts.Search != "" {
		fmt.Println(g.SearchSeries(opts.Search))
	}
}
