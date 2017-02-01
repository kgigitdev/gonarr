package gonarr

type GonarrOptions struct {
	Config       string `short:"c" long:"config" description:"config file" default:"gonarr.json"`
	Search       string `short:"s" long:"search" description:"Search series"`
	ListSeries   bool   `short:"l" long:"list-series" description:"List all series in your collection"`
	Info         bool   `short:"i" long:"info" description:"Show info about one series in your collection"`
	SeasonNumber int    `long:"season-number" description:"Season Number"`

	SeriesID int `long:"series" description:"Series Id"`

	ToggleMonitor bool `short:"m" long:"toggle-monitor" description:"Toggle the monitoring flag for a season"`

	SetMonitor bool `long:"set-monitor" description:"Set the monitoring flag for a seaon"`

	Status bool `long:"status" description:"Get system status"`

	SeasonSearch bool `long:"season-search" description:"Invoke the SeasonSearch command"`

	RefreshSeries bool `long:"refresh-series" description:"Invoke the RefreshSeries command"`

	RescanSeries bool `long:"rescan-series" description:"Invoke the RescanSeries command"`

	ListCommands bool `long:"list-commands" description:"List the command ids of all commands currently in flight"`

	ListCommand int `long:"list-command" description:"Get the status of one command (by command id)"`

	Full bool `long:"full" description:"List full JSON"`

	DebugOut bool `long:"debugout" description:"Print all JSON being sent"`
}
