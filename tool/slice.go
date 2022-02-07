package tool

func RemoveElementFromSlice(elements []interface{}, target interface{}) ([]interface{}, bool) {

	for elementIndex, element := range elements {
		if element == target {
			return append(elements[:elementIndex], elements[elementIndex+1:]...), true
		}
	}
	return elements, false

}

func FindDiffFromStringSlice(oldSlice []string, newSlice []string) ([]interface{}, []interface{}) {

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

func GetDistinctStringSlice(elements []string) []string {

	distinctSlices := []string{}

	for _, elelement := range elements {
		if !IsStringDuplicate(elelement, distinctSlices) {
			distinctSlices = append(distinctSlices, elelement)
		}
	}

	return distinctSlices

}
