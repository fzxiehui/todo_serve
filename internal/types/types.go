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

type CreateTodoRequest struct {
	Content string `json:"content"`
	Done    bool   `json:"done"`
}

type CreateTodoResponse struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
	Date    string `json:"date"`
	Done    bool   `json:"done"`
}

type GetTodoResponse struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
	Date    string `json:"date"`
	Done    bool   `json:"done"`
}

type UpdateTodoRequest struct {
	Done bool `json:"done"`
}

type UpdateTodoResponse struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
	Date    string `json:"date"`
	Done    bool   `json:"done"`
}

type QueryTodoRequest struct {
	Text string `json:"text"`
}

type QueryTodoResponse struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
	Date    string `json:"date"`
	Done    bool   `json:"done"`
}

type QueryTodoListResponse struct {
	Total int                 `json:"total"`
	List  []QueryTodoResponse `json:"list"`
}
