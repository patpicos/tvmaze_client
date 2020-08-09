package tvmazeclient

import (
	"time"
)

// Result - Search results are indexed by relevance (score)
type Result struct {
	Score float64
	Show  Show
}

type Schedule struct {
	Time string   `json:"time"`
	Days []string `json:"days"`
}
type Rating struct {
	Average float64 `json:"average"`
}
type Country struct {
	Name     string `json:"name"`
	Code     string `json:"code"`
	Timezone string `json:"timezone"`
}
type Network struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Country Country `json:"country"`
}
type Externals struct {
	Tvrage  int    `json:"tvrage"`
	Thetvdb int    `json:"thetvdb"`
	Imdb    string `json:"imdb"`
}
type Image struct {
	Medium   string `json:"medium"`
	Original string `json:"original"`
}
type Self struct {
	Href string `json:"href"`
}
type Previousepisode struct {
	Href string `json:"href"`
}
type Links struct {
	Self            Self            `json:"self"`
	Previousepisode Previousepisode `json:"previousepisode"`
}

// Episode - Details about a single episode
// 	Array of Episode - http://api.tvmaze.com/shows/:id/episodes
// 	Single Episode - http://api.tvmaze.com/shows/:id/episodebynumber?season=:season&number=:number
type Episode struct {
	ID       int       `json:"id"`
	URL      string    `json:"url"`
	Name     string    `json:"name"`
	Season   int       `json:"season"`
	Number   int       `json:"number"`
	Airdate  string    `json:"airdate"`
	Airtime  string    `json:"airtime"`
	Airstamp time.Time `json:"airstamp"`
	Runtime  int       `json:"runtime"`
	Image    Image     `json:"image"`
	Summary  string    `json:"summary"`
	Links    Links     `json:"_links"`
}
