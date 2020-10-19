package container

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Start : Start one or more stopped containers
// a, --attach               Attach STDOUT/STDERR and forward signals
// --checkpoint string       Restore from this checkpoint
// --checkpoint-dir string   Use a custom checkpoint storage directory
// --detach-keys string      Override the key sequence for detaching a container
// -i, --interactive         Attach container's STDIN
func (el *Client) Start(
	attach bool,
	checkpoint string,
	checkpointDir string,
	detachKeys string,
	interactive string,
) error {
	body := map[string]interface{}{}
	b, err := json.Marshal(body)
	if err != nil {
		return err
	}
	node := el.selectBestNode()
	buffer := bytes.NewBuffer(b)
	url := fmt.Sprintf("%v", node)
	request, err := http.NewRequest("POST", url, buffer)
	if err != nil {
		return err
	}
	response, error := http.DefaultClient.Do(request)
	contentData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	content := &map[string]interface{}{}
	json.Unmarshal(contentData, content)
	fmt.Println(content)
	return error
}
