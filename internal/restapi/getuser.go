package faceit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"gopkg.in/go-playground/validator.v9"
)

const game = "destiny2"

// GetUserByID gets a Faceit user based in the faceit GUID
func (f *faceit) GetUserByID(guid string) (*User, error) {
	url := fmt.Sprintf("https://open.faceit.com/data/v4/players/%s", guid)

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

	var body RawUser
	json.Unmarshal([]byte(rawbody), &body)

	v := validator.New()
	if err = v.Struct(body); err != nil {
		return nil, err
	}

	skillLevel := 1

	if _, ok := body.Games[game]; ok {
		skillLevel = body.Games[game].SkillLevel
	}

	user := &User{
		ID:         body.ID,
		Username:   body.Username,
		SkillLevel: skillLevel,
	}

	return user, nil
}

func (f *faceit) GetUserByName(name string) (*User, error) {
	url := fmt.Sprintf("https://open.faceit.com/data/v4/players?nickname=%s", name)

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

	var body RawUser
	json.Unmarshal([]byte(rawbody), &body)

	v := validator.New()
	if err = v.Struct(body); err != nil {
		return nil, err
	}

	skillLevel := 1

	if _, ok := body.Games[game]; ok {
		skillLevel = body.Games[game].SkillLevel
	}

	user := &User{
		ID:         body.ID,
		Username:   body.Username,
		SkillLevel: skillLevel,
	}

	return user, nil
}
