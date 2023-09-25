package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	listenAddress string
	//storage
}

func (s *ApiServer) Run() {
	r := mux.NewRouter()

	log.Println("Server running on port:", s.listenAddress)
	http.ListenAndServe(s.listenAddress, r)
}

func NewApiServer(listenAddr string) *ApiServer {
	return &ApiServer{
		listenAddress: listenAddr,
	}
}

// add a file or folder with registration handlers (handlereg, handleregauth, handlelogin, handleloginauth)
