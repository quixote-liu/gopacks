package handler

import (
	"moft/util"
	"net/http"
)

func ReceiveFile(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		util.ResponseJSONErr(w, http.StatusBadRequest, map[string]interface{}{
			"error":   err.Error(),
			"message": "parse request form failed",
		})
		return
	}

	// read file.
	 r.FormFile("file")
}
