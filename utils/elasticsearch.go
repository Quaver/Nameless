package utils

import (
	"fmt"
	config "github.com/Swan/Nameless/config"
	"github.com/imroc/req"
	"net/http"
)

// UpdateElasticSearchMapset Tells the API to update a mapset in elasticsearch
func UpdateElasticSearchMapset(id int) error {
	header := req.Header{"Accept": "application/json"}
	param := req.Param{
		"key": config.Data.APISecretKey,
		"id":  id,
	}
	
	r, err := req.Post(fmt.Sprintf("%v/v1/mapsets/elastic", config.Data.APIBaseUrl), header, param)
	
	if err != nil {
		return err
	}
	
	response := r.Response()
	
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("elasticsearch update failed with status - %v", response.StatusCode)
	}
	
	return nil
}
