# go-pkg
some common pkg   
1. route:base on httproute   
demo:
```go
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
```
