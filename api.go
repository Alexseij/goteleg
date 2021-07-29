package goteleg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Type for working with telegram API
type API struct {
	apiURL string
	client *http.Client
}

// https://api.telegram.org/bot<token>/METHOD_NAME
func newApi(token string, client *http.Client) API {
	return API{
		apiURL: fmt.Sprintf("https://api.telegram.org/bot%s/", token),
		client: client,
	}
}

func getData(url string, params interface{}, client *http.Client) ([]byte, error) {

	jsonParams, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(jsonParams)

	fmt.Println()
	buf := bytes.NewBuffer(json)

	resp, err := client.Post(url, "aplictation/json", buf)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("goteleg : http request has status : %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (a API) SendMessage(text string, chatID int64, options *MessageOption) (MessageResponce, error) {

	var messageResponce MessageResponce

	url := fmt.Sprintf(
		"%ssendMessage?chat_id=%d&text=%s",
		a.apiURL,
		chatID,
		url.QueryEscape(text),
	)

	data, err := getData(url, options, a.client)
	if err != nil {
		return messageResponce, err
	}

	json.Unmarshal(data, &messageResponce)

	return messageResponce, nil

}

func (a API) ReplyToMessage(optinons interface{}) (MessageResponce, error) {
	var messageResponce MessageResponce

	url := fmt.Sprintf(
		"%ssendMessage",
		a.apiURL,
	)

	data, err := getData(url, optinons, a.client)
	if err != nil {
		return messageResponce, err
	}

	json.Unmarshal(data, &messageResponce)

	return messageResponce, nil
}

func (a API) GetUpdates(options interface{}) (UpdatesResponce, error) {
	var updatesResponce UpdatesResponce

	url := fmt.Sprintf(
		"%sgetUpdates",
		a.apiURL,
	)

	data, err := getData(url, options, a.client)
	if err != nil {
		return updatesResponce, err
	}

	json.Unmarshal(data, &updatesResponce)

	return updatesResponce, nil
}

func (a API) SendPhoto(chatID int, photo string, options *PhotoOption) (MessageResponce, error) {
	var messageResponce MessageResponce

	url := fmt.Sprintf(
		"%ssendPhoto",
		a.apiURL,
	)

	data, err := getData(url, options, a.client)
	if err != nil {
		return messageResponce, err
	}

	json.Unmarshal(data, &messageResponce)

	return messageResponce, nil
}
