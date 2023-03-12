package main

func main() {

}

func myPow(x float64, n int) float64 {
	if n < 0 {

		return 1 / myPow(x, -n)
	}
	if n == 0 {
		return 1
	}

	if n%2 == 0 {
		y := myPow(x, n/2)
		return y * y
	}

	return x * myPow(x, n-1)
}
