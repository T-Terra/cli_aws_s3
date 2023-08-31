package s3

import (
	"fmt"
	"flag"
	"os"
)

func Read_files_command() {
	command := flag.String("file", "", "Read file with path of files S3")
	flag.Parse()
	currentFile, err := os.ReadDir("./")
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(currentFile); i++ {
		if currentFile[i].Name() == *command {
			fmt.Printf("Arquivo existe: %v\n", currentFile[i].Name())
		} else {
			newFile, err := os.Create("path.txt") // Create a new file
			if err != nil {
				fmt.Println(err)
			}
			defer newFile.Close()
			fmt.Println("Arquivo Criado com sucesso!!!")
		}
	}
}