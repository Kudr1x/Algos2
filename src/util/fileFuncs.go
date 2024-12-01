package util

import (
	"fmt"
	"os"
)

func Write(data string, name string) {
	file, err := os.OpenFile("/home/kudrix/GolandProjects/Algos2/raw/"+name+".txt", os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}

	defer file.Close()

	_, err = file.WriteString(data + "\n")

	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
		return
	}
}

func Clear(name string) {
	file, err := os.OpenFile("/home/kudrix/GolandProjects/Algos2/raw/"+name+".txt", os.O_WRONLY|os.O_TRUNC, 0644)

	if err != nil {
		fmt.Printf("ошибка при открытии файла: %v", err)
		return
	}

	defer file.Close()

	if err := file.Truncate(0); err != nil {
		fmt.Printf("ошибка при обрезке файла: %v", err)
		return
	}
}
