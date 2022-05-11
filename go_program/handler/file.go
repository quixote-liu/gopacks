package handler

import (
	"bufio"
	"io/fs"
	"log"
	"moft/model"
	"moft/util"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func ReceiveFile(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		util.ResponseJSONErr(w, http.StatusBadRequest, model.H{
			"error":   err.Error(),
			"message": "parse request form failed",
		})
		return
	}

	// open file path.
	dirPath := model.CONF.GetString("file", "path")
	err := os.Mkdir(dirPath, fs.ModeDir)
	if err != nil && !os.IsExist(err) {
		log.Printf("create directory failed: %v", err)
		return
	}

	// read file.
	f, fh, err := r.FormFile("file")
	if err != nil {
		util.ResponseJSONErr(w, http.StatusBadRequest, model.H{
			"error":   err.Error(),
			"message": "get form file failed",
		})
		return
	}
	defer f.Close()

	// read file content to scan.
	scan := bufio.NewScanner(f)

	fileName := fh.Filename + "_" + time.Now().String()
	fileName = joinDirAndFileName(dirPath, fileName)

	// create file.
	file, err := os.Create(fileName)
	if err != nil {
		log.Printf("create file(%s) failed: %v", fileName, err)
		util.ResponseJSONErr(w, http.StatusInternalServerError, nil)
		return
	}

	// wirte data.
	_, err = file.Write(scan.Bytes())
	if err != nil {
		log.Printf("write data to file(%s) failed: %v", fileName, err)
		util.ResponseJSONErr(w, http.StatusInternalServerError, nil)
		return
	}
}

func joinDirAndFileName(dir, fileName string) string {
	dir = filepath.Clean(dir)
	if !strings.HasPrefix(dir, "/") {
		dir += "/"
	}
	return dir + fileName
}
