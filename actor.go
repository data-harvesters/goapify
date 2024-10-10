package goapify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Actor struct {
	Key   string
	Token string

	DatasetId string

	client *http.Client
}

func NewActor(key, token, datasetID string) *Actor {
	return &Actor{
		Key:       key,
		Token:     token,
		DatasetId: datasetID,
		client: &http.Client{
			Timeout: 1 * time.Minute,
		},
	}
}

func (a *Actor) Input(payload any) error {
	url := fmt.Sprintf("https://api.apify.com/v2/key-value-stores/%s/records/INPUT?token=%s", a.Key, a.Token)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := a.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return err
	}

	return nil
}

func (a *Actor) Output(payload any) error {
	url := fmt.Sprintf("https://api.apify.com/v2/datasets/%s/items?token=%s", a.DatasetId, a.Token)

	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
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
