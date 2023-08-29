package main

import (
	"cli_aws_s3/aws/s3"
	"os"
)

func main() {
	s3.Read_files_command()
	subArgs := len(os.Args[2:])
	if subArgs > 1 {
		s3.Sub_commands()
	}
}