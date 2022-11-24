package dto

type Todo struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type GetAllTodosResponse struct {
	Data  []Todo `json:"data"`
	Total int    `json:"total"`
}
