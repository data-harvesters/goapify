package goapify

type UrlConfigurations []UrlConfiguration

type UrlConfiguration struct {
	Url string `json:"url"`

	Method   *string         `json:"method"`
	Payload  *string         `json:"payload"`
	UserData *map[string]any `json:"userData"`
	Headers  map[string]any  `json:"headers"`
}

func (c *UrlConfiguration) GetMethod() string {
	if c.Method == nil {
		return ""
	}

	return *c.Method
}

func (c *UrlConfiguration) GetPayload() string {
	if c.Payload == nil {
		return ""
	}

	return *c.Payload
}

func (c *UrlConfiguration) GetUserData() map[string]any {
	if c.UserData == nil {
		return nil
	}

	return *c.UserData
}

func (c *UrlConfiguration) GetHeaders() map[string]any {
	if c.Headers == nil {
		return nil
	}

	return *&c.Headers
}

func (c UrlConfigurations) GetUrls() []string {
	var urls []string
	for _, urlConfig := range c {
		urls = append(urls, urlConfig.Url)
	}

	return urls
}
