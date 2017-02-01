package main

import (
	"fmt"
	"log"
	"os"

	flags "github.com/jessevdk/go-flags"
	"github.com/kgigitdev/gonarr"
)

func main() {
	var opts gonarr.GonarrOptions
	parser := flags.NewParser(&opts, flags.Default)
	args, err := parser.ParseArgs(os.Args)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	log.Println("Args: ", args)

	g := gonarr.NewGonarrFromOptions(opts)

	log.Println("Gonarr is: ", g)

	if opts.Status {
		fmt.Println(g.GetSystemStatus())
	} else if opts.ListCommands {
		fmt.Println(g.ListCommands())
	} else if opts.ListCommand > 0 {
		fmt.Println(g.ListCommand(opts.ListCommand))
	} else if opts.SeasonSearch {
		if opts.SeriesID == 0 {
			log.Fatal("No series id supplied.")
		}
		if opts.SeasonNumber == 0 {
			log.Fatal("No season number supplied.")
		}
		fmt.Println(g.SonarrCommand("SeasonSearch", opts.SeriesID, opts.SeasonNumber))
	} else if opts.RescanSeries {
		if opts.SeriesID == 0 {
			log.Fatal("No series id supplied.")
		}
		fmt.Println(g.RescanSeries(opts.SeriesID))
	} else if opts.RefreshSeries {
		if opts.SeriesID == 0 {
			log.Fatal("No series id supplied.")
		}
		fmt.Println(g.RefreshSeries(opts.SeriesID))
	} else if opts.ListSeries {
		series := g.GetAllSeries()
		fmt.Println(series)
	} else if opts.Info {
		if opts.SeriesID == 0 {
			log.Fatal("No series id supplied.")
		}
		fmt.Println(g.GetOneSeries(opts.SeriesID))
	} else if opts.SetMonitor || opts.ToggleMonitor {
		if opts.SeriesID == 0 {
			log.Fatal("No series id supplied.")
		}
		if opts.SeasonNumber == 0 {
			log.Fatal("No season number supplied.")
		}
		cmd := g.GetOneSeries(opts.SeriesID)
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
