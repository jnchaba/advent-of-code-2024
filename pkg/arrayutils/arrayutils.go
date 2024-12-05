package arrayutils

func Contains[T comparable](arr []T, target T) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}

func IndexOf[T comparable](arr []T, target T) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

func RemoveAtIndex[T comparable](arr []T, index int) (T, []T) {
	element := arr[index]
	transformed := append(arr[:index], arr[index+1:]...)
	return element, transformed
}

func InsertAtIndex[T comparable](arr []T, index int, element T) []T {
	// array before index
	before := arr[:index]
	// array after index
	after := arr[index:]
	// new slice with our desired element preceding the 'after' array
	withElement := append([]T{element}, after...)
	// combine before and withElement
	final := append(before, withElement...)
	return final
}
