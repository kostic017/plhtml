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

func (this *callStack) len() int {
    return this.length
}

func (this *callStack) peek() *actRecord {
    if this.length == 0 {
        return nil
    }
    return this.top.value
}

func (this *callStack) pop() *actRecord {
    if this.length == 0 {
        return nil
    }
    n := this.top
    this.top = n.prev
    this.length--
    return n.value
}

func (this *callStack) push(value *actRecord) {
    n := &node{value, this.top}
    this.top = n
    this.length++
}
