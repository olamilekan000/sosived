package helpers

import (
	"fmt"
	"io"
	"net/http"
)

type ServiceCaller struct{}

func (s *ServiceCaller) MakeServiceRequest(authToken string, method string, url string, body io.Reader) (*http.Response, error) {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, body)

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", authToken))

	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
