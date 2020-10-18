package model

type DockerState struct {
	ImageExists      bool `json:"imageExists"`
	ContainerExists  bool `json:"containerExists"`
	ContainerRunning bool `json:"containerRunning"`
}
