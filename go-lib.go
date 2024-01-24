package lib

import "fmt"

const (
	Color = "yello"
	Age   = 40
)

// Will return Hello + the name parameter
func GetHelloMsg(name string) string {
	return fmt.Sprintf("Hello %v!", name)
}
