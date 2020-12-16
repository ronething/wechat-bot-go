package trie

import (
	"bytes"
	"fmt"
	"strings"
)

type Router struct {
	root     *node // 根节点
	handlers map[string]HandlerFunc
}

//NewRouter
func NewRouter() *Router {
	return &Router{
		root:     &node{},
		handlers: make(map[string]HandlerFunc),
	}
}

func (r *Router) GetHandlerFunc(key string) HandlerFunc {
	return r.handlers[key]
}

// Only one * is allowed
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *Router) AddRoute(comment, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)

	key := pattern
	r.root.insert(comment, pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *Router) getRoute(path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)

	n := r.root.search(searchParts, 0)

	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}

	return nil, nil
}

func (r *Router) GetRouteByPath(path string) (*node, map[string]string) {
	return r.getRoute(path)
}

func (r *Router) getRoutes() []*node {
	nodes := make([]*node, 0)
	r.root.travel(&nodes)
	return nodes
}

func (r *Router) PrintRoutes() string {
	nodes := r.getRoutes()
	var text bytes.Buffer
	for _, node := range nodes {
		text.WriteString(fmt.Sprintf("- %s %s\n", node.pattern, node.comment))
	}
	return text.String()
}
