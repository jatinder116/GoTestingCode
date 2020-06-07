package main

import (
	"fmt"
    "GoTestingCode/testPkg"
)

func main() {
    // ============= square and cube =========================
	a, b := testPkg.Squareandcube(19)
    fmt.Println( "square of a number is:",a);
    fmt.Println("cube is a number is:",b);

    // ============== Print N prime Numbers ======================
    p := testPkg.Primenumbers(100)
    fmt.Println( "The N prime numbers is:",p);
}