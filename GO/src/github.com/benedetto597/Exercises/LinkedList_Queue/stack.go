package stack

import (
	"time"
)

type Node struct {
	name_model  string
	launch_date time.Time
	price       float64
	next        *Node
}

type Stack struct {
	last *Node
}

func (s *Stack) add(name_model string, launch_date time.Time, price float64) {
	new_node := new(Node)
	new_node.name_model = name_model
	new_node.launch_date = launch_date
	new_node.price = price

	if s.last == nil {
		s.last = new_node
	} else {
		current := s.last
		for current.next != nil {
			current = current.next
		}
		current.next = new_node
	}
}
