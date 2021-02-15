package models

type Container struct {
	ID        string      `json:"Id"`
	Image     string      `json:"Image"`
	ImageID string `json:"ImageID"`
	Command   []string      `json:"Command"`
	Created   int      `json:"Created"`
	Ports     string      `json:"Ports"`
	Names     []string      `json:"Names"`
	IsInfra   bool        `json:"IsInfra"`
	Status    string      `json:"Status"`
	State     string         `json:"State"`
	Pid       int         `json:"Pid"`
	Size      interface{} `json:"Size"`
	Pod       string      `json:"Pod"`
	PodName   string      `json:"PodName"`
	CreatedAt string   `json:"CreatedAt"`
	Exited 	bool `json:"Exited"`
	ExitedAt  int   `json:"ExitedAt"`
	StartedAt int   `json:"StartedAt"`
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
	Mounts []string `json:"Mounts"`
}
