package web

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/sirupsen/logrus"
)

var webview_directory, webview_nodedefs string

func init() {
	wd, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("getcwd error: %v", err))
	}

	webview_directory = path.Join(wd, "mapserver_webview")
	err = os.MkdirAll(webview_directory, 0777)
	if err != nil {
		panic(fmt.Errorf("mkdirall '%s' error: %v", webview_directory, err))
	}

	webview_nodedefs = path.Join(webview_directory, "nodedefs.json")
}

func (api *Api) PostLuantiNodedefs(resp http.ResponseWriter, req *http.Request) {
	if !api.check_auth(resp, req) {
		return
	}

	f, err := os.OpenFile(webview_nodedefs, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		logrus.WithError(err).Error("nodedefs file open failed")
		resp.WriteHeader(500)
		return
	}
	defer f.Close()

	n, err := io.Copy(f, req.Body)
	if err != nil {
		logrus.WithError(err).Error("nodedefs file write failed")
		resp.WriteHeader(500)
		return
	}

	logrus.Infof("Processed %d bytes of nodedefinitions", n)
	resp.WriteHeader(200)
}

type LuantiMedia struct {
	Filename string `json:"filename"`
	Data     string `json:"data"`
	Type     string `json:"type"`
}

func (api *Api) PostLuantiMedia(resp http.ResponseWriter, req *http.Request) {
	if !api.check_auth(resp, req) {
		return
	}

	bunch := []*LuantiMedia{}
	err := json.NewDecoder(req.Body).Decode(&bunch)
	if err != nil {
		logrus.WithError(err).Error("media Decode")
		resp.WriteHeader(500)
		return
	}

	for _, media := range bunch {
		if !isValidMediaFilename(media.Filename) {
			// not a plain filename, would be written outside of the media directory
			logrus.Warnf("skipping media with invalid filename: '%s'", media.Filename)
			continue
		}

		data, err := base64.RawStdEncoding.DecodeString(media.Data)
		if err != nil {
			logrus.WithError(err).Errorf("base64 Decode: %s", media.Filename)
			resp.WriteHeader(500)
			return
		}

		subpath := "textures"
		if media.Type == "model" {
			subpath = "models"
		}

		os.MkdirAll(path.Join(webview_directory, subpath), 0777)
		filepath := path.Join(webview_directory, subpath, media.Filename)

		err = os.WriteFile(filepath, data, 0666)
		if err != nil {
			logrus.WithError(err).Errorf("file write: '%s'", filepath)
			resp.WriteHeader(500)
			return
		}
	}

	logrus.Infof("Processed %d media entrie(s)", len(bunch))
	resp.WriteHeader(200)

}

// serves the node definitions posted by the mapserver mod
func (api *Api) GetViewNodedefs(resp http.ResponseWriter, req *http.Request) {
	http.ServeFile(resp, req, webview_nodedefs)
}

// max number of files in a media batch request, see public/src/util/ViewMediaLoader.js
const maxViewMediaBatch = 128

// returns true if the name is a plain filename without any directory parts
func isValidMediaFilename(filename string) bool {
	return filename != "" && filename == path.Base(filename) && filename != ".." && filename != "." &&
		!strings.ContainsAny(filename, "/\\")
}

// returns the path and type of a texture or model posted by the mapserver mod, empty if it does not exist
func findViewMedia(filename string) (string, string) {
	if !isValidMediaFilename(filename) {
		return "", ""
	}

	for _, mediatype := range []string{"texture", "model"} {
		filepath := path.Join(webview_directory, mediatype+"s", filename)
		fi, err := os.Stat(filepath)
		if err == nil && !fi.IsDir() {
			return filepath, mediatype
		}
	}

	return "", ""
}

// serves a texture or model posted by the mapserver mod
func (api *Api) GetViewMedia(resp http.ResponseWriter, req *http.Request) {
	filepath, _ := findViewMedia(req.PathValue("filename"))
	if filepath == "" {
		resp.WriteHeader(http.StatusNotFound)
		return
	}

	http.ServeFile(resp, req, filepath)
}

// serves a batch of textures and models, expects a list of filenames
// files that do not exist are omitted from the result
func (api *Api) PostViewMedia(resp http.ResponseWriter, req *http.Request) {
	filenames := []string{}
	err := json.NewDecoder(http.MaxBytesReader(resp, req.Body, 64*1024)).Decode(&filenames)
	if err != nil || len(filenames) > maxViewMediaBatch {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	bunch := []*LuantiMedia{}
	for _, filename := range filenames {
		filepath, mediatype := findViewMedia(filename)
		if filepath == "" {
			continue
		}

		data, err := os.ReadFile(filepath)
		if err != nil {
			logrus.WithError(err).Errorf("file read: '%s'", filepath)
			resp.WriteHeader(http.StatusInternalServerError)
			return
		}

		bunch = append(bunch, &LuantiMedia{
			Filename: filename,
			Data:     base64.RawStdEncoding.EncodeToString(data),
			Type:     mediatype,
		})
	}

	resp.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(resp).Encode(bunch)
	if err != nil {
		logrus.WithError(err).Error("media batch encode")
	}
}
