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
	Hostname  string `json:"hostname"`
	Port      int    `json:"port"`
	ApiPrefix string `json:"api_prefix"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	APIKey    string `json:"api_key"`
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
	url := fmt.Sprintf("http://%s:%d/%s%s",
		g.Hostname, g.Port, g.ApiPrefix, command)
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

// GetAllSeries lists all series in your collection
func (g *Gonarr) GetAllSeries() MySeriesList {
	b := g.makeRequest("series")
	var mySeriesList MySeriesList
	json.Unmarshal(b, &mySeriesList)
	return mySeriesList
}

func (g *Gonarr) GetOneSeries(seriesId int) MySeries {
	url := fmt.Sprintf("series/%d", seriesId)
	b := g.makeRequest(url)
	var mySeries MySeries
	json.Unmarshal(b, &mySeries)
	return mySeries
}

func (g *Gonarr) GetSystemStatus() {
	b := g.makeRequest("system/status")
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

func (g *Gonarr) PrintJSON(i interface{}) {
	fmt.Println(toJSON(i))
}

func (g *Gonarr) Filter(i interface{}, keys ...string) interface{} {
	mkeys := make(map[string]bool)
	for _, k := range keys {
		mkeys[k] = true
	}

	sl, ok := i.([]interface{})
	if ok {
		var o []interface{}
		for _, e := range sl {
			o = append(o, g.Filter(e, keys...))
		}
		return o
	}
	m, ok := i.(map[string]interface{})
	if ok {
		o := make(map[string]interface{})
		for k, v := range m {
			if _, haskey := mkeys[k]; haskey {
				o[k] = g.Filter(v, keys...)
			}
		}
		return o
	}
	return i
}
