package faceit

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "gopkg.in/go-playground/validator.v9"
)

func (f *Faceit) GetUser(guid string) (*FaceitUser, error) {
    url := fmt.Sprintf("https://open.faceit.com/data/v4/players/%s", guid)

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


    var body RawUser
    json.Unmarshal([]byte(rawbody), &body)

    v := validator.New()
    if err = v.Struct(body); err != nil {
        return nil, err
    }

    skillLevel := 1

    if _, ok := body.Games["destiny2"]; ok {
        skillLevel = body.Games["destiny2"].SkillLevel
    }

    user := &FaceitUser{
        Id:         body.Id,
        Username:   body.Username,
        SkillLevel: skillLevel,
    }

    return user, nil
}
