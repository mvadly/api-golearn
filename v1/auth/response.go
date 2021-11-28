package auth

type ResponseLogin struct {
	ID       uint32 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type ResponseProfile struct {
	ID       uint32 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}
