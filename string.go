package uid

type StringProvider[T uid] struct {
	s T
}

func (c *StringProvider[T]) Value() T {
	return c.s
}


func NewStringProvider[T uid](s T) (UID[T], error) {
	pr := &StringProvider[T]{s: s}
	return pr, nil
}
