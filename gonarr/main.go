package main

import (
	"fmt"
	"log"
	"os"

	flags "github.com/jessevdk/go-flags"
	"github.com/kgigitdev/gonarr"
)

var opts struct {
	Config       string `short:"c" long:"config" description:"config file" default:"gonarr.json"`
	Search       string `short:"s" long:"search" description:"Search series"`
	ListSeries   bool   `short:"l" long:"list-series" description:"List all series in your collection"`
	Info         bool   `short:"i" long:"info" description:"Show info about one series in your collection"`
	SeasonNumber int    `long:"season-number" description:"Season Number"`

	SeriesId int `long:"series" description:"Series Id"`

	ToggleMonitor bool `short:"m" long:"toggle-monitor" description:"Toggle the monitoring flag for a season"`

	SetMonitor bool `long:"set-monitor" description:"Set the monitoring flag for a seaon"`

	Status bool `long:"status" description:"Get system status"`

	SeasonSearch bool `long:"season-search" description:"Invoke the SeasonSearch command"`

	RefreshSeries bool `long:"refresh-series" description:"Invoke the RefreshSeries command"`

	RescanSeries bool `long:"rescan-series" description:"Invoke the RescanSeries command"`

	ListCommands bool `long:"list-commands" description:"List the command ids of all commands currently in flight"`

	ListCommand int `long:"list-command" description:"Get the status of one command (by command id)"`

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
		fmt.Println(g.GetSystemStatus())
	} else if opts.ListCommands {
		fmt.Println(g.ListCommands())
	} else if opts.ListCommand > 0 {
		fmt.Println(g.ListCommand(opts.ListCommand))
	} else if opts.SeasonSearch {
		if opts.SeriesId == 0 {
			log.Fatal("No series id supplied.")
		}
		if opts.SeasonNumber == 0 {
			log.Fatal("No season number supplied.")
		}
		fmt.Println(g.SeasonSearch(opts.SeriesId, opts.SeasonNumber))
	} else if opts.RescanSeries {
		if opts.SeriesId == 0 {
			log.Fatal("No series id supplied.")
		}
		fmt.Println(g.RescanSeries(opts.SeriesId))
	} else if opts.RefreshSeries {
		if opts.SeriesId == 0 {
			log.Fatal("No series id supplied.")
		}
		fmt.Println(g.RefreshSeries(opts.SeriesId))
	} else if opts.ListSeries {
		series := g.GetAllSeries()
		fmt.Println(series)
	} else if opts.Info {
		if opts.SeriesId == 0 {
			log.Fatal("No series id supplied.")
		}
		fmt.Println(g.GetOneSeries(opts.SeriesId))
	} else if opts.SetMonitor || opts.ToggleMonitor {
		if opts.SeriesId == 0 {
			log.Fatal("No series id supplied.")
		}
		if opts.SeasonNumber == 0 {
			log.Fatal("No season number supplied.")
		}
		cmd := g.GetOneSeries(opts.SeriesId)
		for i, season := range cmd.Seasons {
			if season.SeasonNumber == opts.SeasonNumber {
				if opts.SetMonitor {
					season.Monitored = true
				} else {
					season.Monitored = !season.Monitored
				}
				cmd.Seasons[i] = season
				break
			}
		}
		fmt.Println("Posting ...")
		b := g.UpdateOneSeries(cmd)
		s := string(b)
		fmt.Println(s)
	} else if opts.Search != "" {
		fmt.Println(g.SearchSeries(opts.Search))
	}
}
