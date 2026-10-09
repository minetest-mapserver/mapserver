package web

import (
	"crypto/subtle"
	"net/http"
)

func (api *Api) check_auth(resp http.ResponseWriter, req *http.Request) bool {
	key := req.Header.Get("Authorization")
	if subtle.ConstantTimeCompare([]byte(key), []byte(api.Context.Config.WebApi.SecretKey)) != 1 {
		resp.WriteHeader(403)
		resp.Write([]byte("invalid key!"))
		return false
	}

	return true
}
