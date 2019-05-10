package reddit

/*
https://www.reddit.com/r/gamedeals/hot.json?limit=100
*/

// Deal is the model for a Reddit's GamesDeal post
type Deal struct {
	Title     string  `json:"title"`
	URL       string  `json:"url"`
	Permalink string  `json:"permalink"`
	Author    string  `json:"author"`
	Text      string  `json:"selftext"`
	UtcDate   float64 `json:"created_utc"`
}

func (d *Deal) isValid() bool {
	return d.Title != "" && d.URL != "" && d.UtcDate > 0
}
