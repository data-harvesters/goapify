package goapify

import (
	"fmt"
	"math/rand"
	"net/url"
	"time"
)

type ProxyConfigurationOptions struct {
	UseApifyProxy bool     `json:"useApifyProxy"`
	Groups        []string `json:"apifyProxyGroups"`
	CountryCode   string   `json:"apifyProxyCountry"`

	ProxyUrls *[]string `json:"proxyUrls"`

	password string
	hostName string
	port     string
}

type ProxyConfiguration struct {
	options *ProxyConfigurationOptions
}

func newProxyConfiguration(options *ProxyConfigurationOptions) *ProxyConfiguration {
	return &ProxyConfiguration{
		options: options,
	}
}

func (p *ProxyConfiguration) Proxy() (*url.URL, error) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	if p.options.ProxyUrls != nil && !p.options.UseApifyProxy {
		proxyUrls := *p.options.ProxyUrls

		proxyUrl := proxyUrls[r.Intn(len(*p.options.ProxyUrls))]

		u, err := url.Parse(proxyUrl)
		if err != nil {
			return nil, err
		}

		return u, nil
	}
	group := p.options.Groups[r.Intn(len(p.options.Groups))]

	connectionString := fmt.Sprintf(`http://%s:%s@%s:%s`,
		fmt.Sprintf("group-%s", group),
		p.options.password,
		p.options.hostName,
		p.options.port)

	proxyUrl, err := url.Parse(connectionString)
	if err != nil {
		return nil, err
	}

	return proxyUrl, nil
}
