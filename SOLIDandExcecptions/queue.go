package main

import "fmt"

// Просто реализация очереди

type Element struct {
	cmd      ICommand
	retrying int
}

type Queue struct {
	queue []Element
}

func (q *Queue) Put(cmd ICommand, retrying int) error {
	q.queue = append(q.queue, Element{cmd: cmd, retrying: retrying})
	return nil
}

func (q *Queue) Take() (Element, error) {
	if len(q.queue) == 0 {
		return Element{}, fmt.Errorf("queue is empty")
	} else {
		frontElement := q.queue[0]
		q.queue = q.queue[1:]
		return frontElement, nil
	}
}
