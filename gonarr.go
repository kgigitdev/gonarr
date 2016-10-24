package gonarr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Gonarr ...
type Gonarr struct {
	Hostname string `json:"hostname"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	APIKey   string `json:"api_key"`
}

func (g Gonarr) String() string {
	j, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(j)
}

// NewGonarrFromConfigFile creates and returns a pointer to a new
// Gonarr object, from a persisted JSON configuration file at the
// specified path.
func NewGonarrFromConfigFile(path string) *Gonarr {
	fh, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer fh.Close()
	config, err := ioutil.ReadAll(fh)
	if err != nil {
		log.Fatal(err)
	}
	g := &Gonarr{}
	json.Unmarshal(config, g)
	return g
}

func (g *Gonarr) makeRequest(command string) []byte {
	url := fmt.Sprintf("http://%s:%d/api/%s", g.Hostname, g.Port, command)
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("X-Api-Key", g.APIKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}

// ListCommands ...
func (g *Gonarr) ListCommands() []byte {
	return g.makeRequest("commands")
}

// ListCalendar ...
func (g *Gonarr) ListCalendar() []byte {
	return g.makeRequest("calendar")
}

// ListSeries lists all series in your collection
func (g *Gonarr) ListSeries() {
	url := fmt.Sprintf("series")
	b := g.makeRequest(url)
	fmt.Println(string(b))
}

// SearchSeries searches for series matching the search term
func (g *Gonarr) SearchSeries(searchTerm string) SeriesInformationList {
	// FIXME: Escape things like spaces, etc.
	url := fmt.Sprintf("series/lookup?term=%s", searchTerm)
	b := g.makeRequest(url)
	var seriesInformationList SeriesInformationList
	json.Unmarshal(b, &seriesInformationList)
	return seriesInformationList
}
