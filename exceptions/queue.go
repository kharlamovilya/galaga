package exceptions

import "fmt"

// Просто реализация очереди

type Element struct {
	cmd      ICommand
	retrying int
}

func NewElement(cmd ICommand, retrying int) *Element {
	return &Element{cmd: cmd, retrying: retrying}
}

func (e Element) Cmd() ICommand {
	return e.cmd
}

func (e Element) Retrying() int {
	return e.retrying
}

type Queue struct {
	queue []Element
}

func (q *Queue) GetQueue() []Element {
	return q.queue
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

func (q *Queue) Pop() error {
	if len(q.queue) == 0 {
		return fmt.Errorf("queue is empty, no element to pop")
	}
	q.queue = q.queue[0 : len(q.queue)-1]
	return nil
}
