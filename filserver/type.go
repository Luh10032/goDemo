package filserver

const FileDir = "file/"

type CreateBody struct {
	Name    string `json:"name" binding:"required"`
	Content string `json:"content" binding:"required"`
}
