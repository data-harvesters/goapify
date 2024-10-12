package goapify

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

var (
	InputNotFoundError = errors.New("failed to find input")
)

type Actor struct {
	ctx    context.Context
	cancel context.CancelFunc

	ID     string
	TaskID string

	apifyUserID string

	key       string
	token     string
	datasetId string

	input map[string]any

	ProxyConfiguration *ProxyConfiguration

	client *http.Client
}

func NewActor() *Actor {
	ctx, cancel := context.WithCancel(context.Background())
	ensureEnvironment()

	return &Actor{
		ctx:         ctx,
		cancel:      cancel,
		ID:          variables["ACTOR_ID"],
		TaskID:      variables["ACTOR_TASK_ID"],
		apifyUserID: variables["APIFY_USER_ID"],
		key:         variables["APIFY_DEFAULT_KEY_VALUE_STORE_ID"],
		token:       variables["APIFY_TOKEN"],
		datasetId:   variables["APIFY_DEFAULT_DATASET_ID"],
		input:       make(map[string]any),
		client: &http.Client{
			Timeout: 1 * time.Minute,
		},
	}
}

func (a *Actor) Context() context.Context {
	return a.ctx
}

func (a *Actor) Exit() {
	a.cancel() // cancel actor context
}

func (a *Actor) GetInput(key string) (any, error) {
	if v, ok := a.input[key]; ok {
		return v, nil
	}

	return nil, InputNotFoundError
}

func (a *Actor) Input(payload any) error {
	url := fmt.Sprintf("https://api.apify.com/v2/key-value-stores/%s/records/INPUT?token=%s", a.key, a.token)

	req, err := http.NewRequestWithContext(a.ctx, "GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := a.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &payload)
	if err != nil {
		return err
	}

	var p map[string]any
	err = json.Unmarshal(b, &p)
	if err != nil {
		return err
	}
	a.input = p

	return nil
}

func (a *Actor) Output(payload any) error {
	url := fmt.Sprintf("https://api.apify.com/v2/datasets/%s/items?token=%s", a.datasetId, a.token)

	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(a.ctx, "POST", url, bytes.NewReader(data))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := a.client.Do(req)
	if err != nil {
		return err
	}

	dd, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		return fmt.Errorf("unable to add dataset: %s", string(dd))
	}

	return nil
}

func (a *Actor) CreateProxyConfiguration(proxyOptions *ProxyConfigurationOptions) error {

	if proxyOptions.UseApifyProxy {
		ensureProxyEnvironment()
	}

	if !proxyOptions.UseApifyProxy {
		if proxyOptions.ProxyUrls == nil {
			return nil
		}
	}

	a.ProxyConfiguration = newProxyConfiguration(proxyOptions)

	return nil
}
