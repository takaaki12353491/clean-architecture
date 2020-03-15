package outputdata

import "golang.org/x/oauth2"

// Login is auth login info
type Login struct {
	State string `json:"state"`
	URL   string `json:"redirect_url"`
}

type Callback struct {
	Token string
}

type Auth struct {
	GithubToken *oauth2.Token
}
