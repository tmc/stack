package stack

type Stack[T any] struct {
	Items []T
}

func (s *Stack[T]) Push(x T) {
	s.Items = append(s.Items, x)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.Items) == 0 {
		var zero T
		return zero, false
	}
	v := s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]
	return v, true
}
