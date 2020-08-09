package tvmazeclient

import (
	"time"
)

// Result - Search results are indexed by relevance (score)
type Result struct {
	Score float64
	Show  Show
}

//Schedule defines the day(s) and time a show airs
type Schedule struct {
	Time string   `json:"time"`
	Days []string `json:"days"`
}

//Rating defines a user generate rating for the show
type Rating struct {
	Average float64 `json:"average"`
}

//Country defines where a show is aired
type Country struct {
	Name     string `json:"name"`
	Code     string `json:"code"`
	Timezone string `json:"timezone"`
}

//Network defines the network and country a show is aired
type Network struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Country Country `json:"country"`
}

//Externals includes reference identifiers to other TV services
type Externals struct {
	Tvrage  int    `json:"tvrage"`
	Thetvdb int    `json:"thetvdb"`
	Imdb    string `json:"imdb"`
}

//Image has links to thumbnails for a show
type Image struct {
	Medium   string `json:"medium"`
	Original string `json:"original"`
}

//Self contains the URL to the current episode
type Self struct {
	Href string `json:"href"`
}

//Previousepisode contains the URL to the previous episode
type Previousepisode struct {
	Href string `json:"href"`
}

//Links has URLs to the current and previous episodes
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
