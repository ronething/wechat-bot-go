package trie

import (
	"fmt"
	"strings"
)

type node struct {
	pattern  string //有效路径，类似于单词匹配的 isWord
	part     string
	children []*node
	comment  string // 备注说明
	isWild   bool   // part 前缀有 : 或者 * 的时候为 true
}

func (n *node) Pattern() string {
	return n.pattern
}

func (n *node) String() string {
	return fmt.Sprintf("node{comment=%s,pattern=%s, part=%s, isWild=%t}", n.comment, n.pattern, n.part, n.isWild)
}

func (n *node) insert(comment, pattern string, parts []string, height int) {
	if len(parts) == height { //到顶了
		n.pattern = pattern
		n.comment = comment
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(comment, pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" { //证明不是一个有效的路径
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}

func (n *node) travel(list *([]*node)) {
	if n.pattern != "" {
		*list = append(*list, n)
	}
	for _, child := range n.children {
		child.travel(list)
	}
}

func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}
