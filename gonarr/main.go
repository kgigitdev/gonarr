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

	if opts.ListSeries {
		g.ListSeries()
	} else if opts.Search != "" {
		fmt.Println(g.SearchSeries(opts.Search))
	}
}
