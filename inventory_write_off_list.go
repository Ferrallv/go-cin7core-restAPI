package gocin7corerestapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const inventory_write_off_list_url_slug = "ref/inventoryWriteOffList"

type InventoryWriteOffListService interface {
	WriteOffList(options map[string]string) (InventoryWriteOffList, error)
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
