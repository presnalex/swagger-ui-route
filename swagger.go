package swagger

import (
	"io/fs"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	path_slash   = "/swagger-ui/"
	path_noslash = "/swagger-ui"
)

func Register(r *mux.Router, appName string, embededFiles fs.FS) {
	basePath := "/" + appName
	r.PathPrefix(
		basePath + path_slash).Handler(
		http.StripPrefix(basePath+path_slash,
			http.FileServer(getFileSystem(embededFiles)),
		),
	)
	r.HandleFunc(basePath+path_noslash, func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, basePath+path_slash, http.StatusMovedPermanently)
	})
}

func getFileSystem(embededFiles fs.FS) http.FileSystem {
	fsys, err := fs.Sub(embededFiles, "swagger-ui")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}
