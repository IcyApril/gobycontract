package gobycontract_test

import (
	"testing"
	"os"
	"github.com/IcyApril/gobycontract"
	"fmt"
	"strconv"
)

func ExampleSecondsToSecondsAndMinutes() {
	minutes, seconds := SecondsToSecondsAndMinutes(125)
	fmt.Println(strconv.Itoa(minutes) + " minutes " + strconv.Itoa(seconds) + " seconds")
	// output:
	// 2 minutes 5 seconds
}

func SecondsToSecondsAndMinutes(seconds int) (minutes int, remainingSeconds int) {
	gobycontract.Require(seconds >= 0, "Input seconds must be positive")

	minutes = seconds/60
	remainingSeconds = seconds % 60

	gobycontract.Ensure(minutes > 0, "Output minutes most be positive")
	gobycontract.Ensure(remainingSeconds > 0, "Output remaining seconds most be positive")
	gobycontract.Ensure(remainingSeconds < 59, "There can be no more than 59 remaining seconds")

	return
}

func TestSumContract(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

func Sum(a int, b int) (result int) {
	gobycontract.Require(a > 0, "Argument a must be > 0")
	gobycontract.Require(b > 0, "Argument b must be > 0")

	result = sum(a, b)

	gobycontract.Ensure(result == (a + b), "Return value must be a + b")

	return
}

func sum(a int, b int) int {
	return a + b
}

func TestBrokenSumContract(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	os.Unsetenv("GOBYCONTRACT_DONTPANIC")

	total := BrokenSum(5, 5)
	if total != 0 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

func TestBrokenSumContractNoPanic(t *testing.T) {
	os.Setenv("GOBYCONTRACT_DONTPANIC", "true")

	total := BrokenSum(5, 5)
	if total != 0 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

func BrokenSum(a int, b int) (result int) {
	gobycontract.Require(a > 0, "Argument a must be > 0")
	gobycontract.Require(b > 0, "Argument b must be > 0")

	result = brokenSum(a, b)

	gobycontract.Ensure(result == (a + b), "Return value must be a + b")

	return
}

func brokenSum(a int, b int) int {
	return a - b
}