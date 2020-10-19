package container

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Port struct {
	IP          string
	PrivatePort int
	PublicPort  int
	Type        string
}

type Bridge struct {
	IPAMConfig          string
	Links               string
	Aliases             string
	NetworkID           string
	EndpointID          string
	Gateway             string
	IPAddress           string
	IPPrefixLen         string
	IPv6Gateway         string
	GlobalIPv6Address   string
	GlobalIPv6PrefixLen int
	MacAddress          string
	DriverOpts          interface{}
}

type Networks struct {
	Bridge Bridge `json:"bridge"`
}

type NetworkSettings struct {
	Networks Networks
}

type HostConfig struct {
	NetworkMode string
}

type Container struct {
	ID              string `json:"Id"`
	Names           []string
	Image           string
	ImageID         string
	Command         string
	Created         time.Time
	Ports           []Port `json:"Ports"`
	Labels          interface{}
	State           string
	Status          string
	HostConfig      HostConfig
	NetworkSettings NetworkSettings
	Mounts          []interface{}
}

// Containers .
type Containers []Container

func (el *Container) UnmarshalJSON(data []byte) error {
	el.ID = "12"
	return nil
}

func (el Container) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{"na": "ss"})
}

// Ls : List containers
func (el *Client) Ls() ([]Container, error) {
	result := make(Containers, 0)
	body := map[string]interface{}{}
	b, err := json.Marshal(body)
	if err != nil {
		return result, err
	}
	node := el.selectBestNode()
	buffer := bytes.NewBuffer(b)
	url := fmt.Sprintf("%v/containers/json", node)
	request, err := http.NewRequest("GET", url, buffer)
	if err != nil {
		return result, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return result, err
	}
	contentData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(contentData, &result)
	return result, err
}
