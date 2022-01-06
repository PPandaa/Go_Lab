package tool

func RemoveElementFromSlice(elements []interface{}, target interface{}) []interface{} {

	for elementIndex, element := range elements {
		if element == target {
			return append(elements[:elementIndex], elements[elementIndex+1:]...)
		}
	}
	return elements

}

func FindDiffFromSlice(oldSlice []string, newSlice []string) ([]interface{}, []interface{}) {

	for oldIndex, old := range oldSlice {
		for newIndex, new := range newSlice {
			if old == new {
				oldSlice[oldIndex] = "Psame"
				newSlice[newIndex] = "Psame"
			}
		}
	}

	missingElements := []interface{}{}
	for _, old := range oldSlice {
		if old != "Psame" {
			missingElements = append(missingElements, old)
		}
	}
	newElements := []interface{}{}
	for _, new := range newSlice {
		if new != "Psame" {
			newElements = append(newElements, new)
		}
	}

	return missingElements, newElements

}
