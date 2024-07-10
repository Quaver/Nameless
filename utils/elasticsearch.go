package utils

import (
	"fmt"
	"github.com/Swan/Nameless/config"
	"github.com/imroc/req"
	"net/http"
)

// UpdateElasticSearchMapset Tells the API to update a mapset in elasticsearch
func UpdateElasticSearchMapset(id int) error {
	if err := updateElasticMapsetV1(id); err != nil {
		return err
	}

	if err := updateElasticMapsetV2(id); err != nil {
		return err
	}

	return nil
}

// Updates elastic search for api-v1
func updateElasticMapsetV1(id int) error {
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

// Updates elastic search for api-v2
func updateElasticMapsetV2(id int) error {
	header := req.Header{
		"Authorization": fmt.Sprintf("Bearer %v", config.Data.APIBotJWT),
	}

	r, err := req.Get(fmt.Sprintf("%v/v2/mapset/%v/elastic", config.Data.APIBaseUrl, id), header)

	if err != nil {
		return err
	}

	response := r.Response()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("elasticsearch update failed with status - %v", response.StatusCode)
	}

	return nil
}
