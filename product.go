package gocin7corerestapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const product_url_slug = "product"

type ProductService interface {
	List(options map[string]string) ([]Product, error)
	ListAll(options map[string]string) ([]Product, error)
}

type ProductOp struct {
	client *Client
}

func (s *ProductOp) List(options map[string]string) ([]Product, error) {
	var products []Product
	endpoint := s.client.BuildCin7Endpoint(product_url_slug, options)
	req, err := s.client.BuildCin7Request("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	cin7_resp, err := s.ReturnPaginatedCin7Data(req)
	if err != nil {
		return nil, err
	}
	products = cin7_resp.Products
	return products, nil
}

func (s *ProductOp) ListAll(options map[string]string) ([]Product, error) {
	var products []Product
	endpoint := s.client.BuildCin7Endpoint(product_url_slug, options)
	req, err := s.client.BuildCin7Request("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	cin7_resp, err := s.ReturnPaginatedCin7Data(req)
	if err != nil {
		return nil, err
	}
	products = cin7_resp.Products // Block to loop through requests and get paginated data

	// Check if limit is explicit, set to default for cin7 otherwise
	if _, ok := options["Limit"]; !ok {
		options["Limit"] = "100"
	}

	// limit of loops here in case  something bad happens
	count := 0
	for {
		count += 1
		// calculate how many records are left
		set_limit, err := strconv.ParseInt(options["Limit"], 0, 64)
		if err != nil {
			return nil, err
		}
		remainder := cin7_resp.Total - (int(set_limit) * cin7_resp.Page)
		if remainder < 1 {
			// got all records
			return products, nil
		} else {
			// set options to next page for request
			options["Page"] = strconv.Itoa(cin7_resp.Page + 1)
		}

		// do the request loop again
		endpoint := s.client.BuildCin7Endpoint(product_url_slug, options)
		req, err := s.client.BuildCin7Request("GET", endpoint, nil)
		if err != nil {
			return nil, err
		}
		cin7_resp, err = s.ReturnPaginatedCin7Data(req)
		if err != nil {
			return nil, err
		}
		// append our list of paginated values
		products = append(products, cin7_resp.Products...)

		if count > 4 {
			return products, fmt.Errorf("Pagination is greater than 4, error sent to communicate that no more runs have been sent. Stopped to prevent to many requests. Data collected so far is still returned")
		}
	}
}

func (s *ProductOp) ReturnPaginatedCin7Data(req *http.Request) (*ProductPaginatedResponse, error) {
	var paginated_response *ProductPaginatedResponse
	resp, err := s.client.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error encountered during request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.Request)
		return nil, fmt.Errorf("Request Error: %v", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&paginated_response)
	if err != nil {
		return nil, fmt.Errorf("Error encountered while decoding in ReturnPaginatedCin7Data: %v", err)
	}
	return paginated_response, nil
}
