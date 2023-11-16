package jee

// handleFunc(w http.ResponseWriter r *http.Request)
type HandleFunc func(c *Context)
