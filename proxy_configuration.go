package goapify

import (
	"fmt"
	"net/url"
)

type ProxyConfigurationOptions struct {
	Password string
	Group    string

	HostName string
	Port     string

	CountryCode string
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
	connectionString := fmt.Sprintf(`http://%s:%s@%s:%s`, p.options.Group, p.options.Password, p.options.HostName, p.options.Port)

	proxyUrl, err := url.Parse(connectionString)
	if err != nil {
		return nil, err
	}

	return proxyUrl, nil
}
