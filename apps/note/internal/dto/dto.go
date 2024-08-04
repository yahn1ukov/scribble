package dto

type CreateInput struct {
	Title   string
	Content *string
}

type UpdateInput struct {
	Title   *string
	Content *string
}
