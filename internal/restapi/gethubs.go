package faceit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"gopkg.in/go-playground/validator.v9"
)

const (
	pcGameID = "destiny2"
)

type (
	// HubPayload is the wrapper for the hub payload
	HubPayload struct {
		Items []*Hub `json:"items" validate:"required"`
	}

	// Hub is the hub payload
	Hub struct {
		Hubid  string `json:"hub_id" validate:"required"`
		Name   string `json:"name" validate:"required"`
		GameID string `json:"game_id" validate:"required"`
	}
)

func (f *faceit) GetUserHubs(guid string) ([]*Hub, error) {
	url := fmt.Sprintf("https://open.faceit.com/data/v4/players/%s/hubs", guid)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := f.SC.Do(req)
	if err != nil {
		return nil, err
	}

	rawbody, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		fmt.Println(string(rawbody))
		err = fmt.Errorf("Server responded with: %d", resp.StatusCode)
		return nil, err
	}

	var body HubPayload
	json.Unmarshal([]byte(rawbody), &body)

	v := validator.New()
	if err = v.Struct(body); err != nil {
		return nil, err
	}

	hubs := make([]*Hub, 0)

	for _, hub := range body.Items {
		id := hub.GameID
		if id == pcGameID {
			log.Infof("Hub ID: %s Hub Name: %s", hub.Hubid, hub.Name)
			hubs = append(hubs, hub)
		}
	}

	return hubs, nil
}
