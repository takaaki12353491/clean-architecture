package outputdata

// Login is auth login info
type Login struct {
	State string `json:"state"`
	URL   string `json:"redirect_url"`
}

type Callback struct {
	Token string
}
