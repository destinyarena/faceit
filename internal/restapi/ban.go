package faceit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// BanPayload is the payload used for bans :)
type BanPayload struct {
	HubID   string `json:"hubId"`
	UserID  string `json:"userId"`
	Reasion string `json:"reason"`
}

// Ban is the function that drops the hammer :)
func (f *faceit) Ban(hubid, guid, reason string) error {
	url := fmt.Sprintf("https://api.faceit.com/hubs/v1/hub/%s/ban/%s", hubid, guid)

	payload := BanPayload{
		HubID:   hubid,
		UserID:  guid,
		Reasion: reason,
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
		return fmt.Errorf("Server returned: %d", resp.StatusCode)
	}

	return nil
}
