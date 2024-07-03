package respond

import (
	"io"
	"net/http"
)

func File(w http.ResponseWriter, code int, name string, contentType string, content io.Reader) {
	w.Header().Set("Content-Disposition", "attachment; filename=\""+name+"\"")
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(code)
	io.Copy(w, content)
}
