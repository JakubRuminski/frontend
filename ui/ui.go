package ui

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"strings"
)

// go:embed build
var staticFS embed.FS

// AddRoutes serves the static file system for the UI React App.
func AddRoutes( ) *staticFileSystem {
	
	embeddedBuildFolder := newStaticFileSystem()
	
	// TODO: replicate this without gin...
	// router.Use(static.Serve("/", embeddedBuildFolder))

	return embeddedBuildFolder
}


type staticFileSystem struct {
	http.FileSystem
}


func newStaticFileSystem() *staticFileSystem {
	sub, err := fs.Sub(staticFS, "build")

	if err != nil {
		panic(err)
	}

	return &staticFileSystem{ FileSystem: http.FS(sub) }
}


func (s *staticFileSystem) Exists(prefix string, path string) bool {
	buildpath := fmt.Sprintf("build%s", path)

	// support for folders
	if strings.HasSuffix(path, "/") {
		_, err := staticFS.ReadDir(strings.TrimSuffix(buildpath, "/"))
		return err == nil
	}

	// support for files
	f, err := staticFS.Open(buildpath)
	if f != nil {
		_ = f.Close()
	}
	return err == nil
}