type stackEntry struct {
	val  int
	next *stackEntry
}

type Stack struct {
	head *stackEntry
}

func (this *Stack) init(val int) {
	this.head = &stackEntry{
		val:  val,
		next: this.head,
	}
}

func (this *Stack) Push(val int) {

	if this.head == nil {
		this.init(val)
		return
	}

	newEntry := &stackEntry{
		val:  val,
		next: this.head,
	}

	this.head = newEntry
}

func (this *Stack) Pop() *stackEntry {
	popped := this.head
	this.head = this.head.next
	return popped
}

func (this *Stack) Top() *stackEntry {

	return this.head
}

func NewStack() *Stack {
	return &Stack{}
}

type MinStack struct {
	backingStack *Stack
	minStack     *Stack
}

func Constructor() MinStack {
	return MinStack{
		backingStack: NewStack(),
		minStack:     NewStack(),
	}
}

func (this *MinStack) Push(val int) {
	this.backingStack.Push(val)

	min := this.minStack.Top()
	if min == nil || val <= min.val {
		this.minStack.Push(val)
	}
}

func (this *MinStack) Pop() {
	entry := this.backingStack.Pop()
	topMin := this.minStack.Top()

	if topMin != nil && entry.val == topMin.val {
		this.minStack.Pop()
	}
}

func (this *MinStack) Top() int {
	return this.backingStack.Top().val
}

func (this *MinStack) GetMin() int {
	// if this.head.minHead == nil {
	// 	return 0
	// }
	// return this.head.minHead.val

	return this.minStack.Top().val

}
