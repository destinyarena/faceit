package faceit

import (
	"net/http"

	"github.com/destinyarena/faceit/internal/logging"
)

var log = logging.New()

type (
	// Faceit exports functions that allow interaction with the Faceit REST API
	Faceit interface {
		Ban(hubid, guid, reason string) error
		Unban(hubid, guid string) error
		GetUserHubs(guid string) ([]*Hub, error)
		GetInvite(hubid string) (string, error)
		GetUserByID(guid string) (*User, error)
		GetUserByName(name string) (*User, error)
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
