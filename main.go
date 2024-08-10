package golang_message_queue

import (
	"fmt"
)

type Item struct {
	value interface{} // hold any data type
	next  *Item
}
type Queue struct {
	front  *Item
	rear   *Item
	length int
}

func (queue *Queue) Enqueue(v interface{}) {
	fmt.Println(v)
	newItem := &Item{
		value: v,
		next:  queue.front,
	}
	if queue.isEmpty() {
		queue.front = newItem
		queue.rear = newItem
	} else {
		queue.rear.next = newItem
		queue.rear = newItem
	}
	queue.length = queue.length + 1
}
func (queue *Queue) Len() int {
	return queue.length
}
func (queue *Queue) isEmpty() bool {
	return queue.Len() == 0
}
func (queue *Queue) Dequeue() interface{} {
	if queue.isEmpty() {
		return nil
	}
	valueToDequeue := queue.front.value
	queue.front = queue.front.next
	queue.length--
	return valueToDequeue
}
func (queue *Queue) Front() interface{} {
	if queue.isEmpty() {
		return nil
	}
	return queue.front.value
}
func main() {
	var integerQueue Queue = Queue{}
	integerQueue.Enqueue(100)                                   // 100 is the new queue
	integerQueue.Enqueue(200)                                   // 100-200 is the new queue
	integerQueue.Dequeue()                                      // Dequeues 100 | 200 is the new queue
	fmt.Println("The new Queue front is", integerQueue.Front()) // 100
	integerQueue.Enqueue(300)                                   // 200 - 300 is the new queue
	fmt.Println("The new Queue front is", integerQueue.Front()) // 100
}
