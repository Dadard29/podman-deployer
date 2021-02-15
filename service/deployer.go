package service

import (
	"encoding/json"
	"fmt"
	"github.com/Dadard29/podman-deployer/models"
	"github.com/Dadard29/podman-deployer/podman"
	"io/ioutil"
	"log"
	"net/http"
)

type deployParameter struct {
	ImageName string
	ContainerName string
	PodName string
	Volume string
}

func deployRoute(w http.ResponseWriter, r *http.Request) {
	// check auth
	token := r.Header.Get("Authorization")
	if token != globalService.config.getToken() {
		sendResponse(w, "invalid token", http.StatusUnauthorized)
		return
	}


	// deserialize body
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		sendResponse(w, "invalid body", http.StatusBadRequest)
		return
	}

	var params deployParameter
	if err := json.Unmarshal(data, &params); err != nil {
		log.Println(err)
		sendResponse(w, "invalid json", http.StatusBadRequest)
		return
	}

	deploy(params, w)
}

func deploy(params deployParameter, w http.ResponseWriter) {

	var i podman.PodmanExecInterface
	i = podman.PodmanExec{}

	// stop running  container with name
	containers := i.ListRunningContainers()
	for _, c := range containers {
		if c.Names[0] == params.ContainerName {
			// check the image is correct
			if c.Image != params.ImageName {
				sendResponse(
					w, fmt.Sprintf("a running container with this name is using a different image than %s",
						params.ImageName), http.StatusInternalServerError)
				return
			}

			if _, err := i.StopContainer(c); err != nil {
				log.Println(err)
				sendResponse(w, "error stopping existing container", http.StatusInternalServerError)
				return
			}
			log.Println(fmt.Sprintf("container %s (%s) stopped", c.Names, c.Image))
			break
		}
	}

	// delete stopped container with name
	allContainers := i.ListAllContainers()
	for _, c := range allContainers {
		if c.Names[0] == params.ContainerName {
			// check the container is stopped
			if c.State != "exited" {
				sendResponse(w, "container has not been correctly stopped", http.StatusInternalServerError)
				return
			}

			// check the image is correct
			if c.Image != params.ImageName {
				sendResponse(
					w, fmt.Sprintf("a stopped container with this name is using a different image than %s",
						params.ImageName), http.StatusInternalServerError)
				return
			}

			if _, err := i.DeleteContainer(c); err != nil {
				log.Println(err)
				sendResponse(w, "error removing existing container", http.StatusInternalServerError)
				return
			}
			log.Println(fmt.Sprintf("container %s (%s) deleted", c.Names, c.Image))
			break
		}
	}

	// pull the image from name
	log.Println(fmt.Sprintf("pulling image %s", params.ImageName))
	containerImage, err := i.PullImage(params.ImageName)
	if err != nil {
		log.Println(err)
		sendResponse(w, fmt.Sprintf("failed to pull image with name %s: %s", params.ImageName, err.Error()), http.StatusInternalServerError)
		return
	}

	var newContainer models.Container
	log.Println(fmt.Sprintf("starting container with name %s", params.ContainerName))
	if newContainer, err = i.RunContainer(
		params.ContainerName, containerImage, params.Volume, params.PodName); err != nil {
		log.Println(err)
		sendResponse(w, "error starting container", http.StatusInternalServerError)
		return
	}


	msg := fmt.Sprintf("container %s (%s) started", newContainer.Names, newContainer.Image)
	log.Println(msg)
	sendResponse(w, msg, http.StatusOK)

}
