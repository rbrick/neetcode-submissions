type node struct {
	prev, next *node
	key, value int
}

type LinkedList struct {
	head, tail *node
}

func (list *LinkedList) Insert(key, value int) *node {
	if list.head == nil {

		h := &node{
			prev: nil, next: nil,
			key: key, value: value,
		}

		list.head = h
		list.tail = h

		return h
	} else {
		// insert a new node
		n := &node{
			prev:  list.tail,
			key:   key,
			value: value,
		}

		list.tail.next = n
		list.tail = n

		return n
	}
}

func (l *LinkedList) Remove(n *node) {
	// orphan the node
	nprev, nnext := n.prev, n.next

	if nprev != nil {
		nprev.next = nnext
	}
	if nnext != nil {
		nnext.prev = nprev
	}

	if l.tail == n {
		l.tail = n.prev
	}

	if l.head == n {
		l.head = n.next
	}

	n.next = nil
	n.prev = nil

}

func (l *LinkedList) PushToFront(n *node) {

	if l.head == n {
		return
	}

	// push to front makes it the head (mru)

	// how do we do this?

	// it's creating a connection between n.prev <-> n.next
	// then doing the same with the head

	l.Remove(n) // remove node from list. now an orphan

	currHead := l.head
	currHead.prev = n
	n.next = currHead

	l.head = n
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

type LRUCache struct {
	cap int

	// head and tail are our linked list

	// head = MRU (most recently used)
	// tail = LRU (least recently used)

	// so when we Get(key) => push node to front (head)
	// Put(key) => either update the value OR insert key & push to front (head), evict if greater than capacity

	// all of these operations can be done in O(1) complexity

	linkedList *LinkedList

	backingMap map[int]*node
}

func Constructor(capacity int) LRUCache {

	// lets create an empty list
	return LRUCache{
		cap:        capacity,
		backingMap: make(map[int]*node),
		linkedList: NewLinkedList(),
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.backingMap[key]

	if !ok {
		return -1
	}

	this.linkedList.PushToFront(node)

	return node.value
}

func (this *LRUCache) Put(key int, value int) {

	if n, ok := this.backingMap[key]; ok {
		this.linkedList.PushToFront(n)
		n.value = value
		return
	}

	n := this.linkedList.Insert(key, value)

	this.backingMap[key] = n

	this.linkedList.PushToFront(n)

	if len(this.backingMap) > this.cap {
		t := this.linkedList.tail
		this.linkedList.Remove(t)
		delete(this.backingMap, t.key)
	}

}
