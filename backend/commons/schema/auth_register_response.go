package schema

type AuthRegisterResponse struct {
	Created  bool   `json:"created"`
	JwtToken string `json:"jwt_token,omitempty"`
	Uuid     string `json:"uuid,omitempty"`
}
