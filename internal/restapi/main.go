package faceit

import (
	"net/http"

	"github.com/destinyarena/faceit/internal/logging"
)

var log = logging.New()

type (
	// Faceit exports functions that allow interaction with the Faceit REST API
	Faceit interface {
		Ban(string, string, string) error
		GetUserHubs(string) ([]*Hub, error)
		GetInvite(string) (string, error)
		GetUserByID(string) (*User, error)
		GetUserByName(string) (*User, error)
		Unban(string, string, string) error
	}

	faceit struct {
		UC *http.Client
		SC *http.Client
	}
)

// New returns a new Faceit REST API Wrapper
func New(apitoken, usertoken string) (Faceit, error) {
	uc := newClient(usertoken)
	sc := newClient(apitoken)
	return &faceit{
		UC: uc,
		SC: sc,
	}, nil
}
