package route

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"testing"
	"time"
)

func TestHandle(t *testing.T) {
	hello := func(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
		w.WriteHeader(http.StatusAccepted)
	}
	Routes(
		NewRouteGroup("test api").PrefixPath("/test").Add(
			Get("/hello", hello),
			Post("/hello", hello),
		),
	)
	server := &http.Server{Addr: "127.0.0.1:8080", Handler: Handle()}
	go func() {
		time.Sleep(2 * time.Second)
		defer func() {
			err := server.Close()
			if err != nil {
				t.Logf("close server catch a err:%s", err.Error())
			}
		}()
		resp, err := http.DefaultClient.Get("http://127.0.0.1:8080/test/hello")
		if err == nil && resp != nil && resp.StatusCode == http.StatusAccepted {
			t.Log("get 127.0.0.1:8080/test/hello success")
			resp, err = http.DefaultClient.Post("http://127.0.0.1:8080/test/hello", "application/json", nil)
			if err == nil && resp != nil && resp.StatusCode == http.StatusAccepted {
				t.Log("post 127.0.0.1:8080/test/hello success")
				return
			}
			t.Logf("resp %#v, err %#v", resp, err)
			return
		}
		t.Logf("resp %#v, err %#v", resp, err)
		t.Failed()
	}()

	err := server.ListenAndServe()
	if err != nil {
		t.Logf("listen and serve catch a err:%s", err.Error())
	}
}
