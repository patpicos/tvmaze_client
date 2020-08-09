package tvmazeclient

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYear(t *testing.T) {
	s := Show{Premiered: "2019-01-10"}
	assert.Equal(t, s.PremiereYear(), 2019, "The extracted year should match")
}

func TestBillions(t *testing.T) {
	s, _ := GetShow(3606)
	assert.GreaterOrEqual(t, s.SeasonCount(), 5, "Billions has at least 5 seasons")
	assert.Equal(t, s.PremiereYear(), 2016, "The extracted year should match")
	assert.False(t, s.Foreign(), "English is not foreign")
}

func TestForeign(t *testing.T) {
	s, _ := GetShow(23619) //Korean
	assert.True(t, s.Foreign(), "Korean is foreign")
	assert.True(t, s.Cancelled(), "Show is cancelled")
}

func TestURLs(t *testing.T) {
	s, _ := GetShow(3606)
	assert.Equal(t, s.ShowSearchURL(), "http://api.tvmaze.com/search/shows?q=Billions")
	assert.Equal(t, s.ShowURL(), "http://api.tvmaze.com/shows/3606")
	assert.Equal(t, s.SeasonsURL(), "http://api.tvmaze.com/shows/3606/seasons")
}

func TestSeasonStringer(t *testing.T) {
	s, _ := GetShow(3606)
	seasons, err := ListSeasons(s.ID)
	assert.Nil(t, err)
	assert.Equal(t, seasons[0].String(), `|ID             |11227
|Number         |1
|URL            |http://api.tvmaze.com/seasons/11227
|Network        |Showtime
|Start          |2016-01-17
|End            |2016-04-10
|Current        |false
|Overview       |What happens when two voracious power players at the top of their fields go head to head? Brilliant hedge fund titan Bobby "Axe" Axelrod and brash U.S. District Attorney Chuck Rhoades play a dangerous, winner-take-all game of cat and mouse where the stakes run into ten figures. Both are ultimately forced to answer the question: what is power worth?
`)
}
