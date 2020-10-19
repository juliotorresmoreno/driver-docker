package container

import (
	"github.com/juliotorresmoreno/drivers/docker/client"
)

// Client .
type Client struct {
	client *client.Client
}

// NewClient .
func NewClient(client *client.Client) *Client {
	cli := &Client{}
	cli.client = client

	return cli
}

// Attach : Attach local standard input, output, and error streams to a running container
func (el *Client) Attach() {

}

// Commit : Create a new image from a container's changes
func (el *Client) Commit() {

}

// Cp : Copy files/folders between a container and the local filesystem
func (el *Client) Cp() {

}

// Create : Create a new container
func (el *Client) Create() {

}

// Diff : Inspect changes to files or directories on a container's filesystem
func (el *Client) Diff() {

}

// Exec : Run a command in a running container
func (el *Client) Exec() {

}

// Export : Export a container's filesystem as a tar archive
func (el *Client) Export() {

}

// Inspect : Display detailed information on one or more containers
func (el *Client) Inspect() {

}

// Kill : Kill one or more running containers
func (el *Client) Kill() {

}

// Logs : Fetch the logs of a container
func (el *Client) Logs() {

}

// Pause : Pause all processes within one or more containers
func (el *Client) Pause() {

}

// Port : List port mappings or a specific mapping for the container
func (el *Client) Port() {

}

// Prune : Remove all stopped containers
func (el *Client) Prune() {

}

// Rename : Rename a container
func (el *Client) Rename() {

}

// Restart : Restart one or more containers
func (el *Client) Restart() {

}

// Rm : Remove one or more containers
func (el *Client) Rm() {

}

// Run : Run a command in a new container
func (el *Client) Run() {

}

// Stats : Display a live stream of container(s) resource usage statistics
func (el *Client) Stats() {

}

// Stop one or more running containers
func (el *Client) Stop() {

}

// Top : Display the running processes of a container
func (el *Client) Top() {

}

// Unpause : Unpause all processes within one or more containers
func (el *Client) Unpause() {

}

// Update : Update configuration of one or more containers
func (el *Client) Update() {

}

// Wait : Block until one or more containers stop, then print their exit codes
func (el *Client) Wait() {

}
