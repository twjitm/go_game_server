package server

import (
	"io"
	"net/http"
	"os"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		_, _ = io.WriteString(w, "<div><form method=\"POST\" action=\"/upload\" "+
			" enctype=\"multipart/form-data\">"+
			"Choose an image to upload: <input name=\"image\" type=\"file\" />"+
			"<input type=\"submit\" value=\"Upload\" />"+
			"</form></div>")

		return
	}
	if r.Method == "POST" {
		f, h, e := r.FormFile("image")
		if e != nil {
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()
		fileName := h.Filename
		file, e := os.Create("/data/" + fileName)
		if e != nil {
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		if _, err := io.Copy(file, f); err != nil {
			http.Error(w, err.Error(),
				http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/view?id="+fileName, http.StatusFound)
	}
}

func View(w http.ResponseWriter, r *http.Request) {

	fileName := r.FormValue("id")
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, "/data/"+fileName)
}
