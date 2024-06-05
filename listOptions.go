package govultr

// ListOptions are the available query params
type ListOptions struct {
	Cursor      string `url:"cursor,omitempty"`
	MainIP      string `url:"main_ip,omitempty"`
	Label       string `url:"label,omitempty"`
	Tag         string `url:"tag,omitempty"`
	Region      string `url:"region,omitempty"`
	Description string `url:"description,omitempty"`
	PerPage     int    `url:"per_page,omitempty"`
}
