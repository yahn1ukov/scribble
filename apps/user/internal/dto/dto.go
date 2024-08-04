package dto

type CreateInput struct {
	Email     string
	FirstName *string
	LastName  *string
	Password  string
}

type UpdateInput struct {
	Email     *string
	FirstName *string
	LastName  *string
}

type UpdatePasswordInput struct {
	OldPassword string
	NewPassword string
}
