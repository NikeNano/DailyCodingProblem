package TwoNumbersAddingUp

func bruteForce(inputValues []int, answer int) bool {
	//m = make(map[string]int)
	for indexOuter, valueOuter := range inputValues {
		for indexInter, valueInner := range inputValues {
			if indexOuter != indexInter { // edge case, not multiply with it self.
				if valueOuter+valueInner == answer {
					return true
				}
			}
		}
	}
	return false

}

func bruteForceDynamic(inputValues []int, answer int) bool {
	for indexOuter, valueOuter := range inputValues {
		if valueOuter > answer { // Since the value is to large allready ...
			continue
		}
		for indexInter, valueInner := range inputValues[indexOuter:] { // Since we have done it allready
			if indexOuter != indexInter { // edge case, not multiply with it self.
				if valueOuter+valueInner == answer {
					return true
				}
			}
		}
	}
	return false

}
