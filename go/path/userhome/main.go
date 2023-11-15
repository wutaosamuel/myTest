package main
import (
	"os"
	"fmt"
	"log"
)

func main() {
	dirName, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dirName)
}