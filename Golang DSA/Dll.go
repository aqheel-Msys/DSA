package main

import "fmt"

type node struct {
	data int
	next *node
}

type link struct {
	head *node
	tail *node
	size int
}

func (l *link) len() int {
	return l.size
}

func (l *link) isempty() bool {
	return l.size == 0

}

func (l *link) adf(n int) {
	new := &node{n, nil}

	if l.isempty() {
		l.head = new
		l.tail = new
	} else {
		new.next = l.head
	}
	l.head = new
	l.size += 1

}

func (l *link) adl(n int) {
	new := &node{n, nil}
	if l.isempty() {
		l.head = new
	} else {
		l.tail.next = new
	}
	l.tail = new
	l.size += 1

}

func (l *link) ada(n, i int) {
	new := &node{n, nil}
	if i == 0 {
		l.adf(n)
	} else if i == l.len()-1 {
		l.adl(n)

	} else {
		p := l.head
		a := 0
		for a < i-1 {
			p = p.next
			a += 1
		}
		new.next = p.next
		p.next = new
		l.size += 1

	}

}

func (l *link) rmf() int {
	var a int
	if l.isempty() {
		fmt.Println("linkedlist is empty")
	} else {
		a = l.head.data
		l.head = l.head.next
	}
	l.size -= 1
	return a

}

func (l *link) rml() int {
	var a int
	if l.isempty() {
		fmt.Println("linkedlist is empty")
	} else {
		p := l.head
		i := 1
		for i < l.len()-1 {
			p = p.next
			i += 1
		}

		a = p.next.data
		l.tail = p
		p.next = nil
		l.size -= 1

	}
	return a
}
func (l *link) rma(i int) int {
	var a int
	if i == 0 {
		l.rmf()
	} else if i == l.len()-1 {
		l.rml()
	} else {
		p := l.head
		e := 0
		for e < i-1 {
			p = p.next
			e += 1
		}
		a = p.next.data
		p.next = p.next.next
		l.size -= 1

	}
	return a

}
func (l *link) display() {
	a := l.head
	for a != nil {
		fmt.Print(a.data, "-->")
		a = a.next
	}
	fmt.Println()
}

func main() {
	l := link{}
	l.adf(10)
	l.adf(20)
	l.adf(30)
	l.display()
	l.adl(20)
	l.adl(10)
	l.display()
	l.ada(50, l.len()-1)
	l.display()
	fmt.Println(l.rmf())
	l.display()
	fmt.Println(l.rml())
	l.display()
	fmt.Println(l.rma(l.len() - 1))
	l.display()

}
