package serve

import (
	_ "code.huawei.com/server/controller/v1"
)

import (
	"github.com/aseTo2016/go-pkg/pkg/route"
	"log"
	"net/http"
	"os"
	"strings"
)

func Serve() {
	ip := os.Getenv("LISTEN_IP")
	if len(ip) == 0 {
		ip = "127.0.0.1"
	}
	port := os.Getenv("LISTEN_PORT")
	if len(port) == 0 {
		port = "8080"

	}
	address := strings.Join([]string{ip, port}, ":")
	log.Printf("start listen %s", address)
	err := http.ListenAndServe(address, route.Handle())
	if err != nil {
		log.Printf("listen and serve catch a error: %s", err.Error())
	}
}
