package nomor_empat

import "fmt"

func PrintSquareRoot(arr []int) {
	for _, v := range arr {
		ans, err := FindExponentOfTwo(v)
		if err != nil{
			fmt.Println(err)
			continue
		}
		words := fmt.Sprintf("2**%d = %d\n", ans, v)
		fmt.Println(words)
	}
}

func FindExponentOfTwo(val int) (res int, err error) {
	temp := 1
	for i := 2; i<=val; i++ {
		temp *= 2
		res++
		if temp > val {
			return 0, fmt.Errorf("bukan pangkat dari 2")
		}
		if temp == val {
			return res, nil
		}
	}
	return res, nil
}