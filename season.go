package tvmazeclient

import (
	"fmt"
	"strings"
	"time"

	strip "github.com/grokify/html-strip-tags-go"
)

// Season - Details about a single season
// 	Array of Seasons - http://api.tvmaze.com/shows/:id/seasons
type Season struct {
	ID           int         `json:"id" bson:"_id"`
	URL          string      `json:"url"`
	Number       int         `json:"number"`
	Name         string      `json:"name"`
	EpisodeOrder int         `json:"episodeOrder"`
	PremiereDate string      `json:"premiereDate"`
	EndDate      string      `json:"endDate"`
	Network      Network     `json:"network"`
	WebChannel   interface{} `json:"webChannel"`
	Image        Image       `json:"image"`
	Summary      string      `json:"summary"`
	Links        Links       `json:"_links"`
}

//Seasons slice of seasons for a show identifier
type Seasons struct {
	ID      int      `json:"id" bson:"_id"`
	Seasons []Season `json:"seasons" bson:"seasons"`
}

func (s Season) String() string {
	var builder strings.Builder
	builder.Grow(100)

	fmt.Fprintf(&builder, "|%-15s|%d\n", "ID", s.ID)
	fmt.Fprintf(&builder, "|%-15s|%d\n", "Number", s.Number)
	fmt.Fprintf(&builder, "|%-15s|%s\n", "URL", s.Links.Self.Href)

	fmt.Fprintf(&builder, "|%-15s|%s\n", "Network", s.Network.Name)

	fmt.Fprintf(&builder, "|%-15s|%s\n", "Start", s.PremiereDate)
	fmt.Fprintf(&builder, "|%-15s|%s\n", "End", s.EndDate)
	fmt.Fprintf(&builder, "|%-15s|%t\n", "Current", s.Current())

	fmt.Fprintf(&builder, "|%-15s|%s\n", "Overview", strip.StripTags(s.Summary))

	return builder.String()
}

// Current - Returns whether the season is active.
// Today's date must be between the start and end date OR the season completed within the last 30 days
// TODO: Update to make testable by injecting an evaluation date
func (s Season) Current() bool {
	premiere, err1 := time.Parse("2006-01-02", s.PremiereDate)
	end, err2 := time.Parse("2006-01-02", s.EndDate)
	today := time.Now()

	if err1 != nil || err2 != nil {
		return false
	}

	if today.After(premiere) && today.Before(end) {
		return true
	}

	today = today.AddDate(0, -1, 0)
	if today.Before(end) {
		return true
	}
	return false
}
