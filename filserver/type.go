package filserver

type readBody struct {
	File string `json:"file" binding:"required"`
}
type createBody struct {
	Name    string `json:"name" binding:"required"`
	Content string `json:"content" binding:"required"`
}
