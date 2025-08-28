package authentication

type LoginRequest struct {
	Id       string   `json:"_id"`
	Name     string   `json:"name"`
	Password string   `json:"password"`
	TypeKey  string   `json:"type,omitempty"`
	Roles    []string `json:"roles,omitempty"`
	Date     string   `json:"date,omitempty"`
}

type LoginResponse struct {
	Ok    string `json:"ok"`
	Id    string `json:"id"`
	Rev   string `json:"rev"`
	Token string `json:"token"`
}
