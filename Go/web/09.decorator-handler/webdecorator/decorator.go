package webdecorator

import "net/http"

type DecoratorFunc func(http.ResponseWriter, *http.Request, http.Handler)

type DecoratorHandler struct {
	decorator DecoratorFunc
	handler   http.Handler
}

func (self *DecoratorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	self.decorator(w, r, self.handler)
}

func NewHandler(h http.Handler, decorator DecoratorFunc) http.Handler {
	return &DecoratorHandler{
		decorator: decorator,
		handler:   h,
	}
}
