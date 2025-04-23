package gocin7corerestapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const inventory_write_off_list_url_slug = "inventoryWriteOffList"

type InventoryWriteOffListService interface {
	WriteOffList(options map[string]string) ([]InventoryWriteOff, error)
	WriteOffListAll(options map[string]string) ([]InventoryWriteOff, error)
}

type InventoryWriteOffListOp struct {
	client *Client
}

func (s *InventoryWriteOffListOp) WriteOffList(options map[string]string) ([]InventoryWriteOff, error) {
	var inventory_write_off_list []InventoryWriteOff
	endpoint := s.client.BuildCin7Endpoint(inventory_write_off_list_url_slug, options)
	req, err := s.client.BuildCin7Request("GET", endpoint, nil)
	if err != nil {
		return inventory_write_off_list, err
	}
	cin7_resp, err := s.ReturnPaginatedCin7Data(req)
	if err != nil {
		return inventory_write_off_list, err
	}
	inventory_write_off_list = cin7_resp.InventoryWriteOffs
	return inventory_write_off_list, nil
}

func (s *InventoryWriteOffListOp) WriteOffListAll(options map[string]string) ([]InventoryWriteOff, error) {
	var inventory_write_off_list []InventoryWriteOff
	endpoint := s.client.BuildCin7Endpoint(inventory_write_off_list_url_slug, options)
	req, err := s.client.BuildCin7Request("GET", endpoint, nil)
	if err != nil {
		return inventory_write_off_list, err
	}
	cin7_resp, err := s.ReturnPaginatedCin7Data(req)
	if err != nil {
		return inventory_write_off_list, err
	}
	inventory_write_off_list = cin7_resp.InventoryWriteOffs

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
			return inventory_write_off_list, nil
		} else {
			// set options to next page for request
			options["Page"] = strconv.Itoa(cin7_resp.Page + 1)
		}

		// do the request loop again
		endpoint := s.client.BuildCin7Endpoint(inventory_write_off_list_url_slug, options)
		req, err := s.client.BuildCin7Request("GET", endpoint, nil)
		if err != nil {
			return nil, err
		}
		cin7_resp, err = s.ReturnPaginatedCin7Data(req)
		if err != nil {
			return nil, err
		}
		// append our list of paginated values
		inventory_write_off_list = append(inventory_write_off_list, cin7_resp.InventoryWriteOffs...)

		if count > 4 {
			return inventory_write_off_list, fmt.Errorf("Pagination is greater than 4, error sent to communicate that no more runs have been sent. Stopped to prevent to many requests. Data collected so far is still returned")
		}
	}
}

func (s *InventoryWriteOffListOp) ReturnPaginatedCin7Data(req *http.Request) (*InventoryWriteOffList, error) {
	var paginated_response *InventoryWriteOffList
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
