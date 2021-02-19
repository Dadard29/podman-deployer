package podman

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	. "github.com/Dadard29/podman-deployer/models"
	"log"
	"os/exec"
	"strings"
)

type PodmanExecInterface interface {
	ListRunningContainers() []Container
	ListAllContainers() []Container
	GetContainer(id string) (Container, error)
	RunContainer(name string, image Image, volume string, podName string) (Container, error)
	StopContainer(container Container) (Container, error)
	DeleteContainer(container Container) (Container, error)
	ListImages() []Image
	GetImage(imageName string) (Image, error)
	PullImage(imageName string) (Image, error)
}

type PodmanExec struct {

}

func (PodmanExec) execCommand(args []string, json bool) ([]byte, error) {
	if json {
		args = append(args, "--format", "json")
	}
	log.Println(fmt.Sprintf("executing podman with args %v", args))
	cmd := exec.Command("podman", args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return []byte{}, err
	}
	if err := cmd.Start(); err != nil {
		return []byte{}, err
	}

	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(stdout); err != nil {
		return []byte{}, err
	}

	if err := cmd.Wait(); err != nil {
		return []byte{}, err
	}

	return buf.Bytes(), nil
}

func (p PodmanExec) ListRunningContainers() []Container {
	var containers []Container
	stdout, err := p.execCommand([]string{"ps"}, true)
	if err != nil {
		log.Println(err)
		return containers
	}

	if err := json.Unmarshal(stdout, &containers); err != nil {
		log.Println(err)
		return containers
	}

	return containers
}

func (p PodmanExec) ListAllContainers() []Container {
	var containers []Container
	stdout, err := p.execCommand([]string{"ps", "-a"}, true)
	if err != nil {
		log.Println(err)
		return containers
	}

	if err := json.Unmarshal(stdout, &containers); err != nil {
		log.Println(err)
		return containers
	}

	return containers
}

func (p PodmanExec) GetContainer(id string) (Container, error) {
	containers := p.ListAllContainers()
	for _, c := range containers {
		if strings.Contains(id, c.ID) {
			return c, nil
		}
	}
	return Container{}, errors.New(fmt.Sprintf("container with id %s not found", id))
}

func (p PodmanExec) RunContainer(name string, image Image, volume string, podName string) (Container, error) {
	var container Container
	cmdList := []string{"run", "-d", "--name", name}
	if volume != "" {
		cmdList = append(cmdList, "-v")
		cmdList = append(cmdList, volume)
	}

	if podName != "" {
		cmdList = append(cmdList, "--pod")
		cmdList = append(cmdList, podName)
	}

	cmdList = append(cmdList, image.Names[0])

	stdout, err := p.execCommand(cmdList, false)
	if err != nil {
		return container, err
	}

	containerId := strings.Trim(string(stdout), "\n")
	c, err := p.GetContainer(containerId)
	if err != nil {
		return c, errors.New("the run command did happen, but the created container cannot be found, what the hell is going on")
	}

	return c, nil
}

func (p PodmanExec) StopContainer(container Container) (Container, error) {
	_, err := p.execCommand([]string{"stop", container.ID}, false)
	if err != nil {
		return container, err
	}

	c, err := p.GetContainer(container.ID)
	if err != nil {
		return c, errors.New("the stop command did happen, but the stopped container cannot be found, what the hell is going on")
	}

	return c, nil
}

func (p PodmanExec) DeleteContainer(container Container) (Container, error) {
	_, err := p.execCommand([]string{"rm", container.ID}, false)
	if err != nil {
		return container, err
	}

	return container, nil
}

func (p PodmanExec) ListImages() []Image {
	var images []Image
	stdout, err := p.execCommand([]string{"images"}, true)
	if err != nil {
		log.Println(err)
	}

	if err := json.Unmarshal(stdout, &images); err != nil {
		log.Println(err)
	}

	return images
}

func (p PodmanExec) GetImage(imageName string) (Image, error) {
	images := p.ListImages()
	for _, i := range images {
		for _, n := range i.Names {
			if n == imageName {
				return i, nil
			}
		}
	}

	return Image{}, errors.New(fmt.Sprintf("image with name %s not found", imageName))
}

func (p PodmanExec) PullImage(imageName string) (Image, error) {
	var i Image
	_, err := p.execCommand([]string{"pull", imageName}, false)
	if err != nil {
		return i, err
	}

	i, err = p.GetImage(imageName)
	if err != nil {
		return i, err
	}

	return i, nil
}