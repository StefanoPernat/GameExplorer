package reddit

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"time"
)

var client *http.Client

func init() {
	client = &http.Client{
		Timeout: time.Second * 15,
	}
}

//GetTodayHotDeals retrives top 100 hot deals from reddit api
func GetTodayHotDeals() ([]Deal, error) {
	address := fmt.Sprintf("%s/hot.json?limit=%d", redditBaseURL, limit)

	fmt.Println(address)

	req, err := http.NewRequest("get", address, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request for Deals is not %d, returned %d instead", http.StatusOK, res.StatusCode)
	}

	if res.Body == nil {
		return nil, fmt.Errorf("response body seems to be empty")
	}

	defer res.Body.Close()

	fromReddit, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return parseListResponse(fromReddit)
}

func parseListResponse(response []byte) ([]Deal, error) {
	deals := make([]Deal, 0)

	var parsed map[string]interface{}
	err := json.Unmarshal(response, &parsed)
	if err != nil {
		return nil, err
	}

	if _, found := parsed["data"]; !found {
		return nil, errors.New("Unable to parse respone")
	}

	data := parsed["data"].(map[string]interface{})
	if _, found := data["children"]; !found {
		return nil, errors.New("Unable to parse respone")
	}

	list := data["children"].([]interface{})

	for _, deal := range list {
		singleEntry := deal.(map[string]interface{})
		if _, found := singleEntry["data"]; !found {
			return nil, errors.New("Unable to parse respone")
		}

		game := singleEntry["data"].(map[string]interface{})

		singleDeal := Deal{}

		if value, ok := game["title"].(string); ok {
			singleDeal.Title = value
		}

		if value, ok := game["url"].(string); ok {
			singleDeal.URL = value
		}

		if value, ok := game["permalink"].(string); ok {
			singleDeal.Permalink = value
		}

		if value, ok := game["author"].(string); ok {
			singleDeal.Author = value
		}
		if value, ok := game["selftext"].(string); ok {
			singleDeal.Text = value
		}
		if value, ok := game["created_utc"].(float64); ok {
			singleDeal.UtcDate = value
		}

		if singleDeal.isValid() {
			today := time.Now()

			sec, dec := math.Modf(singleDeal.UtcDate)
			date := time.Unix(int64(sec), int64(dec*(1e9)))

			if date.Day() == today.Day() && date.Month() == today.Month() && date.Year() == today.Year() {
				deals = append(deals, singleDeal)
			}

		}
	}

	return deals, nil
}
