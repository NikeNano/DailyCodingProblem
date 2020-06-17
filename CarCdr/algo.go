package CarCdr

func cons (a int, b int ) func() {
	func pair (function func) int{
		return function(a, b)
	}
	return pair
}