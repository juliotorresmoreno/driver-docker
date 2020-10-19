package client

type Client struct {
	Hosts []string
}

func NewClient(hosts []string) *Client {
	client := &Client{}
	client.Hosts = hosts
	return client
}
