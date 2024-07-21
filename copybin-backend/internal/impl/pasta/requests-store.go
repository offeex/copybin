package pasta

type CreateRequest struct {
	text string
}

type UpdateRequest struct {
	newText string
	pastaId int64
}
