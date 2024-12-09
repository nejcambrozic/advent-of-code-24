package utils

func DeepCopy2dStrSpice(original [][]string) [][]string {
	output := make([][]string, len(original))

	for i := range original {
		output[i] = make([]string, len(original[i]))
		copy(output[i], original[i])
	}
	return output
}
