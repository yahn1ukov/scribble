package dto

type UploadInput struct {
	Name        string
	Size        int64
	ContentType string
	Content     []byte
}
