package dto

type CreateInput struct {
	Title       string
	Description *string
}

type UpdateInput struct {
	Title       *string
	Description *string
}
