package main

import (
	"cli_aws_s3/aws/s3"
	"fmt"
	"os"
	"regexp"
)

func main() {
	result, err := regexp.MatchString("-file.*", os.Args[1])
	if err != nil {
		panic(err)
	}
	if len(os.Args) > 1 && result {
		s3.Create_file_command()
	} else {
		fmt.Println("Executando outra ação")
		s3.ExeInstallAwsCli()
	}
}