package oxforddict

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"path"
	"runtime"

	"github.com/suguanyang/libpp"
)

type appConfiguration struct {
	AppID   string `json:"app_id"`
	AppKey  string `json:"app_key"`
	BaseAPI string `json:"base_api"`
}

type wordQuery struct {
	sourceLanguage string
	word           string
}

func loadAppconfig() (*appConfiguration, error) {
	configuration := &appConfiguration{}

	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return configuration, errors.New("it was not possible to recover the information")
	}
	libpp.ReadJSONFile(path.Join(path.Dir(filename), "./app.config"), configuration)
	return configuration, nil
}

func sendRequest(query string) (*http.Response, error) {
	client := &http.Client{}
	appConfig, loadError := loadAppconfig()
	if loadError != nil {
		log.Fatal(loadError)
	}
	req, err := http.NewRequest("GET", appConfig.BaseAPI+query, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("app_id", appConfig.AppID)
	req.Header.Set("app_key", appConfig.AppKey)

	return client.Do(req)
}

func parseResponseBody(res *http.Response, container interface{}) error {
	defer res.Body.Close()
	if res.StatusCode == 404 {
		return errors.New("Please check your input")
	}
	if res.StatusCode == 500 {
		return errors.New("Server Error")
	}
	json.NewDecoder(res.Body).Decode(container)
	return nil
}

// LookUpWords request word
func LookUpWords(word string, sl string) (*[]Result, error) {
	query := "entries/" + sl + "/" + word
	res, err := sendRequest(query)
	if err != nil {
		log.Fatal(err)
	}
	data := new(OxfordSuccessResponse)
	queryError := parseResponseBody(res, data)
	return &data.Results, queryError
}
