package podman

import (
	"fmt"
	"testing"
)

var podmanExec = PodmanExec{}

func TestPodmanExec_ListRunningContainers(t *testing.T) {
	fmt.Println(podmanExec.ListRunningContainers())
}

func TestPodmanExec_ListImages(t *testing.T) {
	fmt.Println(podmanExec.ListImages())
}

func TestPodmanExec_ListAllContainers(t *testing.T) {
	fmt.Println(podmanExec.ListAllContainers())
}

func TestPodmanExec_PullImage(t *testing.T) {

	if _, err := podmanExec.PullImage("docker.io/library/nginx:latest"); err != nil {
		t.Error(err)
	}
}