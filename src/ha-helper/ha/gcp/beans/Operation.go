package beans

type Operation struct {
	Kind          string `json:"kind"`
	ID            string `json:"id"`
	Name          string `json:"name"`
	Zone          string `json:"zone"`
	OperationType string `json:"operationType"`
	TargetLink    string `json:"targetLink"`
	TargetID      string `json:"targetId"`
	Status        string `json:"status"`
	User          string `json:"user"`
	Progress      int    `json:"progress"`
	InsertTime    string `json:"insertTime"`
	SelfLink      string `json:"selfLink"`
}
