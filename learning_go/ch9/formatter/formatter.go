package formatter

import "fmt"

// Format returns a descriptive string representation of an int
func Format(num int) string {
	return fmt.Sprintf("The number is %d", num)
}
