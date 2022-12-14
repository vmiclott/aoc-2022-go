package day11

type queue []*int

func (q *queue) push(item *int) {
	*q = append(*q, item)
}

func (q *queue) pop() *int {
	item := (*q)[0]
	*q = (*q)[1:len(*q)]
	return item
}
