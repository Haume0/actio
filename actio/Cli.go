package actio

import (
	"fmt"
	"os/exec"
)

// Exec : Execute a shell command
func Exec(command string) {
	cmd := exec.Command("sh", "-c", command)
	//komutu çalıştırıp sürekli çıktıyı dinlemek için
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}
	if err := cmd.Start(); err != nil {
		fmt.Println(err)
	}
	for {
		buf := make([]byte, 1024)
		_, err := stdout.Read(buf)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print(string(buf))
	}
}
