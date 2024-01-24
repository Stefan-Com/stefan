package stefan_golang_library

import "fmt"

const (
	color = "yello"
	age   = 40
)

// Will return Hello + the name parameter
func GetHelloMsg(name string) string {
	return fmt.Sprintf("Hello %v!", name)
}
