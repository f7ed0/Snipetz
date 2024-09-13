package schema

type ConnectionRequest struct {
	MicroserviceType string `json:"type"`
	URI              string `json:"uri"`
}
