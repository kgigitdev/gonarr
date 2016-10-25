package gonarr

import "encoding/json"

// Season stores information for one or more entries in the
// "seasons" field of SeriesInformation
type Season struct {
	SeasonNumber int  `json:"seasonNumber,omitempty"`
	Monitored    bool `json:"monitored"`
}

type SeriesImage struct {
	CoverType string `json:"coverType,omitempty"`
	URL       string `json:"url,omitempty"`
}

type Ratings struct {
	Votes int `json:"votes,omitempty"`
	Value int `json:"value,omitempty"`
}

type AddedFragment struct {
	Added string `json:"added,omitempty"`
}

type AirTimeFragment struct {
	AirTime string `json:"airTime,omitempty"`
}

type CleanTitleFragment struct {
	CleanTitle string `json:"cleanTitle,omitempty"`
}

type FirstAiredFragment struct {
	FirstAired string `json:"firstAired,omitempty"`
}

type GenresFragment struct {
	Genres []string `json:"genres,omitempty"`
}

type ImagesFragment struct {
	Images []SeriesImage `json:"images,omitempty"`
}

type IMDBIDFragment struct {
	IMDBID string `json:"imdbId,omitempty"`
}

type MonitoredFragment struct {
	Monitored bool `json:"monitored,omitempty"`
}

type NetworkFragment struct {
	Network string `json:"network,omitempty"`
}

type OverviewFragment struct {
	Overview string `json:"overview,omitempty"`
}

type ProfileIDFragment struct {
	ProfileID int `json:"profileId,omitempty"`
}

type QualityProfileIDFragment struct {
	QualityProfileID int `json:"qualityProfileId,omitempty"`
}

type RatingsFragment struct {
	Ratings Ratings `json:"ratings,omitempty"`
}

type RuntimeFragment struct {
	Runtime int `json:"runtime,omitempty"`
}

type SeasonCountFragment struct {
	SeasonCount int `json:"seasonCount,omitempty"`
}

type SeasonsFragment struct {
	Seasons []Season `json:"seasons,omitempty"`
}

type SeriesTypeFragment struct {
	SeriesType string `json:"seriesType,omitempty"`
}

type SortTitleFragment struct {
	SortTitle string `json:"sortTitle,omitempty"`
}

type StatusFragment struct {
	Status string `json:"status,omitempty"`
}

type TagsFragment struct {
	Tags []string `json:"tags,omitempty"`
}

type TitleFragment struct {
	Title string `json:"title,omitempty"`
}

type TitleSlugFragment struct {
	TitleSlug string `json:"titleSlug,omitempty"`
}

type TVMazeIDFragment struct {
	TVMazeID int `json:"tvMazeId,omitempty"`
}

type TVRageIDFragment struct {
	TVRageID int `json:"tvRageId,omitempty"`
}

type TVDBIDFragment struct {
	TVDBID int `json:"tvdbId,omitempty"`
}

type UseSceneNumberingFragment struct {
	UseSceneNumbering bool `json:"useSceneNumbering,omitempty"`
}

type YearFragment struct {
	Year int `json:"year,omitempty"`
}

type SeasonFolderFragment struct {
	SeasonFolder bool `json:"seasonFolder,omitempty"`
}

type IDFragment struct {
	ID int `json:"id,omitempty"`
}

type PathFragment struct {
	Path string `json:"path,omitempty"`
}

// CommonFragments is a compositional struct containing all fragments
// common to "all series" and "all my series".
type CommonFragments struct {
	AddedFragment
	AirTimeFragment
	CleanTitleFragment
	FirstAiredFragment
	GenresFragment
	ImagesFragment
	IMDBIDFragment
	MonitoredFragment
	NetworkFragment
	OverviewFragment
	PathFragment
	ProfileIDFragment
	QualityProfileIDFragment
	RatingsFragment
	RuntimeFragment
	SeasonCountFragment
	SeasonFolderFragment
	SeasonsFragment
	SeriesTypeFragment
	StatusFragment
	TagsFragment
	TitleFragment
	TitleSlugFragment
	TVMazeIDFragment
	TVDBIDFragment
	UseSceneNumberingFragment
	YearFragment
}

// SeriesInformation is a struct for storing a series information
type SeriesInformation struct {
	CommonFragments

	// Thes might just not have appeared in the JSON output I
	// happened to see; no guarantee that they can't appear in
	// "my series"
	RemotePoster  string `json:"remotePoster,omitempty"`
	Certification string `json:"certification,omitempty"`
}

// SeriesInformationList ...
type SeriesInformationList []SeriesInformation

func (s SeriesInformationList) String() string {
	return toJSON(s)
}

type MySeries struct {
	IDFragment
	CommonFragments
}

type MySeriesList []MySeries

func (m MySeries) String() string {
	return toJSON(m)
}

func (m MySeriesList) String() string {
	return toJSON(m)
}

func toJSONB(i interface{}) []byte {
	j, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return []byte(err.Error())
	}
	return j
}

func toJSON(i interface{}) string {
	return string(toJSONB(i))
}
