
package main

import "fmt"

type Node struct { // can also be created as tree
	data  int
	left  *Node
	right *Node
}

func (n *Node) insert(k int) {
	if k < n.data {
		//Move to left
		if n.left == nil {
			n.left = &Node{data: k}
		} else {
			n.left.insert(k)
		}

	} else if k > n.data {
		//Move to right
		if n.right == nil {
			n.right = &Node{data: k}
		} else {
			n.right.insert(k)
		}

	}
	//fmt.Println(n)
}

func (n *Node) search(k int) bool {
	if n == nil {
		return false
	}
	if k < n.data {
		//Move to left & search
		return n.left.search(k)

	} else if k > n.data {
		//Move to right
		return n.right.search(k)
	}
	return true

}
func (n *Node) iterative(k int) *Node {
	for n != nil { // its a while loop
		if n.data == k {
			return n
		} else if k < n.data {
			n = n.left
		} else {
			n = n.right
		}
	}
	//fmt.Println(k, "is not found")
	return nil
}

func (n *Node) Preorder(r *Node) {
	if r != nil {
		fmt.Printf("%d %s", r.data, "-->")
		n.Preorder(r.left)
		n.Preorder(r.right)
	}

}
func (n *Node) Postorder(r *Node) {
	if r != nil {
		n.Postorder(r.left)
		n.Postorder(r.right)
		fmt.Printf("%d %s", r.data, "-->")
	}
}
func (n *Node) Inorder(r *Node) {
	if r != nil {
		n.Inorder(r.left)
		fmt.Printf("%d %s", r.data, "-->")
		n.Inorder(r.right)
	}
}

func (n *Node) inorderpredessor(r *Node) int {
	r = r.left
	for r.right != nil {
		r = r.right
	}
	return r.data

}

func (n *Node) delete(k int) *Node {
	if n == nil {
		return nil
	}
	if k < n.data {
		//moving to left & performing deletion in left tree
		if n.left != nil {
			n.left = n.left.delete(k)
		} else {
			fmt.Println("The Given Node is not present")
		}
	} else if k > n.data {
		//moving to right & performing deletion right tree
		if n.right != nil {
			n.right = n.right.delete(k)
		} else {
			fmt.Println("The Given Node is not present")
		}
	} else {
		// we found the node to be deleted
		if n.right == nil { // Node to be deleted has no right child
			temp := n.left
			n = nil     // deleting the node
			return temp // since deleted node has no right child, it can be replaced with its only child(left)
		} else if n.left == nil { // Node to be deleted has no left child
			temp := n.right
			n = nil     // deleting the node
			return temp // since deleted node has no right child, it can be replaced with its only child(right)
		} else {
			ipre := n.inorderpredessor(n)
			n.data = ipre
			n.left = n.left.delete(ipre)
		}
	}
	return n
}

func main() {
	s := &Node{data: 100}
	fmt.Println("\n\n Inserting..........................................................")
	s.insert(50)
	s.insert(600)
	s.insert(300)
	fmt.Println(s)
	fmt.Println(s.right)

	fmt.Println("\n\n Traversing..........................................................")
	fmt.Printf("%s \t", "PREORDER")
	s.Preorder(s)
	fmt.Println()
	fmt.Printf("%s \t", "POSTORDER")
	s.Postorder(s)
	fmt.Println()
	fmt.Printf("%s \t", "INORDER")
	s.Inorder(s)

	fmt.Println("\n\n Searching..........................................................")
	fmt.Println("For 300-->", s.search(300), "And  For 60 -->", s.search(60))
	fmt.Println("For 60 -->", s.iterative(60), "And For 300-->", s.iterative(100))

	fmt.Println("\n\n deleting..........................................................")
	s.Inorder(s)
	fmt.Println()
	s.delete(50)
	s.delete(600)
	s.Inorder(s)
	fmt.Println()
	s.delete(100)
	s.Inorder(s)

}
