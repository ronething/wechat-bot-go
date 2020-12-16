package trie

import (
	"testing"
)

func newTestRouter() *Router {
	r := NewRouter()
	r.AddRoute("根", "/", nil)
	r.AddRoute("1", "/hello/:name", nil)
	r.AddRoute("2", "/hello/b/c", nil)
	r.AddRoute("3", "/hi/:name", nil)
	r.AddRoute("4", "/assets/*filepath", nil)
	return r
}

func TestGetRoutes(t *testing.T) {
	r := newTestRouter()
	n, ps1 := r.getRoute("/hello/panda")
	t.Logf("n is %v\n", n)
	t.Logf("ps1 is %v\n", ps1)
}
