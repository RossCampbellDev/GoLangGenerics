package main

import "golang.org/x/exp/constraints"  // a set of lots of var types eg strings, numbers

type Num interface {
  int | int8 | int16 | float32 | float64
}

// this function can take 2 variables of any type from the Num interface
// and return a value of that same type
func Add[T Num](a T, b T) T {
  return a + b
}

// now do a mapping function with generic typing!
func MapThis(values []T, mapFunc func(T) T) []T {
  var newValues = []T
  for _, v := range values {
    newVal := mapFunc(v)
    newValues = append(newValues, newVal)
  }
}



type CustomAttr interface {
  constraints.Ordered
}

type User[T CustomAttr] struct {
  ID int
  Name string
  SomeData T
}

func main() {
  // make use of our generic map function with an anonymous function that will square vals
  result := MapThis([]int{1,2,3}. func(n int) int {
    return n * 2
  })
  
  // specify that the SomeData attribute must be string type this time
  u := User[string]{
    ID: 0,
    Name: "Ross",
    SomeData: "this must be a string"
  }
}
