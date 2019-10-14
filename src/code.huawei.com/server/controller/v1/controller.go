package v1

import (
	"code.huawei.com/server/service"
	"github.com/aseTo2016/go-pkg/pkg/route"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func init() {
	route.Routes(route.NewRouteGroup("go mesh api").PrefixPath("/demo").Add(
		route.Get("/hello", Hello),
		route.Get("/greeting", Greeting),
	))
}

func Hello(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	_, err := w.Write([]byte("Hello, go mesher demo"))
	if err != nil {
		log.Printf("write into failed. %s", err.Error())
	}
}

func Greeting(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	body, err := service.Greeting()
	status := http.StatusOK
	if err != nil {
		log.Printf("greeting catch a error, %s", err.Error())
		status = http.StatusInternalServerError
	}
	w.WriteHeader(status)
	_, err = w.Write(body)
	if err != nil {
		log.Printf("write response catch a error, %s", err.Error())
	}
}
