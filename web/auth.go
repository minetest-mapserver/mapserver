package web

import "net/http"

func (api *Api) check_auth(resp http.ResponseWriter, req *http.Request) bool {
	if req.Header.Get("Authorization") != api.Context.Config.WebApi.SecretKey {
		resp.WriteHeader(403)
		resp.Write([]byte("invalid key!"))
		return false
	}

	return true
}
