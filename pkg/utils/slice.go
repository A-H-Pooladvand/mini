package utils

func InSlice[A comparable](needle A, collection []A) bool {
	for _, value := range collection {
		if needle == value {
			return true
		}
	}

	return false
}

func Map[A any](collection []A, callback func(value A, key int) A) (items []A) {
	for key, value := range collection {
		items = append(items, callback(value, key))
	}

	return items
}
