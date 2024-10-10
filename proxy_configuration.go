package goapify

type ProxyConfigurationOptions struct {
	Password    string
	Groups      []string
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
