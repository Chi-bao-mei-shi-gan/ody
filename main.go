package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	var newName string
	var showHelp bool

	// 解析命令行參數
	flag.StringVar(&newName, "n", "", "The new name to replace 'newProject'")
	flag.BoolVar(&showHelp, "h", false, "Show help message")
	flag.Parse()

	if showHelp {
		fmt.Println("Usage: ody [-n newName] [-h]")
		fmt.Println()
		fmt.Println("Options:")
		fmt.Println("  -n string")
		fmt.Println("        The new name to replace 'newProject'")
		fmt.Println("  -h")
		fmt.Println("        Show help message")
		return
	}

	if newName == "" {
		log.Fatal("You must specify a new name with -n")
	}

	// 獲取當前工作目錄
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// 克隆存儲庫到當前目錄
	cloneDir := filepath.Join(currentDir, newName)
	if err := os.MkdirAll(cloneDir, 0755); err != nil {
		log.Fatal(err)
	}
	if err := exec.Command("git", "clone", "https://github.com/Chi-bao-mei-shi-gan/new_project.git", cloneDir).Run(); err != nil {
		log.Fatal(err)
	}

	// 刪除 .git 目錄
	err = os.RemoveAll(filepath.Join(cloneDir, ".git"))
	if err != nil {
		log.Fatal(err)
	}

	// 遞歸地遍歷所有文件並替換 'newProject' 為 newName
	if err := filepath.Walk(cloneDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			// 讀取文件內容
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			// 替換內容
			newContent := bytes.Replace(content, []byte("newProject"), []byte(newName), -1)
			// 寫回文件
			if err := ioutil.WriteFile(path, newContent, info.Mode()); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done! Your new project is ready at", cloneDir)
}
