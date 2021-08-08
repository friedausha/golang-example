package multiplier

import "fmt"

func Multiply(a int, b int) (int, error){
	if a < b {
		return 0, fmt.Errorf("%d kurang dari %d", a, b)
	}
	return a * b, nil
}
