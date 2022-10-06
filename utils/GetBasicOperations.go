package utils

func GetBasicOperations(a, b int) (int, int, int, float32) {
	var addition int = getAddition(a, b)
	var subtraction int = getSubtraction(a, b)
	var multiplication int = getMultiplication(a, b)
	var division float32 = getDivision(a, b)

	return addition, subtraction, multiplication, division
}

func getAddition(a, b int) int {
	return a + b
}
func getSubtraction(a, b int) int {
	return a - b
}
func getMultiplication(a, b int) int {
	return a * b
}
func getDivision(a, b int) float32 {
	newA := float32(a)
	newB := float32(b)
	return newA / newB
}
