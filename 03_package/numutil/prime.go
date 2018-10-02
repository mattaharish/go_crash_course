package numutil

func Prime(n int) bool {
	ctr := 0
	for i := 2; i < n; i++ {
		if n%i == 0 {
			ctr++
		}
	}
	if ctr > 0 {
		return false
	} else {
		return true
	}
}

func Odd(n int) bool {
	if n%2 == 0 {
		return false
	} else {
		return true
	}
}

func Even(n int) bool {
	if n%2 == 0 {
		return true
	} else {
		return false
	}
}

func SumOfArray(arr []int) int {

	sum := 0
	for index := 0; index < len(arr); index++ {
		sum = sum + arr[index]
	}
	return sum
}

func AvgOfArray(arr []float32) float32 {
	var sum float32 = 0.0

	for _, num := range arr {
		sum += num
	}
	return sum / float32(len(arr))
}
