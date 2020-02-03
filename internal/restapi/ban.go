package faceit

import (
    "fmt"
    "errors"
    "bytes"
    "net/http"
    "encoding/json"
)

type BanPayload struct {
    HubId   string `json:"hubId"`
    UserId  string `json:"userId"`
    Reasion string `json:"reason"`
}

func (f *Faceit) Ban(hubid, guid, reasion string) error {
    url := fmt.Sprintf("https://api.faceit.com/hubs/v1/hub/%s/ban/%s", hubid, guid)

    payload := BanPayload{
        HubId: hubid,
        UserId: guid,
        Reasion: reasion,
    }

    body, err := json.Marshal(payload)
    if err != nil {
        return err
    }

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
    if err != nil {
        return err
    }

    resp, err := f.UC.Do(req)
    if err != nil {
        return err
    }

    defer resp.Body.Close()

    if resp.StatusCode != 200 && resp.StatusCode != 201 {
        err = errors.New(fmt.Sprintf("Server returned: %d", resp.StatusCode))
        return err
    }

    return nil
}
