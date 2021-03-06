package main

import (
	"github.com/emate/packer-ebsnoami/builder/amazon/ebsnoami"
	"github.com/mitchellh/packer/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(ebsnoami.Builder))
	server.Serve()
}
