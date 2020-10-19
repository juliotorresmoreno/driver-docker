package manager

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Manager .
type Manager struct {
}

// Mutate .
func (el *Manager) Mutate(
	method string,
	url string,
	headers http.Header,
	body interface{},
) (*http.Response, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return &http.Response{}, err
	}
	buffer := bytes.NewBuffer(b)
	request, err := http.NewRequest(method, url, buffer)
	return http.DefaultClient.Do(request)
}

// Query .
func (el *Manager) Query() {

}
