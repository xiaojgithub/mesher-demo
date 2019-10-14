package route

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"sync"
)

type Route struct {
	Method string
	Path   string
	Handle httprouter.Handle
}

func NewHandle(method, path string, handle httprouter.Handle) Route {
	return Route{
		Method: method,
		Path:   path,
		Handle: handle,
	}
}

func Get(path string, handle httprouter.Handle) Route {
	return NewHandle(http.MethodGet, path, handle)
}

func Post(path string, handle httprouter.Handle) Route {
	return NewHandle(http.MethodPost, path, handle)
}

func Put(path string, handle httprouter.Handle) Route {
	return NewHandle(http.MethodPut, path, handle)
}

func Patch(path string, handle httprouter.Handle) Route {
	return NewHandle(http.MethodPatch, path, handle)
}

func Delete(path string, handle httprouter.Handle) Route {
	return NewHandle(http.MethodDelete, path, handle)
}

func Head(path string, handle httprouter.Handle) Route {
	return NewHandle(http.MethodHead, path, handle)
}

func Options(path string, handle httprouter.Handle) Route {
	return NewHandle(http.MethodOptions, path, handle)
}

type routeGroup struct {
	rs         []Route
	name       string
	prefixPath string
	toPath     func(string) string
}

func NewRouteGroup(name string) *routeGroup {
	return &routeGroup{
		rs:   make([]Route, 0, 10),
		name: name,
		toPath: func(path string) string {
			return path
		},
	}
}

func (rg *routeGroup) PrefixPath(prefixPath string) *routeGroup {
	rg.prefixPath = prefixPath
	rg.toPath = func(path string) string {
		return fmt.Sprintf("%s%s", rg.prefixPath, path)
	}
	return rg
}

func (rg *routeGroup) Add(routes ...Route) *routeGroup {
	rg.rs = append(rg.rs, routes...)
	return rg
}

var (
	locker         sync.Mutex
	allRouteGroups = make(map[string]*routeGroup, 20)
	httpRouter     = httprouter.New()
)

//Routes
func Routes(rg *routeGroup) {
	if rg == nil || rg.name == "" {
		panic(fmt.Sprintf("register route is invalid, %#v", rg))
	}
	locker.Lock()
	defer locker.Unlock()
	_, exist := allRouteGroups[rg.name]
	if exist {
		log.Printf("%s route group more exist", rg.name)
		return
	}
	allRouteGroups[rg.name] = rg
}

//Handle 获取http.Handler，用于http listen and serve
func Handle() http.Handler {
	for name, rg := range allRouteGroups {
		for _, route := range rg.rs {
			path := rg.toPath(route.Path)
			log.Printf("register route %s, %s, %s", name, route.Method, path)
			httpRouter.Handle(route.Method, path, route.Handle)
		}
	}
	return httpRouter
}
