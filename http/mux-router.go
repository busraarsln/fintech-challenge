package http

import (
	"log"
	"net/http"

	"github.com/busraarsln/fintech-challenge/utils"
	"github.com/gorilla/mux"
)

var limiter = utils.NewIPRateLimiter(1, 10)

type muxRouter struct{}

var (
	muxDispatcher = mux.NewRouter().StrictSlash(true)
)

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}
func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}
func (*muxRouter) PUT(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("PUT")
}
func (*muxRouter) DELETE(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("DELETE")
}
func (*muxRouter) Handle(uri string, sh http.Handler) {
	muxDispatcher.Handle(uri, sh)
}
func (*muxRouter) SERVE(port string) {
	//fmt.Printf("Mux HTTP server running on port %v", port)
	//http.ListenAndServe(port, muxDispatcher)

	if err := http.ListenAndServe(port, limitMiddleware(muxDispatcher)); err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}
}

func limitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		limiter := limiter.GetLimiter(r.RemoteAddr)
		if !limiter.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
