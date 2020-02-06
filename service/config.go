package service

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	SecretKey = "PODMAN_DEPLOYER_SECRET"
	HostKey = "PODMAN_DEPLOYER_HOST"
	PortKey = "PODMAN_DEPLOYER_PORT"
)

type config struct {
	host string
	port int
}

func (c config) getToken() string {
	return c.generateToken(os.Getenv(SecretKey))
}

func (c config) checkSecret() bool {
	return os.Getenv(SecretKey) != ""
}


func (c config) generateToken(secret string) string {
	hash := sha256.New()
	hash.Write([]byte(secret))

	return hex.EncodeToString(hash.Sum(nil))
}

func getDefaultConfig() config {
	defaultHost, err := os.Hostname()
	if err != nil {
		log.Fatalln(err)
	}

	defaultPort := 8080

	return config{
		host: defaultHost,
		port: defaultPort,
	}
}

func retrieveConfig() config {
	// check secret key
	c := getDefaultConfig()
	if ! c.checkSecret() {
		log.Fatalln(fmt.Sprintf("you must fulfilled the env variable %s with a secret !", SecretKey))
	}

	if host := os.Getenv(HostKey); host != "" {
		c.host = host
	}

	if port := os.Getenv(PortKey); port != "" {
		portInt, err := strconv.Atoi(port)
		if err != nil {
			log.Fatalln(fmt.Sprintf("port from variable %s cannot be converted to int value", PortKey))
		}
		c.port = portInt
	}

	return c
}
