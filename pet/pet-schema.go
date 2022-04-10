package pet

type Pet struct {
	_id     interface{} `json:"_id"`
	Name    string      `json:"name"`
	Class   string      `json:"class"`
	OwnerID string      `json:"ownerID"`
}
