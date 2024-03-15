package gocin7corerestapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const product_availability_url_slug = "ref/productavailability"

type ProductAvailabilityService interface {
	List(options map[string]string) ([]ProductAvailability, error)
	ListAll(options map[string]string) ([]ProductAvailability, error)
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

func (s *ProductAvailabilityOp) ListAll(options map[string]string) ([]ProductAvailability, error) {
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
	product_availability = cin7_resp.ProductAvailabilityList // Block to loop through requests and get paginated data

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
			return product_availability, nil
		} else {
			// set options to next page for request
			options["Page"] = strconv.Itoa(cin7_resp.Page + 1)
		}

		// do the request loop again
		endpoint := s.client.BuildCin7Endpoint(product_availability_url_slug, options)
		req, err := s.client.BuildCin7Request("GET", endpoint, nil)
		if err != nil {
			return nil, err
		}
		cin7_resp, err = s.ReturnPaginatedCin7Data(req)
		if err != nil {
			return nil, err
		}
		// append our list of paginated values
		product_availability = append(product_availability, cin7_resp.ProductAvailabilityList...)

		if count > 4 {
			return nil, fmt.Errorf("Shit got round too many times during testing, stop battering API")
		}
	}
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
