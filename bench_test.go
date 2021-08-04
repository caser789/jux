package mux

import (
	"net/http"
	"testing"
)

func BenchmarkMux(b *testing.B) {
	router := new(Router)
	handler := func(w http.ResponseWriter, r *http.Request) {}
	router.HandleFunc("/v1/{v1}", handler)

	request, _ := http.NewRequest("GET", "/v1/anything", nil)
	for i := 0; i < b.N; i++ {
		router.ServeHTTP(nil, request)
	}
}
