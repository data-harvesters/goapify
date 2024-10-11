package internal

var (
	dockerFileTemplate = `FROM golang:1.22.1

WORKDIR /app
COPY . .

RUN go mod download

RUN go build -o ./${name}
CMD ["./${name}"]
`

	actorJsonTemplate = `{
    "actorSpecification": 1,
    "name": "${name}",
    "title": "${name}",
    "version": "1.0.0"
}`

	inputSchemaTemplate = `{
    "title": "${name} Input",
    "type": "object",
    "schemaVersion": 1,
    "properties": {
    },
    "required": [
    ]
}`

	actorTemplate = `package main

    import (
	"fmt"
	"os"

	"github.com/data-harvesters/goapify"
)

type input struct {
	*goapify.ProxyConfigurationOptions 'json:"proxyConfiguration"'
}

func main() {
	a := goapify.NewActor(
		os.Getenv("APIFY_DEFAULT_KEY_VALUE_STORE_ID"),
		os.Getenv("APIFY_TOKEN"),
		os.Getenv("APIFY_DEFAULT_DATASET_ID"),
	)

	i := new(input)

	err := a.Input(i)
	if err != nil {
		fmt.Printf("failed to decode input: %v\\n", err)
		panic(err)
	}

	if i.ProxyConfigurationOptions != nil {
		err = a.CreateProxyConfiguration(i.ProxyConfigurationOptions)
		if err != nil {
			panic(err)
		}
	}
}`
)
