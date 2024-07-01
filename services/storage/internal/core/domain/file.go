package domain

import "io"

type File struct {
	URL         string
	Size        int64
	ContentType string
	Content     io.Reader
}
