package faceit

import (
	"net/http"

	"github.com/destinyarena/faceit/internal/logging"
)

var log = logging.New()

type Faceit struct {
	UC *http.Client
	SC *http.Client
}

func New(apitoken, usertoken string) *Faceit {
	uc := newClient(usertoken)
	sc := newClient(apitoken)
	return &Faceit{
		UC: uc,
		SC: sc,
	}
}
