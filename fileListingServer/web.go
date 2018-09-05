package main

import (
	"net/http"
	_ "net/http/pprof"

	"log"
	"os"
	"selfLearning/fileListingServer/filelisting"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

// 函数式 编程
func errWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		//recover
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic : %v", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		err := handler(writer, request)

		if err != nil {
			//log.Warn("Error handling request : %s", err.Error())
			log.Printf("Error occured "+"handler request : %s", err.Error())
			// user Error
			if userErr, ok := err.(userErr); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}
			// system Error
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer,
				http.StatusText(code), code)
		}
	}
}

type userErr interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/",
		errWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		panic(err)
	}
}
