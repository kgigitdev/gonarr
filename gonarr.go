package gonarr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

func (g *Gonarr) makeGetRequest(command string) string {
	return g.makeRequest("GET", command, nil)
}

func (g *Gonarr) makePostRequest(command string, payload io.Reader) string {
	return g.makeRequest("POST", command, payload)
}

func (g *Gonarr) makePutRequest(command string, payload io.Reader) string {
	return g.makeRequest("PUT", command, payload)
}

func (g *Gonarr) makeRequest(method string, command string, payload io.Reader) string {
	url := fmt.Sprintf("http://%s:%d/%s%s",
		g.Hostname, g.Port, g.ApiPrefix, command)
	fmt.Println(url)
	req, err := http.NewRequest(method, url, payload)
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
	return string(body)
}

func (g *Gonarr) UpdateOneSeries(cmd MySeries) string {
	url := fmt.Sprintf("series")
	payload := toJSONB(cmd)
	return g.makePutRequest(url, bytes.NewBuffer(payload))
}

// RefreshSeries invokes the RefreshSeries API command
func (g *Gonarr) RefreshSeries(seriesId int) string {
	url := fmt.Sprintf("command/RefreshSeries?seriesId=%d", seriesId)
	return g.makePostRequest(url, nil)
}

// RescanSeries invokes the RescanSeries API command
func (g *Gonarr) RescanSeries(seriesId int) string {
	url := fmt.Sprintf("command/RescanSeries?seriesId=%d", seriesId)
	return g.makePostRequest(url, nil)
}

// SeasonSearch invokes the SeasonSearch API command
func (g *Gonarr) SonarrCommand(commandName string, seriesId int, seasonNumber int) string {
	cmd := NewSonarrCommand(commandName, seriesId, seasonNumber)
	payload := toJSONB(cmd)
	r := bytes.NewBuffer(payload)
	return g.makePostRequest("command", r)
}

// SeriesSearch invokes the SeasonSearch API command
func (g *Gonarr) SeriesSearch(seriesId int) string {
	url := fmt.Sprintf("command/SeriesSearch?seriesId=%d",
		seriesId)
	return g.makePostRequest(url, nil)
}

// ListCommands ...
func (g *Gonarr) ListCommands() string {
	// return g.makeGetRequest("commands")
	return g.makeGetRequest("command")
}

func (g *Gonarr) ListCommand(commandId int) string {
	url := fmt.Sprintf("command/%d", commandId)
	return g.makeGetRequest(url)
}

// ListCalendar ...
func (g *Gonarr) ListCalendar() string {
	return g.makeGetRequest("calendar")
}

// GetAllSeries lists all series in your collection
func (g *Gonarr) GetAllSeries() MySeriesList {
	b := g.makeGetRequest("series")
	var mySeriesList MySeriesList
	json.Unmarshal([]byte(b), &mySeriesList)
	return mySeriesList
}

func (g *Gonarr) GetOneSeries(seriesId int) MySeries {
	url := fmt.Sprintf("series/%d", seriesId)
	b := g.makeGetRequest(url)
	var mySeries MySeries
	json.Unmarshal([]byte(b), &mySeries)
	return mySeries
}

func (g *Gonarr) GetSystemStatus() string {
	return g.makeGetRequest("system/status")
}

// SearchSeries searches for series matching the search term
func (g *Gonarr) SearchSeries(searchTerm string) SeriesInformationList {
	// FIXME: Escape things like spaces, etc.
	url := fmt.Sprintf("series/lookup?term=%s", searchTerm)
	b := g.makeGetRequest(url)
	var seriesInformationList SeriesInformationList
	json.Unmarshal([]byte(b), &seriesInformationList)
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
