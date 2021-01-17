package shorten

// Redirect model
type Redirect struct {
	Code      string `json:"code"`
	URL       string `json:"url"`
	CreatedAt int64
}
