package main

import (
	"fmt"

	"github.com/gauravmahal/GolangTraining/modules/myFirstDocumentedPackage/mymath"
)

func main() {
	fmt.Println("2 + 3 =", mymath.Sum(2, 3))
}
