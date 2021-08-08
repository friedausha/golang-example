package addition

import "fmt"

func Add(a int, b int) (int, error){
	if a < 10 || b < 10 {
		return 0, fmt.Errorf("kurang dari 10")
	}
	return a + b, nil
}
