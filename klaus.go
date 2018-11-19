package klaus

import "net/http"

// Engine ...
type Engine struct {
}

// New ...
func New() *Engine {
	engine := &Engine{}

	return engine
}

// ServeHTTP conforms to the http.Handler interface.
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//c := engine.pool.Get().(*Context)
	//c.writermem.reset(w)
	//c.Request = req
	//c.reset()

	//engine.handleHTTPRequest(c)

	//engine.pool.Put(c)
}

// Run ...
func Run(engine *Engine) (err error) {
	err = http.ListenAndServe(":8080", engine)
	return
}
