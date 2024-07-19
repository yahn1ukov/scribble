package respond

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func File(w http.ResponseWriter, name string, contentType string, content []byte) {
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", name))
	w.Header().Set("Content-Type", contentType)
	io.Copy(w, bytes.NewReader(content))
}
