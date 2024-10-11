package goapify

import (
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"time"
)

type ProxyConfigurationOptions struct {
	UseApifyProxy bool      `json:"useApifyProxy"`
	Groups        *[]string `json:"apifyProxyGroups"`
	CountryCode   *string   `json:"apifyProxyCountry"`

	ProxyUrls *[]string `json:"proxyUrls"`

	password string `json:"-"`
	hostName string `json:"-"`
	port     string `json:"-"`
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

	if p.options.Groups == nil {
		return nil, errors.New("no proxy groups found")
	}
	groups := *p.options.Groups

	group := groups[r.Intn(len(groups))]

	connectionString := fmt.Sprintf(`http://%s:%s@%s:%s`,
		fmt.Sprintf("groups-%s", group),
		p.options.password,
		p.options.hostName,
		p.options.port)

	proxyUrl, err := url.Parse(connectionString)
	if err != nil {
		return nil, err
	}

	return proxyUrl, nil
}
