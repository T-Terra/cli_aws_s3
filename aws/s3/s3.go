package s3

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type ParamsInstallWin struct {
	msiexec string
	param string
	url_to_download string
}

type ParamsInstallLinux struct {
	download []string
	zip []string
	install []string
	uninstall []string
}

var command_create_file *string = flag.String("create_file", "", "Read file with path of files S3")
var command_installing_aws_cli *string = flag.String("i", "", "Installing aws cli in linux or windows os (-i=linux/windows)")
var command_read_file *string = flag.String("file", "", "Read file for get data")
var command_help *string = flag.String("h", "", "Show all commands usage")

func Create_file_command() {
	result_verify_dir := Verify_dir()
	if result_verify_dir == "File not exists" {
		newFile, err := os.Create(*command_create_file) // Create a new file
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
		if file.Name() == *command_create_file {
			fmt.Printf("Arquivo existe: %v\n", file.Name())
			txt_result = "File exists"
		} else {
			txt_result = "File not exists"
		}
	}
	return txt_result
}

func Regex_find_args( pattern string, arg_search string) bool {
	result, err := regexp.MatchString(pattern, arg_search)
	if err != nil {
		panic(err)
	}
	return result
}

func Help() string {
	flag.Parse()
	return *command_help
}

func ExeInstallAwsCli() {
	flag.Parse()

	payloadWin := ParamsInstallWin {
		msiexec: "msiexec.exe",
		param: "/i",
		url_to_download: "https://awscli.amazonaws.com/AWSCLIV2.msi",
	}

	payloadLinux := ParamsInstallLinux {
		download: []string{"curl", "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip", "-o", "awscliv2.zip"},
		zip: []string{"unzip", "awscliv2.zip"},
		install: []string{"sudo", "./aws/install"},
		uninstall: []string{"sudo", "rm", "-rf", "aws/"},
	}

	if *command_installing_aws_cli == "linux" {
		cmd1 := exec.Command(payloadLinux.download[0], payloadLinux.download[1], payloadLinux.download[2], payloadLinux.download[3])

		cmd2 := exec.Command(payloadLinux.zip[0], payloadLinux.zip[1])

		cmd3 := exec.Command(payloadLinux.install[0], payloadLinux.install[1])

		remove_dir_aws := exec.Command(payloadLinux.uninstall[0], payloadLinux.uninstall[1], payloadLinux.uninstall[2], payloadLinux.uninstall[3])

		remove_zip_aws := exec.Command(payloadLinux.uninstall[0], payloadLinux.uninstall[1], payloadLinux.zip[1])

		if err1 := cmd1.Run(); err1 != nil {
			fmt.Println("Erro ao baixar o arquivo:", err1)
			os.Exit(1)
		}

		if err2 := cmd2.Run(); err2 != nil {
			fmt.Println("Erro ao descompactar o arquivo:", err2)
			os.Exit(1)
		}

		if err3 := cmd3.Run(); err3 != nil {
			fmt.Println("Erro ao instalar o AWS CLI:", err3)
			os.Exit(1)
		}

		if err4 := remove_dir_aws.Run(); err4 != nil {
			fmt.Println("Erro ao apagar a pasta AWS", err4)
			os.Exit(1)
		}

		if err5 := remove_zip_aws.Run(); err5 != nil {
			fmt.Println("Erro ao apagar o zip AWS", err5)
			os.Exit(1)
		}

		fmt.Println("AWS CLI instalado com sucesso!")
	} else if *command_installing_aws_cli == "windows" {
		cmd1 := exec.Command(payloadWin.msiexec, payloadWin.param, payloadWin.url_to_download)

		cmd1.Stdout = os.Stdout
		cmd1.Stderr = os.Stderr

		err := cmd1.Run()
		if err != nil {
			fmt.Printf("Erro ao executar o instalação: %v\n", err)
			return
		}

		fmt.Println("Comando executado com sucesso!")
	} else {
		fmt.Println("Sistema não encontrado")
	}
}

func ReadFile() []string {
	flag.Parse()
	file, err := os.ReadFile(*command_read_file)
	if err != nil {
		fmt.Println("Erro ao ler o arquivo")
	}
	stringFormat := string(file)
	splitString := strings.Split(stringFormat, "\n")
	return splitString
}
