package github

import (
	"context"
	"net/http"

	"github.com/google/go-github/github"
	"github.com/rajatparida86/wtfdial"
	"golang.org/x/oauth2"
)

// Authenticator is a github authenticator
type Authenticator struct {
}

// Authenticate authenticates against Github and returns the currently authenticated wtf.User
func (a *Authenticator) Authenticate(token string) (*wtf.User, error) {
	ctxt := context.Background()
	staticToken := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	// A oAuth2 client to be passed to github Client
	oAuth2Client := oauth2.NewClient(ctxt, staticToken)
	// A github client to do the authentiocation
	githubClient := github.NewClient(oAuth2Client)
	// Get the authenticated user by passing the 'user' string as empty
	user, resp, err := githubClient.Users.Get(ctxt, "")
	if err != nil {
		switch resp.StatusCode {
		case http.StatusForbidden, http.StatusUnauthorized:
			return nil, wtf.NewAuthenticationError("Can not authenticate user with Github")
		default:
			return nil, err
		}
	}
	//Return the github user as a wtf.User
	return &wtf.User{
		UserName: user.GetName(),
		ID:       wtf.UserID(user.GetID()),
	}, nil

}
