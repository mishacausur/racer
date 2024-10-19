package printer

import "fmt"

func PrintHello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func main() {
	print(PrintHello("fgrhg"))
}
