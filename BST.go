package main

//Imports
import(
	"fmt"
	"strconv"
	"math/rand"
	"time"
)

//Define node type
type node struct{
	parent *node;
	leaves[2] *node;
	key int;
	val string;
}

//Define tree type
type tree struct{
	head *node;
}

//Recursive add function
func (n *node) add(key int, val string, pN *node) *node{
	if n == nil {
		//Base case, new node
		return &node{key: key, val: val, parent: pN};
	} else if key < n.key {
		//Key is < current node key, call leaves[0]
		new := n.leaves[0].add(key, val, n);
		if new != nil { n.leaves[0] = new }
		return nil;
	} else if key > n.key {
		//Key is > current node key, call leaves[1]
		new := n.leaves[1].add(key, val, n);
		if new != nil { n.leaves[1] = new }
		return nil;
	} else {
		//Base case, update key
		n.val = val;
		return nil;
	}
}

//Add key and val to tree, handling a new tree
func (t *tree) add(key int, val string){
	if t.head == nil {
		t.head = &node{key: key, val: val};
	} else {
		t.head.add(key, val, nil);
	}
}

//Recursively search the tree for a val given a key
func (n *node) find(key int) string {
	if n == nil { return "DNE"; }
	if n.key < key { return n.leaves[1].find(key); }
	if n.key > key { return n.leaves[0].find(key); }
	return n.val;
}

//Find entry for node is the tree
func (t *tree) find(key int) string {
	if t.head != nil {
		return t.head.find(key);
	}
	return "DNE";
}

//Find the lowest value node below the current node
func (n *node) findMin() *node {
	curr := n;
	for curr.leaves[0] != nil {
		curr = curr.leaves[0];
	}
	return curr;
}

//Replace a given node with another one
func (n *node) replace(new *node){
	pN := n.parent;
	if pN != nil {
		if pN.leaves[0] == n {
			pN.leaves[0] = new;
		} else {
			pN.leaves[1] = new;
		}
	}
	if new != nil {
		new.parent = pN;
	}
}

//Recursively remove key value pair from the tree
func (n *node) remove(key int){
	if n == nil {
		//If n is nil, cancel
		return;
	} else if n.key < key {
		//Recursive call down the tree
		n.leaves[1].remove(key);
		return;
	} else if n.key > key {
		//Recursive call down the tree
		n.leaves[0].remove(key);
		return;
	}
	if n.leaves[0] != nil && n.leaves[1] != nil {
		//Replace using the next appropriate element
		new := n.leaves[1].findMin();
		n.key = new.key;
		n.val = new.val;
		new.remove(new.key);
	} else if n.leaves[1] == nil {
		//Replace using the leftmost leaf
		n.replace(n.leaves[0]);
	} else if n.leaves[0] == nil {
		//Replace using the rightmost leaf
		n.replace(n.leaves[1]);
	} else {
		//Replace with nil
		n.replace(nil);
	}
}

//Remove key from tree
func (t *tree) remove(key int){
	if t.head != nil {
		t.head.remove(key);
	}
}

//Print the tree starting at the head
func (t *tree) print(){
	t.head.print();
}

//Recursively print all node key value pairs
func (n *node) print(){
	if n != nil {
		n.leaves[0].print();
		fmt.Println(n.key, ": ", n.val);
		n.leaves[1].print();
	}
}

func main() {
	var t tree;
	var strings[10000] string;

	for i := 0; i < 10000; i++ {
		strings[i] = strconv.Itoa(i);
	}

	var avg int64 = 0;
	var worst int64 = 0;

	r := rand.New(rand.NewSource(time.Now().Unix()))
  	for s, i := range r.Perm(len(strings)) {
		start := time.Now().UnixNano();
		t.add(i, strings[i]);
		now := time.Now().UnixNano() - start;
		if s == 0 { fmt.Println("First insertion: ", now, "ns"); }
		if now > worst { worst = now; }
		avg += now;
	}

	fmt.Println("Worst time: ", worst, "ns");
	fmt.Println("Average time: ", (avg/10000), "ns\n\n");

	found := true;
	avg = 0;
	worst = 0;

	for i := 0; i < 10000; i++ {
		start := time.Now().UnixNano();
		found = found && !(t.find(i) == "DNE");
		now := time.Now().UnixNano() - start;
		if now > worst { worst = now; }
		avg += now;
	}

	fmt.Println("All entries found: ", found);
	fmt.Println("Worst time: ", worst, "ns");
	fmt.Println("Average time: ", (avg/10000), "ns");

}