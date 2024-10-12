package goapify

import (
	"fmt"
	"os"
)

var (
	variables = make(map[string]string)

	variableNames = []string{
		"ACTOR_ID",
		"ACTOR_TASK_ID",

		"APIFY_USER_ID",

		"APIFY_DEFAULT_KEY_VALUE_STORE_ID",
		"APIFY_TOKEN",
		"APIFY_DEFAULT_DATASET_ID",
	}

	proxyVariables = []string{
		"APIFY_PROXY_PASSWORD",
		"APIFY_PROXY_HOSTNAME",
		"APIFY_PROXY_PORT",
	}
)

func ensureEnvironment() {
	for k := range variables {
		env := os.Getenv(k)
		if env == "" {
			panic(fmt.Sprintf("%s environment variable was not found.", k))
		}

		variables[k] = env
	}
}

func ensureProxyEnvironment() {
	for _, k := range proxyVariables {
		env := os.Getenv(k)
		if env == "" {
			panic(fmt.Sprintf("%s environment variable was not found.", k))
		}

		variables[k] = env
	}
}
