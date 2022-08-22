package middleware

import (
	"fmt"
	"github.com/gorilla/mux"
	uerror "github.com/lffwl/utility/error"
	"github.com/lffwl/utility/response"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestHttpResponse(t *testing.T) {

	go func() {
		router := mux.NewRouter()
		router.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
			response.Json(writer, uerror.HighErrorSuccessCode, "xxx", 12312321)
			return
		}).Methods(http.MethodGet)

		router.Use(HttpResponse)
		// 初始化
		srv := &http.Server{
			Handler: router,
			Addr:    ":9999",
		}

		log.Fatal(srv.ListenAndServe())
	}()

	res, err := http.Get("http://127.0.0.1:9999/test")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", body)

}
