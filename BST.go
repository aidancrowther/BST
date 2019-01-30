package main

import(
	"fmt"
)

type node struct{
	parent *node;
	leaves[2] *node;
	data int;
}

type tree struct{
	head *node;
}

func (t *tree) add(i int){
	if t.head == nil {
		t.head = &node{data: i};
	} else {
		t.head.add(i, nil);
	}
}

func (n *node) add(i int, pN *node) *node{
	if n == nil {
		return &node{data: i, parent: pN};
	} else if i < n.data {
		new := n.leaves[0].add(i, n);
		if new != nil { n.leaves[0] = new }
		return nil;
	} else if i > n.data {
		new := n.leaves[1].add(i, n);
		if new != nil { n.leaves[1] = new }
		return nil;
	}
	return nil;
}

func (t *tree) print(){
	if t.head != nil {
		t.head.print();
	}
}

func (n *node) print(){
	if n != nil {
		n.leaves[0].print();
		fmt.Println(n.data);
		n.leaves[1].print();
	}
}

func main() {
	var t tree;
	t.add(0);
	t.print();
}