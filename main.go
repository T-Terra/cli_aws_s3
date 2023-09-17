package main

import (
	"cli_aws_s3/aws/s3"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		create_file := s3.Regex_find_args(".*create_file.*", os.Args[1])
		install_os := s3.Regex_find_args("-i.*", os.Args[1])
		read_file := s3.Regex_find_args(".*file.*", os.Args[1])

		if create_file {
			s3.Create_file_command()
		} else if install_os {
			fmt.Println("Instalando AWS CLI..")
			s3.ExeInstallAwsCli()
		} else if read_file {
			if len(os.Args) != 3 {
				fmt.Println("Adicione o nome do bucket como segundo argumento")
				fmt.Println("Exemplo: -file=<nome do arquivo> <bucket>")
				return
			} else {
				dados := s3.ReadFile()
				bucket := os.Args[2]
				for i := 0; i < len(dados); i++ {
					s3.ExeCommandAws(dados[i], bucket)
				}
			}
		} else {
			s3.Help()
		}
	} else {
		fmt.Println("use a flag -h para verificar o modo de uso")
	}
}