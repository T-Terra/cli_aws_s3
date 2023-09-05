package s3

import (
	"fmt"
	"flag"
	"os"
	"runtime"
	"os/exec"
)

type ParamsInstallWin struct {
	msiexec string
	param string
	url_to_download string
}

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


func ExeInstallAwsCli() {
	payload := ParamsInstallWin {
		msiexec: "msiexec.exe",
		param: "/i",
		url_to_download: "https://awscli.amazonaws.com/AWSCLIV2.msi",
	}
	
	if runtime.GOOS == "linux" {
		cmd1 := exec.Command("curl", "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip", "-o", "awscliv2.zip")

		// Comando 2: Descompactar o arquivo awscliv2.zip
		cmd2 := exec.Command("unzip", "awscliv2.zip")

		// Comando 3: Instalar o AWS CLI
		cmd3 := exec.Command("sudo", "./aws/install")

		// Execute os comandos em sequência
		if err := cmd1.Run(); err != nil {
			fmt.Println("Erro ao baixar o arquivo:", err)
			os.Exit(1)
		}

		if err := cmd2.Run(); err != nil {
			fmt.Println("Erro ao descompactar o arquivo:", err)
			os.Exit(1)
		}

		if err := cmd3.Run(); err != nil {
			fmt.Println("Erro ao instalar o AWS CLI:", err)
			os.Exit(1)
		}

		fmt.Println("AWS CLI instalado com sucesso!")
	} else if runtime.GOOS == "windows" {
		// Comando a ser executado
		cmd1 := exec.Command(payload.msiexec, payload.param, payload.url_to_download)

		// Configurar a saída padrão e erro para a saída do seu programa
		cmd1.Stdout = os.Stdout
		cmd1.Stderr = os.Stderr

		// Executar o comando
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