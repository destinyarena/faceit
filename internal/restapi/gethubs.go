package faceit

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "gopkg.in/go-playground/validator.v9"
)

const (
    PC_GAME_ID   = "destiny2"
    PS4_GAME_ID  = "destiny2_ps4"
    XBOX_GAME_ID = "destiny2_xbox"
)

type (
    FaceitHubPayload struct {
        Items []*FaceitHub `json:"items" validate:"required"`
    }

    FaceitHub struct {
        Hubid  string `json:"hub_id" validate:"required"`
        Name   string `json:"name" validate:"required"`
        GameID string `json:"game_id" validate:"required"`
    }
)

func (f *Faceit) GetUserHubs(guid string) ([]*FaceitHub, error) {
    url := fmt.Sprintf("https://open.faceit.com/data/v4/players/%s/hubs", guid)

    req, _ := http.NewRequest("GET", url, nil)
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

    var body FaceitHubPayload
    json.Unmarshal([]byte(rawbody), &body)

    v := validator.New()
    if err = v.Struct(body); err != nil {
        return nil, err
    }

    hubs := make([]*FaceitHub, 0)

    for _, hub := range body.Items {
        id := hub.GameID
        if id == PC_GAME_ID || id == PS4_GAME_ID || id == XBOX_GAME_ID {
            log.Infof("Hub ID: %s Hub Name: %s", hub.Hubid, hub.Name)
            hubs = append(hubs, hub)
        }
    }


    return hubs, nil
}
