package framework

import (
	"errors"
	"fmt"
	"strings"
)

type Tree struct {
	root *Node
}

func NewTree() *Tree {
	return &Tree{root: &Node{}}
}

type Node struct {
	isLast      bool         // 是否是最后一个节点
	controllers []Controller // 控制器
	segment     string       // 段字符串
	children    []*Node      // 子节点
}

// 是否是通配符
func isWildSegment(segment string) bool {
	return strings.HasPrefix(segment, ":")
}

// /user/:id
// /users/list
// /user/list
// 匹配node
func (n *Node) matchNode(router string) *Node {

	// 将路由切成2段
	segments := strings.SplitN(router, "/", 2)

	// 获取第一段
	segment := segments[0]

	// 如果不是通配符类型，就转换成大写
	if !isWildSegment(segment) {
		segment = strings.ToUpper(segment)
	}

	// 查找该段内容在下一级节点中的所有匹配的子节点
	childNodes := n.filterChildNodes(segment)
	if len(childNodes) <= 0 {
		return nil
	}

	// 如果段大小为1，说明是最后一次查找
	if len(segments) == 1 {
		// 如果有和段匹配的节点，且必须是末尾节点，才算匹配成功
		for _, node := range childNodes {
			if node.isLast {
				return node
			}
		}
	}

	for _, node := range childNodes {
		findNode := node.matchNode(segments[1])
		if findNode != nil {
			return findNode
		}
	}

	return nil
}

// 获取所有子节点
func (n *Node) filterChildNodes(segment string) []*Node {

	if isWildSegment(segment) {
		return n.children
	}

	var nodes []*Node
	for _, child := range n.children {
		if isWildSegment(child.segment) {
			nodes = append(nodes, child)
		} else if child.segment == segment {
			nodes = append(nodes, child)
		}
	}

	return nodes
}

// 新增路由
// /user/:id
// /users/list
// /user/list
func (t *Tree) AddRouter(router string, controllers ...Controller) error {
	n := t.root
	if n.matchNode(router) != nil {
		return errors.New(fmt.Sprintf("exist router: %v", router))
	}

	segments := strings.Split(router, "/")
	for i, segment := range segments {

		if !isWildSegment(segment) {
			segment = strings.ToUpper(segment)
		}

		// 是否是最后一段
		isLast := i == len(segments)-1

		// 从子节点中查找是否存在该段
		var segmentNode *Node
		for _, cnode := range n.children {
			if cnode.segment == segment {
				segmentNode = cnode
				break
			}
		}

		// 如果不存在，需要创建并加入到上级节点的子节点上
		if segmentNode == nil {
			segmentNode = &Node{segment: segment}
			if isLast {
				segmentNode.isLast = isLast
				segmentNode.controllers = controllers
			}
			n.children = append(n.children, segmentNode)
		}

		// 继续遍历
		n = segmentNode
	}

	return nil
}

// 查找controller
func (t *Tree) FindController(router string) []Controller {
	matchNode := t.root.matchNode(router)
	if matchNode == nil {
		return nil
	}
	return matchNode.controllers
}
