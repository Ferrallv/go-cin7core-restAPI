package gocin7corerestapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const product_availability_url_slug = "ref/productavailability"

type ProductAvailabilityService interface {
	List(options map[string]string) ([]ProductAvailability, error)
}

type ProductAvailabilityOp struct {
	client *Client
}

func (s *ProductAvailabilityOp) List(options map[string]string) ([]ProductAvailability, error) {
	var product_availability []ProductAvailability
	endpoint := s.client.BuildCin7Endpoint(product_availability_url_slug, options)
	req, err := s.client.BuildCin7Request("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	cin7_resp, err := s.ReturnPaginatedCin7Data(req)
	if err != nil {
		return nil, err
	}
	product_availability = cin7_resp.ProductAvailabilityList
	return product_availability, nil
}

func (s *ProductAvailabilityOp) ReturnPaginatedCin7Data(req *http.Request) (*ProductAvailablityPaginatedResponse, error) {
	var paginated_response *ProductAvailablityPaginatedResponse
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
