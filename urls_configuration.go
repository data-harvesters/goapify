package goapify

type UrlConfigurations []UrlConfiguration

type UrlConfiguration struct {
	Url string `json:"url"`
}

func (c UrlConfigurations) GetUrls() []string {
	var urls []string
	for _, urlConfig := range c {
		urls = append(urls, urlConfig.Url)
	}

	return urls
}
