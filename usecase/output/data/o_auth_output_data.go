package outputdata

// Auth ...
type Auth struct {
	URL string `json:"redirect_url"`
}

type Callback struct {
	Token string
}
