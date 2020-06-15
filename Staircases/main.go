package first

import "fmt"

func main() {
	/*
		N = 1: [1]
		N = 2: [1, 1], [2]
		N = 3: [1, 2], [1, 1, 1], [2, 1]
		N = 4: [1, 1, 2], [2, 2], [1, 2, 1], [1, 1, 1, 1], [2, 1, 1]

		In order to get to N=3 we first need to get to N=2 and so on.
		Thus f(3) = f(2) + f(1). This generalizes to  f(n) = f(n - 1) + f(n - 2).
	*/
	fmt.Println("Hello world")
}
