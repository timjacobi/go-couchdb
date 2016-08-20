package couchdb

type SearchResult struct {
	Doc    interface{}       `json:"doc"`
	Fields map[string]string `json:"fields"`
	ID     string            `json:"id"`
	Score  float64           `json:"score"`
	Order  []float64         `json:"order"`
}

type Search struct {
	ETag           string         `json:"etag,omitempty"`
	FetchDuration  int64          `json:"fetch_duration,omitempty"`
	Limit          int64          `json:"limit,omitempty"`
	Query          string         `json:"q,omitempty"`
	Rows           []SearchResult `json:"rows"`
	SearchDuration int64          `json:"search_duration,omitempty"`
	Skip           int64          `json:"skip,omitempty"`
	TotalRows      int64          `json:"total_rows"`
	Bookmark       int64          `json:"bookmark"`
}
