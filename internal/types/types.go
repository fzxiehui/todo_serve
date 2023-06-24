package types

/*
 * Auth
 */
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Token    string `json:"token"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

type RegisterResponse struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Token    string `json:"token"`
}
