package setutil

type (
	empty = struct{}

	// TODO: go1.24: use alias
	// set[T comparable] = map[T]empty
	Set[T comparable] map[T]empty
)

func Contains[S ~map[E]empty, E comparable](s S, v E) bool {
	_, found := s[v]
	return found
}

func Add[S ~map[E]empty, E comparable](s S, vs ...E) {
	for _, v := range vs {
		s[v] = empty{}
	}
}

func (s Set[T]) Contains(v T) bool {
	return Contains(s, v)
}

func (s Set[T]) Add(vs ...T) {
	Add(s, vs...)
}
