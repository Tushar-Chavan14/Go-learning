package employee

func Filter(slice []Employee, condition func(Employee) bool) []Employee {

	var result []Employee

	for _, element := range slice {
		if condition(element) {
			result = append(result, element)
		}
	}

	return result
}

func Find(slice []Employee, condition func(Employee) bool) (any, bool) {

	for _, element := range slice {
		if condition(element) {
			return element, true
		}
	}

	return nil, false
}
