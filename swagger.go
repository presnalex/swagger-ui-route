package swagger

import (
	"net/http"
	"os"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/gorilla/mux"
)

const (
	path_slash   = "/swagger-ui/"
	path_noslash = "/swagger-ui"
)

type asset func(string) ([]byte, error)
type assetdir func(string) ([]string, error)
type assetinfo func(string) (os.FileInfo, error)

func Register(r *mux.Router, appName string, a asset, d assetdir, i assetinfo) {
	basePath := "/" + appName
	r.PathPrefix(
		basePath + path_slash).Handler(
		http.StripPrefix(basePath+path_slash,
			http.FileServer(
				&assetfs.AssetFS{
					Asset:     a,
					AssetDir:  d,
					AssetInfo: i,
					Prefix:    "/",
				},
			),
		),
	)
	r.HandleFunc(basePath+path_noslash, func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, basePath+path_slash, http.StatusMovedPermanently)
	})
}
