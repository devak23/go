package stacker

import "errors"

// Our implementation of stack will be a slice denoted by []
// the interface{} is a empty interface implemented by all data types
type Stack []interface{}

// it is conventional for any custom type to report
// its length or capacity. Therefore...
func (stack Stack) Len() int {
  return len(stack)
}

// and this function returns the capacity of the stack
func (stack Stack) Cap() int {
  return cap(stack)
}

// a utility function that tells us if the stack is empty or not
func (stack Stack) IsEmpty() bool {
  return len(stack) == 0
}

// Returns the topmost element from the stack. Note that this method
// does not modify the stack and hence the paramters are passed by value
func (stack Stack) Peek() (interface{}, error) {
  if stack.IsEmpty() {
    return nil, errors.New("Cant Peek into an empty Stack")
  }

  return stack[len(stack) - 1], nil
}

// This method modifies the stack. Thus the parameters are passed by reference
func (stack *Stack) Push(element interface{}) {
  *stack = append(*stack, element)
}

// you can pop any object from the stack
func (stack *Stack) Pop() (interface{}, error) {
  if stack.IsEmpty() {
    return nil, errors.New("Stack is empty!")
  }

  // backup the current stack object as we are going to change the stack
  currStack := *stack

  // extract the element (topmost) from the stack
  element := currStack[len(currStack) - 1]

  // modify the current stack pointer to point to a new stack (slice)
  // which is obtained by slicing the last element out of it
  *stack = currStack[:len(currStack) - 1]

  return element, nil
}


