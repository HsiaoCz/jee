package jee

import "strings"

type node struct {
	// 每个节点上的数据
	part     string
	children map[string]*node
	// 每个节点上存储的视图函数
	handlefunc HandleFunc
	// 子节点
	paramChildren *node
}

func (n *node) addNode(part string) *node {
	// 动态路由
	if strings.HasPrefix(part, ":") && n.paramChildren == nil {
		n.paramChildren = &node{part: part}
		return n.paramChildren
	}
	// 判断当前节点有没有children属性
	if n.children == nil {
		n.children = make(map[string]*node)
	}
	child, ok := n.children[part]
	if !ok {
		child = &node{
			part: part,
		}
		n.children[part] = child
	}
	return child
}

func (n *node) getNode(part string) *node {
	// n的children属性都不存在
	if n.children == nil {
		return nil
	}
	// 正常思路：先到静态路由中找
	child, ok := n.children[part]
	if !ok {
		if n.paramChildren != nil {
			return n.paramChildren
		}
		return nil
	}
	return child
}
