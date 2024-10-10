package utils

import (
	"fmt"
	"math/rand"

	"github.com/cockroachdb/apd/v3"
)

func printRandomNumber() {
	fmt.Println("This is a random number", rand.Int31n(200))
}

func DivideNumbers(arg1 int64, arg2 int64) {
	a := apd.NewBigInt(arg1)
	b := apd.NewBigInt(arg2)
	divide(a, b)
}

// Divide and store in arg1 or smth
func divide(arg1 *apd.BigInt, arg2 *apd.BigInt) {
	// Was
	fmt.Println("The result of", arg1, "divided by", arg2, "is", arg1.Div(arg1, arg2))
}
