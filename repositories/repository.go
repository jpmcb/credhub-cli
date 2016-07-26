package repositories

import (
	"net/http"

	"encoding/json"
	"errors"

	"github.com/pivotal-cf/cm-cli/client"
	cm_errors "github.com/pivotal-cf/cm-cli/errors"
	"github.com/pivotal-cf/cm-cli/models"
)

type Repository interface {
	SendRequest(request *http.Request, identifier string) (models.Item, error)
}

func DoSendRequest(httpClient client.HttpClient, request *http.Request) (*http.Response, error) {
	response, err := httpClient.Do(request)

	if err != nil {
		return nil, cm_errors.NewNetworkError()
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		decoder := json.NewDecoder(response.Body)
		serverError := models.ServerError{}
		err = decoder.Decode(&serverError)
		if err != nil {
			return nil, err
		}

		if response.StatusCode == http.StatusUnauthorized {
			return nil, cm_errors.NewUnauthorizedError()
		}

		return nil, errors.New(serverError.Error)
	}
	return response, nil
}
