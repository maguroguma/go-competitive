package deque

type Deque struct {
	front, back []int
}

func NewDeque() *Deque {
	d := new(Deque)
	d.front, d.back = []int{}, []int{}
	return d
}

func (d *Deque) Len() int {
	return len(d.front) + len(d.back)
}

func (d *Deque) PushFront(a int) {
	d.front = append(d.front, a)
}

func (d *Deque) PushBack(a int) {
	d.back = append(d.back, a)
}

func (d *Deque) PopFront() int {
	if len(d.front) > 0 {
		pop := d.front[len(d.front)-1]
		d.front = d.front[:len(d.front)-1]
		return pop
	}

	if len(d.back) == 0 {
		panic("deque is empty")
	}

	pop := d.back[0]
	d.back = d.back[1:]
	return pop
}

func (d *Deque) PopBack() int {
	if len(d.back) > 0 {
		pop := d.back[len(d.back)-1]
		d.back = d.back[:len(d.back)-1]
		return pop
	}

	if len(d.front) == 0 {
		panic("deque is empty")
	}

	pop := d.front[0]
	d.front = d.front[1:]
	return pop
}

func (d *Deque) List() []int {
	res := []int{}
	for i := len(d.back) - 1; i >= 0; i-- {
		res = append(res, d.back[i])
	}
	for i := 0; i < len(d.front); i++ {
		res = append(res, d.front[i])
	}
	return res
}
