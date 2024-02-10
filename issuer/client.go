package issuer

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type ResponseBody struct {
	StatusCode int
	Body       []byte
}

type Client struct {
	BaseURL  string
	Username string
	Password string
}

type IClient interface {
	Do(path string, paramsJSON []byte) (*ResponseBody, error)
}

func (c *Client) Do(path string, paramsJSON []byte) (*ResponseBody, error) {
	endpoint, err := url.JoinPath(c.BaseURL, path)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(paramsJSON))
	if err != nil {
		return nil, err
	}

	fmt.Println("req: ", req)

	req.Header.Set("User-Agent", "dummy-wallet-api(Go)")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("Accept", "application/json, text/plain, */*")

	req.SetBasicAuth(c.Username, c.Password)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return &ResponseBody{
		Body:       body,
		StatusCode: res.StatusCode,
	}, err
}
