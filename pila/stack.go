package pila

import (
	"encoding/json"
	"fmt"

	"github.com/fern4lvarez/piladb/pkg/stack"
	"github.com/fern4lvarez/piladb/pkg/uuid"
)

// Stack represents a stack entity in piladb.
type Stack struct {
	// ID is a unique identifier of the Stack
	ID fmt.Stringer
	// Name of the Stack
	Name string
	// Database associated to the Stack
	Database *Database

	// base represent the Stack data structure
	base *stack.Stack
}

// NewStack creates a new Stack given a name without
// an association to any Database.
func NewStack(name string) *Stack {
	s := &Stack{}
	s.Name = name
	s.SetID()
	s.base = stack.NewStack()
	return s
}

// Push an element on top of the Stack.
func (s *Stack) Push(element interface{}) {
	s.base.Push(element)
}

// Pop removes and returns the element on top of the Stack.
//If the Stack was empty, it returns false.
func (s *Stack) Pop() (interface{}, bool) {
	return s.base.Pop()
}

// Size returns the size of the Stack.
func (s *Stack) Size() int {
	return s.base.Size()
}

// Peek returns the element on top of the Stack.
func (s *Stack) Peek() interface{} {
	return s.base.Peek()
}

// SetDatabase links the Stack with a given Database and
// recalculates its ID.
func (s *Stack) SetDatabase(db *Database) {
	s.Database = db
	s.SetID()
}

// SetID recalculates the id of the Stack based on its
// Database name and its own name.
func (s *Stack) SetID() {
	if s.Database != nil {
		s.ID = uuid.New(s.Database.Name + s.Name)
		return
	}

	s.ID = uuid.New(s.Name)
}

// Status returns the status of the Stack  in json format.
func (s *Stack) Status() StackStatus {
	status := StackStatus{}
	status.ID = s.ID.String()
	status.Name = s.Name
	status.Size = s.Size()
	status.Peek = s.Peek()

	return status
}

// StackStatus represents the status of a Stack.
type StackStatus struct {
	ID   string      `json:"id"`
	Name string      `json:"name"`
	Peek interface{} `json:"peek"`
	Size int         `json:"size"`
}

// ToJSON converts a StackStatus into JSON.
func (stackStatus StackStatus) ToJSON() ([]byte, error) {
	return json.Marshal(stackStatus)
}
