package oxforddict

import "github.com/suguanyang/libpp"

type appConfiguration struct {
	AppID   string `json:"app_id"`
	AppKey  string `json:"app_key"`
	BaseAPI string `json:"base_api"`
}

func loadAppconfig() appConfiguration {
	configuration := appConfiguration{}

	libpp.ReadJSONFile("app.config", &configuration)
	libpp.PrintStruct(configuration)
	return configuration
}

// RequestOxfordDictionaries request word
func RequestOxfordDictionaries(words []string, sl string) string {
	appConfig := loadAppconfig()
	return appConfig.BaseAPI
}
