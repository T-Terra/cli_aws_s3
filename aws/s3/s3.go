package s3

import (
	"fmt"
	"flag"
	"os"
)

var command *string = flag.String("file", "", "Read file with path of files S3")

func Create_file_command() {
	result_verify_dir := Verify_dir()
	if result_verify_dir == "File not exists" {
		newFile, err := os.Create(*command) // Create a new file
		if err != nil {
			fmt.Println(err)
		}
		defer newFile.Close()
		fmt.Println("Arquivo Criado com sucesso!!!")
	}
}

func Verify_dir() ( txt_result string) {
	flag.Parse()
	currentFile, err := os.ReadDir("./")
	if err != nil {
		panic(err)
	}
	for _, file := range currentFile {
		if file.Name() == *command {
			fmt.Printf("Arquivo existe: %v\n", file.Name())
			txt_result = "File exists"
		} else {
			txt_result = "File not exists"
		}
	}
	return txt_result
}
