package s3

import (
	"fmt"
	"flag"
	"os"
)

func Read_files_command() {
	command := flag.String("file-read", "", "Read file with path of files S3")
	flag.Parse()
	fmt.Printf("Os argumentos são: %v\n", (*command))
}

func Sub_commands() {
	addCmd := flag.NewFlagSet("add", flag.PanicOnError)

	n1 := addCmd.Float64("n1", 0, "Número")

	addCmd.Parse(os.Args[2:])

	fmt.Println(*n1)

}