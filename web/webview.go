package web

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

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
	f, err := os.OpenFile(webview_nodedefs, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		logrus.WithError(err).Error("nodedefs file open failed")
		resp.WriteHeader(500)
		return
	}

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
	bunch := []*LuantiMedia{}
	err := json.NewDecoder(req.Body).Decode(&bunch)
	if err != nil {
		logrus.WithError(err).Error("media Decode")
		resp.WriteHeader(500)
		return
	}

	for _, media := range bunch {
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
