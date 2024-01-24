package stefan

import "fmt"

const (
	Color = "yello"
	Age   = 40
)

// Will return Hello + the given name
func GetHelloMsg(name string) string {
	return fmt.Sprintf("Hello %v!", name)
}
