package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type FileHandler struct {
	SharedFilePath string 
}


func (fh *FileHandler) CreateFile() {
	file, err := os.Create(fh.SharedFilePath)
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("0")
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return
	}

	fmt.Println("Arquivo criado com sucesso!")
}

func (fh *FileHandler) GetFileValue() (int, error) {
	content, err := ioutil.ReadFile(fh.SharedFilePath)
	if err != nil {
		return 0, fmt.Errorf("Erro ao ler o arquivo: %v", err)
	}

	value, err := strconv.Atoi(string(content))
	if err != nil {
		return 0, fmt.Errorf("Erro ao converter para inteiro: %v", err)
	}

	return value, nil
}

func (fh *FileHandler) WriteToFile(value int) error {
	err := ioutil.WriteFile(fh.SharedFilePath, []byte(strconv.Itoa(value)), 0644)
	if err != nil {
		return fmt.Errorf("Erro ao escrever no arquivo: %v", err)
	}

	return nil
}

