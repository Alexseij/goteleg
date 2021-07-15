package goteleg

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Type for working with telegram API
type API string

// https://api.telegram.org/bot<token>/METHOD_NAME
func newApi(token string) API {
	return API(fmt.Sprintf("https://api.telegram.org/bot%s/", token))
}

func sendGetQuery(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (a API) SendMessage(text string, chatID int64, options *MessageOption) (MessageResponce, error) {

	var messageResponce MessageResponce

	url := fmt.Sprintf(
		"%ssendMessage?text=%s&chat_id=%d&%s",
		string(a),
		url.QueryEscape(text),
		chatID,
		parseOptions(options),
	)

	jsonResp, err := sendGetQuery(url)
	if err != nil {
		return messageResponce, err
	}

	json.Unmarshal(jsonResp, &messageResponce)

	return messageResponce, nil

}

func (a API) ReplyToMessage(text string, chatID int64, messageID int) (MessageResponce, error) {
	var messageResponce MessageResponce

	url := fmt.Sprintf(
		"%ssendMessage?text=%s&chat_id=%d&reply_to_message_id=%d",
		string(a),
		url.QueryEscape(text),
		chatID,
		messageID,
	)

	jsonResp, err := sendGetQuery(url)
	if err != nil {
		return messageResponce, err
	}

	json.Unmarshal(jsonResp, &messageResponce)

	return messageResponce, nil
}

func (a API) GetUpdates(offset int, timeout int) (UpdatesResponce, error) {
	var updatesResponce UpdatesResponce

	url := fmt.Sprintf(
		"%sgetUpdates?timeout=%d&offset=%d",
		string(a),
		timeout,
		offset,
	)

	jsonResp, err := sendGetQuery(url)
	if err != nil {
		return updatesResponce, err
	}

	json.Unmarshal(jsonResp, &updatesResponce)

	return updatesResponce, nil
}
