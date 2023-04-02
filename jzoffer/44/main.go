package main

import "strconv"

func main() {

}

func findNthDigit(n int) int {

	//int digit = 1;   // n所在数字的位数
	//long start = 1;  // 数字范围开始的第一个数
	//long count = 9;  // 占多少位
	//while(n > count){
	//	n -= count;
	//	digit++;
	//	start *= 10;
	//	count = digit * start * 9;
	//}
	//long num = start + (n - 1) / digit;
	//return Long.toString(num).charAt((n - 1) % digit) - '0';

	digit := 1
	start := 1
	count := 9

	for n > count {
		n -= count
		digit++
		start *= 10
		count = digit * start * 9
	}
	num := start * (n - 1) / digit
	return int(strconv.Itoa(num)[(n-1)%digit] - '0')
}
