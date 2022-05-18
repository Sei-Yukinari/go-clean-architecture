package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-clean-architecture/src/infrastructure/logger"
	"io/ioutil"
	"net/http"
	"time"
)

type Client interface {
	Get(endpoint string) (*http.Request, error)
	PostWith(endpoint string, params interface{}) (*http.Request, error)
	Do(request *http.Request) (ClientResponse, error)
}

type ResponseStruct struct {
	Status        string
	StatusCode    int
	Header        http.Header
	ContentLength int64
	Body          []byte
}

type ClientResponse interface {
	Get() ResponseStruct
	To(value interface{})
}

type client struct {
	BaseUrl string
}

func New(baseUrl string) Client {
	return &client{BaseUrl: baseUrl}
}

func (h client) Get(endpoint string) (*http.Request, error) {
	return http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s%s", h.BaseUrl, endpoint),
		bytes.NewBuffer([]byte{}),
	)
}

func (h client) PostWith(endpoint string, params interface{}) (*http.Request, error) {
	b, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	return http.NewRequest(
		http.MethodPost,
		h.BaseUrl+endpoint,
		bytes.NewBuffer(b),
	)
}

func (h client) Do(request *http.Request) (ClientResponse, error) {
	client := &http.Client{}
	t := time.Now()
	response, err := client.Do(request)
	logger.Infof("response time(%s): %+vms elapsed", request.URL, time.Since(t).Milliseconds())
	if err != nil {
		return nil, err
	}
	defer func() {
		err := response.Body.Close()
		if err != nil {
			logger.Warnf("response body close error:%v\n", err)
		}
	}()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return &ResponseStruct{
		Status:        response.Status,
		StatusCode:    response.StatusCode,
		Header:        response.Header,
		ContentLength: response.ContentLength,
		Body:          body,
	}, nil
}

func (r ResponseStruct) Get() ResponseStruct {
	return r
}

func (r ResponseStruct) To(value interface{}) {
	err := json.Unmarshal(r.Body, &value)
	if err != nil {
		value = nil
	}
}
