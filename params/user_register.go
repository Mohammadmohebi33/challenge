package params

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"passowrd"`
}

type RegisterResponse struct {
	User UserInfo `json:"user"`
}
