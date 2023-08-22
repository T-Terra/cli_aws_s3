package main

import (
	"os"
	"cli_aws_s3/aws/s3"
)

func main() {
	var arguments []string

	arguments = append(arguments, os.Args[1])
	arguments = append(arguments, os.Args[2])

	s3.Read_args(arguments)
}