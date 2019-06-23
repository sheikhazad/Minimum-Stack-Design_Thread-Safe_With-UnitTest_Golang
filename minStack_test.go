package minimumStack

import (
	"fmt"
	"testing"
)

func TestMinStack(test *testing.T) {

	var ms MinStack
	//Or, ms := Constructor()

	//Push int values
	ms.Push(-1)
	ms.Push(2)
	ms.Push(0)
	ms.Push(-4)

	//Stack Size() test
	if ms.Size() != 4 {
		test.Error("Test failed: Expected stack size = 4 from [-1, 2, 0, -4]")
	} else { //Just for printing
		fmt.Println("Test passed: Expected stack size = 4 from [-1, 2, 0, -4]")
	}

	//Stack Top() test
	top, _ := ms.Top()
	if top != -4 {
		test.Errorf("Test failed: Expected top of stack = -4 from [-1, 2, 0, -4]")
	} else { //Just for printing
		fmt.Println("Test passed: Expected top of stack = -4 from [-1, 2, 0, -4]")
	}

	//getMin() test
	minVal, _ := ms.getMin()
	if minVal != -4 {
		test.Errorf("Test failed: Expected min from stack = -4 from [-1, 2, 0, -4]")
	} else { //Just for printing
		fmt.Println("Test passed: Expected min from stack = -4 from [-1, 2, 0, -4]")
	}

	//Pop() test
	ms.Pop() //[-1, 2, 0]
	top, _ = ms.Top()
	if top != 0 {
		test.Errorf("Test failed: After pop, Expected top of stack = 0 from [-1, 2, 0]")
	} else { //Just for printing
		fmt.Println("Test passed: After pop, Expected top of stack = 0 from [-1, 2, 0]")
	}

	//getMin() test
	minVal, _ = ms.getMin()
	if minVal != -1 {
		test.Errorf("Test failed: After pop, Expected min from stack = -1 from [-1, 2, 0]")
	} else { //Just for printing
		fmt.Println("Test passed: After pop, Expected min from stack = -1 from [-1, 2, 0]")
	}

	// Clear() and isEmpty() test
	ms.Clear()
	if ms.isEmpty() != true {
		test.Errorf("Test failed: After Clear(), Stack is expected to be empty.")
	} else { //Just for printing
		fmt.Println("Test passed: After Clear(), Stack is expected to be empty.")
	}
}
