package utils

// Stack represents a dynamic stack of bool values
type Stack []bool

// Push appends a value to the stack
func (s *Stack) Push(value bool) {
	*s = append(*s, value)
}

// Pop removes and returns the last value from the stack
func (s *Stack) Pop() (bool, bool) {
	if len(*s) == 0 {
		// Handle empty stack
		return false, false
	}
	last := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return last, true
}

// Peek returns the last value from the stack without removing it
func (s *Stack) Peek() (bool, bool) {
	if len(*s) == 0 {
		// Handle empty stack
		return false, false
	}
	return (*s)[len(*s)-1], true
}

// IsEmpty returns true if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// InvertCurrent inverts the last value in the stack
func (s *Stack) InvertCurrent() {
	if len(*s) == 0 {
		// Handle empty stack
		return
	}
	(*s)[len(*s)-1] = !(*s)[len(*s)-1]
}
