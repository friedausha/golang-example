package concat_string

import "fmt"

func ConcatString(a string, b string) (string, string, string, error) {
	if a == "aku" {
		return "", a, b, fmt.Errorf("engga boleh aku")
	}
	if a == b {
		return "", a, b, fmt.Errorf("string sama yaitu %s", a)
	}
	return a + b, a, b, nil
}
