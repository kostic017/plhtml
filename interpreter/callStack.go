package interpreter

type callStack struct {
	top    *node
	length int
}

type node struct {
	value *actRecord
	prev  *node
}

func NewStack() *callStack {
	return &callStack{nil, 0}
}

func (callStack *callStack) len() int {
	return callStack.length
}

func (callStack *callStack) peek() *actRecord {
	if callStack.length == 0 {
		return nil
	}
	return callStack.top.value
}

func (callStack *callStack) pop() *actRecord {
	if callStack.length == 0 {
		return nil
	}
	n := callStack.top
	callStack.top = n.prev
	callStack.length--
	return n.value
}

func (callStack *callStack) push(value *actRecord) {
	n := &node{value, callStack.top}
	callStack.top = n
	callStack.length++
}
