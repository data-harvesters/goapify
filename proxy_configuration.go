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
}

type ProxyConfiguration struct {
	options *ProxyConfigurationOptions
}

func newProxyConfiguration(options *ProxyConfigurationOptions) *ProxyConfiguration {
	return &ProxyConfiguration{
		options: options,
	}
}

type ProxyChoiceOptions struct {
	random bool

	useRandomGroup bool
	specificGroup  string

	specificProxy string
}

type ProxyChoiceOption func(*ProxyChoiceOptions)

func UseRandomProxy() ProxyChoiceOption {
	return func(c *ProxyChoiceOptions) {
		c.random = true
	}
}

func useRandomGroup() ProxyChoiceOption {
	return func(c *ProxyChoiceOptions) {
		c.useRandomGroup = true
	}
}

func useGroup(group string) ProxyChoiceOption {
	return func(c *ProxyChoiceOptions) {
		if c.useRandomGroup {
			return
		}

		c.specificGroup = group
	}
}

func UseSpecificProxy(proxyUrl string) ProxyChoiceOption {
	return func(c *ProxyChoiceOptions) {
		if c.random {
			return
		}

		c.specificProxy = proxyUrl
	}
}

func (p *ProxyConfiguration) Proxy(options ...ProxyChoiceOption) (*url.URL, error) {
	proxyChoiceOptions := &ProxyChoiceOptions{}

	for _, f := range options {
		f(proxyChoiceOptions)
	}

	if p.options.UseApifyProxy {
		if p.options.Groups == nil {
			return nil, errors.New("no apify proxy groups found")
		}
		groups := *p.options.Groups

		if len(groups) == 0 {
			return nil, errors.New("no apify proxy groups found")
		}

		group := "auto"
		if proxyChoiceOptions.useRandomGroup {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))

			group = groups[r.Intn(len(groups))]
		} else {
			for _, g := range groups {
				if g == proxyChoiceOptions.specificGroup {
					group = g
					break
				}
			}
		}

		connectionString := fmt.Sprintf(`http://%s:%s@%s:%s`,
			fmt.Sprintf("groups-%s", group),
			variables["APIFY_PROXY_PASSWORD"],
			variables["APIFY_PROXY_HOSTNAME"],
			variables["APIFY_PROXY_PORT"])

		proxyUrl, err := url.Parse(connectionString)
		if err != nil {
			return nil, err
		}

		return proxyUrl, nil
	}

	if p.options.ProxyUrls == nil {
		return nil, errors.New("no proxies given")
	}
	proxyUrls := *p.options.ProxyUrls

	if len(proxyUrls) == 0 {
		return nil, errors.New("no proxies given")
	}

	if proxyChoiceOptions.random {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))

		proxyUrl := proxyUrls[r.Intn(len(*p.options.ProxyUrls))]

		u, err := url.Parse(proxyUrl)
		if err != nil {
			return nil, err
		}

		return u, nil
	}

	if proxyChoiceOptions.specificProxy == "" {
		u, err := url.Parse(proxyUrls[0])
		if err != nil {
			return nil, err
		}

		return u, nil
	}

	for _, v := range proxyUrls {
		if v == proxyChoiceOptions.specificProxy {
			u, err := url.Parse(v)
			if err != nil {
				return nil, err
			}

			return u, nil
		}
	}

	return nil, fmt.Errorf("failed to find proxy with given specified proxy")
}

func (p *ProxyConfiguration) Proxy(options ...ProxyChoiceOption) (*url.URL, error) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	if p.options.ProxyUrls != nil && !p.options.UseApifyProxy {
		proxyUrls := *p.options.ProxyUrls

		if len(proxyUrls) == 0 {
			return nil, errors.New("no proxies given")
		}

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

	if len(groups) == 0 {
		return nil, errors.New("no proxy groups found")
	}

	group := groups[r.Intn(len(groups))]

	connectionString := fmt.Sprintf(`http://%s:%s@%s:%s`,
		fmt.Sprintf("groups-%s", group),
		variables["APIFY_PROXY_PASSWORD"],
		variables["APIFY_PROXY_HOSTNAME"],
		variables["APIFY_PROXY_PORT"])

	proxyUrl, err := url.Parse(connectionString)
	if err != nil {
		return nil, err
	}

	return proxyUrl, nil
}
