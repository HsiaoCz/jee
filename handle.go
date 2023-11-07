package jee

import "net/http"

// handleFunc(w http.ResponseWriter r *http.Request)
type HandleFunc func(w http.ResponseWriter, r *http.Request)
