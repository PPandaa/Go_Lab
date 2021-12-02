package tool

func RemoveElementFromSlice(elements []interface{}, target interface{}) []interface{} {

	for elementIndex, element := range elements {
		if element == target {
			return append(elements[:elementIndex], elements[elementIndex+1:]...)
		}
	}
	return elements

}
