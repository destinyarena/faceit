package faceit

import (
	"fmt"
	"net/http"
)

func (f *faceit) Unban(hubid, guid, reasion string) error {
	url := fmt.Sprintf("https://api.faceit.com/hubs/v1/hub/%s/ban/%s", hubid, guid)

	req, err := http.NewRequest("POST", url, nil)
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
