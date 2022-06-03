package interfaces

import (
	"io"
	"net/http"
)

type ServiceCaller interface {
	MakeServiceRequest(authToken string, method string, url string, body io.Reader) (*http.Response, error)
}
