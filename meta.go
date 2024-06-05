package govultr

// Meta represents the available pagination information
type Meta struct {
	Links *Links `json:"links"`
	Total int    `json:"total"`
}

// Links represent the next/previous cursor in your pagination calls
type Links struct {
	Next string `json:"next"`
	Prev string `json:"prev"`
}
