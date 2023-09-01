package main

import (
	"cli_aws_s3/aws/s3"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "-file"{
		s3.Create_file_command()
	} else {
		fmt.Println("Executando outra ação")
	}
}