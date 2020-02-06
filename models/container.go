package models

import "time"

type Container struct {
	ID        string      `json:"ID"`
	Image     string      `json:"Image"`
	Command   string      `json:"Command"`
	Created   string      `json:"Created"`
	Ports     string      `json:"Ports"`
	Names     string      `json:"Names"`
	IsInfra   bool        `json:"IsInfra"`
	Status    string      `json:"Status"`
	State     int         `json:"State"`
	Pid       int         `json:"Pid"`
	Size      interface{} `json:"Size"`
	Pod       string      `json:"Pod"`
	PodName   string      `json:"PodName"`
	CreatedAt time.Time   `json:"CreatedAt"`
	ExitedAt  time.Time   `json:"ExitedAt"`
	StartedAt time.Time   `json:"StartedAt"`
	Labels    struct {
	} `json:"Labels"`
	PID    string `json:"PID"`
	Cgroup string `json:"Cgroup"`
	IPC    string `json:"IPC"`
	MNT    string `json:"MNT"`
	NET    string `json:"NET"`
	PIDNS  string `json:"PIDNS"`
	User   string `json:"User"`
	UTS    string `json:"UTS"`
	Mounts string `json:"Mounts"`
}
