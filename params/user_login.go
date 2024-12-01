package params

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User   UserInfo `json:"user"`
	Tokens Tokens   `json:"tokens"`
}
