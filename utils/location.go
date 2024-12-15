package utils

type Location struct {
	I int
	J int
}

func IsValidLocation(loc Location, gridSize int) bool {
	return loc.I >= 0 && loc.I < gridSize && loc.J >= 0 && loc.J < gridSize
}
