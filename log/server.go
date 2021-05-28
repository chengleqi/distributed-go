package log

import (
	"io"
	stlog "log"
	"net/http"
	"os"
)

var log *stlog.Logger

type filelog string

func (fl filelog) Write(data []byte) (int, error) {
	file, err := os.OpenFile(string(fl), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return file.Write(data)
}

func Run(destination string) {
	log = stlog.New(filelog(destination), "go", stlog.LstdFlags)
}

func RegisterHandlers() {
	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			bytes, err := io.ReadAll(r.Body)
			if err != nil || len(bytes) == 0 {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			write(string(bytes))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	})
}

func write(message string) {
	log.Printf("%v\n", message)
}