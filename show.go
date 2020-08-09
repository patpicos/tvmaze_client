package tvmazeclient

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	strip "github.com/grokify/html-strip-tags-go"
)

// Show - Individual Show details from the search results list
type Show struct {
	ID            int    `json:"id" bson:"_id"`
	URL           string `json:"url"`
	Name          string `json:"name"`
	SanitizedName string
	Type          string      `json:"type"`
	Language      string      `json:"language"`
	Genres        []string    `json:"genres"`
	Status        string      `json:"status"`
	Runtime       int         `json:"runtime"`
	Premiered     string      `json:"premiered"`
	OfficialSite  string      `json:"officialSite"`
	Schedule      Schedule    `json:"schedule"`
	Rating        Rating      `json:"rating"`
	Weight        int         `json:"weight"`
	Network       Network     `json:"network"`
	WebChannel    interface{} `json:"webChannel"`
	Externals     Externals   `json:"externals"`
	Image         Image       `json:"image"`
	Summary       string      `json:"summary"`
	Updated       int         `json:"updated"`
	Links         Links       `json:"_links"`
}

//Print - Display details at the command line output
func (s Show) String() string {
	var builder strings.Builder
	builder.Grow(100)

	fmt.Fprintf(&builder, "|%-15s|%s\n", "Name", s.Name)
	fmt.Fprintf(&builder, "|%-15s|%d\n", "ID", s.ID)
	fmt.Fprintf(&builder, "|%-15s|%s\n", "Show URL", s.ShowURL())
	fmt.Fprintf(&builder, "|%-15s|%s\n", "Seasons URL", s.SeasonsURL())

	fmt.Fprintf(&builder, "|%-15s|%s\n", "Classification", s.Type)
	fmt.Fprintf(&builder, "|%-15s|%s\n", "Genres", strings.Join(s.Genres, ", "))
	fmt.Fprintf(&builder, "|%-15s|%d\n", "Premiered", s.PremiereYear())

	fmt.Fprintf(&builder, "|%-15s|%s\n", "Language", s.Language)
	fmt.Fprintf(&builder, "|%-15s|%t\n", "Foreign", s.Foreign())

	fmt.Fprintf(&builder, "|%-15s|%s\n", "Status", s.Status)
	fmt.Fprintf(&builder, "|%-15s|%t\n", "Cancelled", s.Cancelled())

	fmt.Fprintf(&builder, "|%-15s|%d\n", "Seasons", s.SeasonCount())
	seasons, _ := ListSeasons(s.ID)
	for _, season := range seasons {
		fmt.Fprintf(&builder, "|S%-14d|%s|%s|%s|%t\n", season.Number, season.URL, season.PremiereDate, season.EndDate, season.Current())
	}
	fmt.Fprintf(&builder, "|%-15s|%t\n", "Foreign", s.Foreign())
	fmt.Fprintf(&builder, "|%-15s|%s\n", "Overview", strip.StripTags(s.Summary))

	// fmt.Println("Runtime:", s.Runtime)
	// fmt.Println("Schedule:", s.Schedule)
	// fmt.Println("Network:", s.Network)
	return builder.String()
}

// SeasonCount - Number of seasons
func (s Show) SeasonCount() int {
	seasons, _ := ListSeasons(s.ID)
	return len(seasons)
}

//ShowSearchURL return the search url for the show
func (s Show) ShowSearchURL() string {
	return fmt.Sprintf("http://api.tvmaze.com/search/shows?q=%s", s.Name)
}

//ShowURL returns the URL for the show
func (s Show) ShowURL() string {
	return fmt.Sprintf("http://api.tvmaze.com/shows/%d", s.ID)
}

//SeasonsURL returns the URL for the show's seasons
func (s Show) SeasonsURL() string {
	return fmt.Sprintf("http://api.tvmaze.com/shows/%d/seasons", s.ID)
}

//PremiereYear - The year the show started
func (s Show) PremiereYear() int {
	runes := []rune(s.Premiered)
	safeSubstring := string(runes[0:4])
	converted, _ := strconv.ParseInt(safeSubstring, 10, 0)
	return int(converted)
}

//Cancelled - Return whether the show is cancelled
func (s Show) Cancelled() bool {
	switch s.Status {
	case "Continuing":
		return false
	case "Running":
		return false
	case "In Development":
		return false
	default:
		return true
	}
}

//Country returns the sanitized country where the show is created
// TODO: Move this business logic out to a consumer
func (s Show) Country() string {
	matched, _ := regexp.MatchString(s.Network.Name, "{?i)ABC|Disney|AMC")
	if s.Network.Country.Code == "US" || strings.Contains(s.Name, "(US)") || strings.Contains(s.Name, " US") || matched {
		return "USA"
	} else if s.Network.Country.Name != "" {
		return s.Network.Country.Name
	}

	matched, _ = regexp.MatchString(s.Network.Name, "(?i)UK|United.Kingdom|BBC|E4")
	if matched || strings.Contains(s.Name, "UK") {
		return "UK"
	}
	matched, _ = regexp.MatchString(s.Network.Name, "(?i)NZ|New.Zealand")
	if matched || strings.Contains(s.Name, "New Zealand") {
		return "New Zealand"
	}
	matched, _ = regexp.MatchString(s.Network.Name, `(?i)\(CA\)|Canada`)
	if matched || strings.Contains(s.Name, "Canada") {
		return "Canada"
	}
	if strings.Contains(s.Network.Name, "Japan") {
		return "Japan"
	}
	if s.Network.Name == "Netflix" {
		return "USA"
	}

	return "INVALID_COUNTRY" // return a bogus country. Further testing needed against data in TVMaze
}

//Foreign - Return whether the show is foreign
// TODO: Move this business logic out to a consumer
func (s Show) Foreign() bool {
	switch s.Language {
	case "English":
		return false
	case "":
		return false
	default:
		return true
	}
}
