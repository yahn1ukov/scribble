// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gqlmodels

import (
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
)

type AuthOutput struct {
	Token string `json:"token"`
}

type CreateNoteInput struct {
	Title   string            `json:"title"`
	Content *string           `json:"content,omitempty"`
	Files   []*graphql.Upload `json:"files,omitempty"`
}

type CreateNotebookInput struct {
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
}

type File struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Size        int64     `json:"size"`
	ContentType string    `json:"contentType"`
	CreatedAt   time.Time `json:"createdAt"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Mutation struct {
}

type Note struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Content   *string   `json:"content,omitempty"`
	Files     []*File   `json:"files"`
	CreatedAt time.Time `json:"createdAt"`
}

type Notebook struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
}

type Query struct {
}

type RegisterInput struct {
	Email     string  `json:"email"`
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
	Password  string  `json:"password"`
}

type UpdateNoteInput struct {
	Title   *string `json:"title,omitempty"`
	Content *string `json:"content,omitempty"`
}

type UpdateNotebookInput struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
}

type UpdateUserInput struct {
	Email     *string `json:"email,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
}

type UpdateUserPasswordInput struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	FirstName *string   `json:"firstName,omitempty"`
	LastName  *string   `json:"lastName,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
}
