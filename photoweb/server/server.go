package server

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

const fileBasePath = "/data/photo/"

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
		file, e := os.Create(fileBasePath + fileName)
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
	http.ServeFile(w, r, fileBasePath+fileName)
}

//文件列表
func FileList(w http.ResponseWriter, r *http.Request) {
	filearr, err := ioutil.ReadDir(fileBasePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var listHtml string
	for _, info := range filearr {
		listHtml += "<li><a href=\"/view?id=" + info.Name() + "\">"+info.Name()+"</a></li>"
	}
	io.WriteString(w, "<div><ol>"+listHtml+"</ol></div>")
}
