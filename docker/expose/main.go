package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/juliotorresmoreno/drivers/docker/client"
	"github.com/juliotorresmoreno/drivers/docker/container"
)

func main() {
	cli := client.NewClient([]string{})
	containersCli := container.NewClient(cli)
	containers, _ := containersCli.Ls()
	fmt.Println("mierda")
	json.NewEncoder(os.Stdout).Encode(containers)
}
