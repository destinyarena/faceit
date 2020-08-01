package faceit

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (f *Faceit) GetInvite(hubid string) (error, string) {
	reqBody, _ := json.Marshal(ReqInvitePayload{
		EntityID:   hubid,
		EntityType: "hub",
		Type:       "regular",
		MaxAge:     0,
		MaxUses:    1,
	})

	url := "https://api.faceit.com/invitations/v1/invite"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return err, ""
	}

	resp, err := f.UC.Do(req)
	if err != nil {
		return err, ""
	}

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		err = fmt.Errorf("Server response code: %d", resp.StatusCode)
		return err, ""
	}

	defer resp.Body.Close()

	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, ""
	}

	var body Invite
	json.Unmarshal([]byte(raw), &body)

	if body.Payload.Code == "" {
		err = errors.New("Invalid invite code")
		return err, ""
	}

	return nil, body.Payload.Code
}
