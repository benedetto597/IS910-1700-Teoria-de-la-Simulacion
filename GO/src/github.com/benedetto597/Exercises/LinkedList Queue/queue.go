/*
	@author benedetto597
	@date 2020-04-21
	@version 1.0.0

	=== DESCRIPTION ===


*/

package linkedlistqueue

import (
	"fmt"
	"time"
)

type Node struct {
	name_model  string
	launch_date time.Time
	price       float64
	next        *Node
}

type Queue struct {
	first *Node
}

func (n *Node) print_node() {
	fmt.Println("\n============= Node =============")
	if n != nil {
		fmt.Println(n.name_model, n.launch_date, n.price)
	} else {
		fmt.Println("Empty Node")
	}
}

func (q *Queue) add(name_model string, launch_date time.Time, price float64) {
	new_node := new(Node)
	new_node.name_model = name_model
	new_node.launch_date = launch_date
	new_node.price = price

	if q.first == nil {
		q.first = new_node
	} else {
		current := q.first
		for current.next != nil {
			current = current.next
		}
		current.next = new_node
	}
}

func (q *Queue) del_last() *Node {
	if q.first == nil {
		return nil
	} else {
		current := q.first
		for current.next.next != nil {
			current = current.next
		}
		current.next = nil
		return current.next
	}
}

func (q *Queue) del_position(pos int) *Node {
	if q.first == nil {
		return nil
	} else if pos > q.get_size() {
		fmt.Printf("\nCan't Delete - Position %d is out of range\n", pos)
		return nil
	} else {
		index := 1
		current := q.first
		if pos == 1 {
			q.first = current.next
		}
		for index != pos-1 && current.next != nil {
			current = current.next
			index++
		}
		current.next = nil
		return current.next
	}
}

func (q *Queue) print() {
	if q.first == nil {
		fmt.Println("Empty Queue")
	} else {
		current := q.first
		fmt.Println("\n============= Queue =============")
		for current != nil {
			fmt.Println(current.name_model, current.launch_date, current.price)
			current = current.next
		}
	}
}

func (q *Queue) get_first() *Node {
	return q.first
}

func (q *Queue) get_last() *Node {
	if q.first == nil {
		return nil
	} else {
		current := q.first
		for current.next != nil {
			current = current.next
		}
		return current
	}
}

func (q *Queue) get_position(pos int) *Node {

	if q.first == nil {
		fmt.Println("\nEmpty Queue")
		return nil
	} else if pos > q.get_size() {
		fmt.Printf("\nCan't Get - Position %d is out of range\n", pos)
		return nil
	} else {
		index := 1
		current := q.first
		for index != pos {
			current = current.next
			index++
		}
		return current
	}
}

func (q *Queue) get_size() int {
	if q.first == nil {
		return 0
	} else {
		index := 1
		current := q.first
		for current.next != nil {
			current = current.next
			index++
		}
		return index
	}
}

func (q *Queue) add_position(name_model string, launch_date time.Time, price float64, pos int) {
	new_node := new(Node)
	new_node.name_model = name_model
	new_node.launch_date = launch_date
	new_node.price = price

	if q.first == nil {
		q.first = new_node
	} else if pos > q.get_size() {
		fmt.Printf("\nCan't Add - Position %d is out of range\n", pos)
	} else {
		index := 2
		current := q.first
		for index != pos {
			current = current.next
			index++
		}
		new_node.next = current.next
		current.next = new_node
	}
}

func main() {
	// Placas de computadora --> Nombre, Fecha de lanzamiento, Precio
	ll_queue := new(Queue)
	timeString := "2021-08-15 02:30:45"
	theTime, err := time.Parse("2006-01-02 03:04:05", timeString)
	if err != nil {
		fmt.Println("Could not parse time:", err)
	}
	ll_queue.add("Asus Saphire", theTime, 100.99)
	ll_queue.add("AS Rock", theTime, 566.98)
	ll_queue.add("Asus Metal", theTime, 774.65)
	ll_queue.add("MSI Phantom", theTime, 7443)
	ll_queue.add("MSI Carbon", theTime, 355)
	ll_queue.add("TTO", theTime, 788)

	ll_queue.print()

	ll_queue.del_last()
	ll_queue.del_last()
	ll_queue.del_last()

	ll_queue.print()

	ll_queue.get_first().print_node()
	ll_queue.get_last().print_node()
	ll_queue.get_position(3).print_node()
	ll_queue.get_position(5).print_node()

	fmt.Println("\nQueue Size:", ll_queue.get_size())

	ll_queue.add_position("Asus Google", theTime, 988, 2)
	ll_queue.add_position("MSI Catracha", theTime, 318.12, 3)

	ll_queue.print()

	ll_queue.del_position(5)
	ll_queue.del_position(1)

	ll_queue.print()
}
