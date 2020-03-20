package outputdata

// Auth ...
type Auth struct {
	State string `json:"state"`
	URL   string `json:"redirect_url"`
}

type Callback struct {
	Token string
}
