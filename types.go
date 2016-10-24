package gonarr

import "encoding/json"

// Season stores information for one or more entries in the
// "seasons" field of SeriesInformation
type Season struct {
	SeasonNumber int  `json:"seasonNumber"`
	Monitored    bool `json:"monitored"`
}

type SeriesImage struct {
	CoverType string `json:"coverType"`
	URL       string `json:"url"`
}

type Ratings struct {
	Votes int `json:"votes"`
	Value int `json:"value"`
}

// SeriesInformation is a struct for storing a series information
type SeriesInformation struct {
	Title             string        `json:"title"`
	SortTitle         string        `json:"sortTitle"`
	SeasonCount       int           `json:"seasonCount"`
	Status            string        `json:"status"`
	Overview          string        `json:"overview"`
	Network           string        `json:"network"`
	AirTime           string        `json:"airTime"`
	Images            []SeriesImage `json:"images"`
	RemotePoster      string        `json:"remotePoster"`
	Seasons           []Season      `json:"seasons"`
	Year              int           `json:"year"`
	ProfileID         int           `json:"profileId"`
	SeasonFolder      bool          `json:"seasonFolder"`
	Monitored         bool          `json:"monitored"`
	UseSceneNumbering bool          `json:"useSceneNumbering"`
	Runtime           int           `json:"runtime"`
	TVDBID            int           `json:"tvdbId"`
	TVRageID          int           `json:"tvRageId"`
	TVMazeID          int           `json:"tvMazeId"`
	FirstAired        string        `json:"firstAired"`
	SeriesType        string        `json:"seriesType"`
	CleanTitle        string        `json:"cleanTitle"`
	IMDBID            string        `json:"imdbId"`
	TitleSlug         string        `json:"titleSlug"`
	Certification     string        `json:"certification"`
	Genres            []string      `json:"genres"`
	Tags              []string      `json:"tags"`
	Added             string        `json:"added"`
	Ratings           Ratings       `json:"ratings"`
	QualityProfileID  int           `json:"qualityProfileId"`
}

// SeriesInformationList ...
type SeriesInformationList []SeriesInformation

func (s SeriesInformationList) String() string {
	j, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(j)
}
