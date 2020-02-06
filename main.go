package main

import (
	"github.com/Dadard29/podman-deployer/service"
)

func main() {
	s := service.NewService()
	s.Run()
}
