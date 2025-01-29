package gocin7corerestapi

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Client struct {
	baseUrl             string
	httpClient          *http.Client
	ProductAvailablilty ProductAvailabilityService
	Product             ProductService
}

/*
Endgoal is to have a function like:
import cin7 "example.com/gocin7corerestapi"
os.Setenv("api_auth_key", "blah")
os.Setenv("api_account_id", "blah")

gocin7 := cin7.NewClient()
products, err := gocin7.ProductAvailability.List("some options here")
*/

func NewClient(url string, httpClient *http.Client) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	if url == "" {
		url = "https://inventory.dearsystems.com/ExternalApi/v2/"
	}
	c := &Client{
		baseUrl:    url,
		httpClient: httpClient,
	}
	c.ProductAvailablilty = &ProductAvailabilityOp{client: c}
	c.Product = &ProductOp{client: c}
	return c, nil
}

func (c *Client) BuildCin7Request(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("Error while building request: %v", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("api-auth-accountid", os.Getenv("CIN7_ACCOUNT_ID"))
	req.Header.Add("api-auth-applicationkey", os.Getenv("CIN7_AUTH_KEY"))
	return req, nil
}

func BuildCin7QueryParams(options map[string]string) string {
	queryParams := url.Values{}
	for k, v := range options {
		queryParams.Add(k, v)
	}
	params := queryParams.Encode()
	params = strings.ReplaceAll(params, "+", "%20")
	return params
}

func (c *Client) BuildCin7Endpoint(slug string, options map[string]string) string {
	var queryParams string
	if options != nil {
		queryParams = "?" + BuildCin7QueryParams(options)
	}
	endpoint := c.baseUrl + slug + queryParams
	return endpoint
}
