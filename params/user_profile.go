package params

type ProfileRequest struct {
	UserID string `json:"user_id"`
}

type ProfileResponse struct {
	Name string `json:"name"`
}
