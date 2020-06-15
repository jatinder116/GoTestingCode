package main

import (
	"fmt"
    "gotestcode/testpkg"
)

func mains() {
    // ============= square and cube =========================
	var a, b int = testpkg.Sqrandcube(19)
    fmt.Println( "square of a number is:",a)
    fmt.Println("cube of a number is:",b)

    // ============== Print N prime Numbers ======================
    var p string = testpkg.Primenum(100)
    fmt.Println( "The N prime numbers are:",p)
}