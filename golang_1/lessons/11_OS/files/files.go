package files

import (
	"fmt"
	"os"
)

func ReadFile(name string) ([]byte, error) {
	// file, err := os.Open("files.txt")
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil

}

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Print(err)
	}

	_, err = file.Write(content)
	defer file.Close()

	if err != nil {

		fmt.Println(err)
		return
	}
	fmt.Println("Запись успешна")

}
