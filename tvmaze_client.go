package tvmazeclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/* URL's of interest:
Search Shows: http://api.tvmaze.com/search/shows?q=girls
Individual show: http://api.tvmaze.com/shows/:id
	Episode List:  http://api.tvmaze.com/shows/:id/episodes
	Season List: http://api.tvmaze.com/shows/:id/seasons

Specific Episodes:
	http://api.tvmaze.com/shows/:id/episodebynumber?season=:season&number=:number
	http://api.tvmaze.com/shows/1/episodesbydate?date=:date

Show Index (Useful for refreshes)
	http://api.tvmaze.com/updates/shows (~47,000 shows)

Rate Limiting:
API calls are rate limited to allow at least 20 calls every 10 seconds per IP address.
If you exceed this rate, you might receive an HTTP 429 error. We say at least, because rate limiting takes
place on the backend but not on the edge cache.
So if your client is only requesting common/popular endpoints like shows or episodes
(as opposed to more unique endpoints like searches or embedding), you're likely to never hit the limit.
For an optimal throughput, simply let your client back off for a few seconds when it receives a 429.
*/

//GetShow - Return an individual TV Show
// 	http://api.tvmaze.com/shows/:id
func GetShow(showID int) (Show, error) {
	url := fmt.Sprintf("http://api.tvmaze.com/shows/%d", showID)
	response, err := http.Get(url)

	if err != nil {
		return Show{}, err //fmt.Errorf("error reading contents from %s, Error: %s StatusCode %d", url, err, response.StatusCode)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return Show{}, fmt.Errorf("Page %s returned HTTP %s", url, response.Status)
	}
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Show{}, fmt.Errorf("error reading contents from %s, Error: %s", url, err)
	}

	var results Show
	err = json.Unmarshal(content, &results)
	if err != nil {
		return Show{}, fmt.Errorf("error decoding JSON from %s, Error: %s", url, err)
	}

	// Add SanitizedName to Show Object
	// results.SanitizedName = utils.SanitizeShowName(results.Name)
	return results, nil
}

//Search - Query the TVMaze API for a search keyword
func Search(query string) []Result {
	url := fmt.Sprintf("http://api.tvmaze.com/search/shows?q=%s", query)
	response, err := http.Get(url)
	if err != nil {
		log.Printf("Error: Fetching url failed %s", url)
		return nil
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error reading from %s", url)
		return nil
	}

	var results []Result

	err = json.Unmarshal(content, &results)
	if err != nil {
		log.Printf("Error decoding from %s", url)
		return nil
	}

	return results
}

//ListSeasons - Retrieve list of seasons for a specific show
func ListSeasons(showID int) ([]Season, error) {
	url := fmt.Sprintf("http://api.tvmaze.com/shows/%d/seasons", showID)
	response, err := http.Get(url)
	if err != nil {
		return nil, err //fmt.Errorf("Error: Fetching url failed %s", url)
	}
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Error: Error reading contents from %s, Error: %s", url, err)
	}

	var results []Season
	err = json.Unmarshal(content, &results)
	if err != nil {
		return nil, fmt.Errorf("Error: Error decoding JSON %s, Error: %s", url, err)
	}

	return results, nil
}

//GetShowUpdates - Get list of show updates
// 	http://api.tvmaze.com/updates/shows
func GetShowUpdates(lastUpdateEpoch int) (map[string]int, error) {
	url := "http://api.tvmaze.com/updates/shows"
	response, err := http.Get(url)
	if err != nil {
		return nil, err //fmt.Errorf("error reading contents from %s", url)
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading contents from %s, Error: %s", url, err)
	}

	var results map[string]int
	err = json.Unmarshal(content, &results)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON from %s, Error: %s", url, err)
	}
	return results, nil
}
