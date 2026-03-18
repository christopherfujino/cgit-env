package notes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"

	"christopherfujino.com/distributed-compute-monorepo/service"
)

type GetNotesResponse struct {
	Files []string `json:"files"`
}

type UpdateNotesRequest struct {
	Path     string `json:"path"`
	Contents string `json:"contents"`
}

type tee struct {
	dir          string
	endpointRoot string
}

func (t tee) Register() {
	var fileServer = http.StripPrefix(t.endpointRoot, http.FileServer(http.Dir(filepath.Join(t.dir, "assets"))))
	http.Handle(fmt.Sprintf("GET %s/", t.endpointRoot), fileServer)

	http.HandleFunc(fmt.Sprintf("GET /api%s", t.endpointRoot), func(w http.ResponseWriter, r *http.Request) {
		var db = Open()
		log.Printf("GET %v\n", r.URL.Path)

		var res = GetNotesResponse{
			Files: db.GetAllPaths(),
		}
		resBytes, err := json.Marshal(res)
		if err != nil {
			w.WriteHeader(500)
			// TODO sanitize this?
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		_, err = w.Write(resBytes)
		if err != nil {
			log.Printf("Error failed to write response: %s\n", err.Error())
		}
	})

	http.HandleFunc("UPDATE /api/notes/update", func(w http.ResponseWriter, r *http.Request) {
		var db = Open()
		log.Printf("UPDATE %v\n", r.URL.Path)
		var buffer = bytes.Buffer{}
		io.Copy(&buffer, r.Body)

		var reqData UpdateNotesRequest
		err := json.Unmarshal(buffer.Bytes(), &reqData)
		if err != nil {
			w.WriteHeader(500)
			// TODO sanitize this?
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		db.Write(reqData.Path, reqData.Contents)
		fmt.Printf("%s: %s\n", reqData.Path, reqData.Contents)
		w.WriteHeader(200)
	})
}

func Create(dirPath, endpointRoot string) service.T {
	return tee{
		dir: dirPath,
		endpointRoot: endpointRoot,
	}
}
