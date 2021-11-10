package httpClient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"server/service/mylib/errorCode"
	"time"
)

type Client struct {
	c *http.Client
}

var (
	netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 10 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 10 * time.Second,
	}
)

const (
	ContentType = "Content-Type"
	JSON        = "application/json"
)

//取得新的http client
func NewClient() (c *Client) {
	return &Client{
		c: &http.Client{
			Timeout:   10 * time.Second,
			Transport: netTransport,
		},
	}
}

func (client *Client) PostJson(url string, data interface{}) (code int, request *http.Request, err error) {
	para, err := json.Marshal(data)
	if err != nil {
		code = errorCode.EncodeJsonError
		return
	}
	request, err = http.NewRequest(http.MethodPost, url, bytes.NewBuffer(para))
	if err != nil {
		code = errorCode.RequestCreateError
		return
	}
	request.Header.Set(ContentType, JSON)
	return
}

func (client *Client) Send(request *http.Request) (code int, result []byte, err error) {
	res, err := client.c.Do(request)
	if err != nil {
		code = errorCode.RequestSendError
		return
	}
	defer res.Body.Close()
	result, err = ioutil.ReadAll(res.Body)
	if err != nil {
		code = errorCode.ResponseDecodeError
		return
	}
	return
}
