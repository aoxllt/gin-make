package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/cheggaaa/pb/v3"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:     "gin-make",
		Aliases: []string{"gm"},
		Short:   "A CLI tool to generate Gin project structures",
	}

	var generateProjectName string
	rootCmd.PersistentFlags().StringVarP(&generateProjectName, "generate", "g", "", "Generate a new Gin project structure with the given name")
	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		if generateProjectName != "" {
			err := generateProject(generateProjectName)
			if err != nil {
				fmt.Printf("Error generating project: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("Project '%s' generated successfully!\n", generateProjectName)
		}
	}

	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(completionCmd)
	rootCmd.Execute()
}

var generateCmd = &cobra.Command{
	Use:     "generate [project-name]",
	Aliases: []string{"g"},
	Short:   "Generate a new Gin project structure",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		err := generateProject(projectName)
		if err != nil {
			fmt.Printf("Error generating project: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Project '%s' generated successfully!\n", projectName)
	},
}

var completionCmd = &cobra.Command{
	Use:     "completion",
	Aliases: []string{"c"},
	Short:   "Generate the autocompletion script for the specified shell",
}

func generateProject(projectName string) error {
	fmt.Printf("Starting generation of project '%s'...\n", projectName)

	// 步骤 1：创建根目录
	fmt.Print("Creating project root directory... ")
	if err := os.Mkdir(projectName, 0755); err != nil {
		fmt.Println("Failed")
		return err
	}
	fmt.Println("Done")
	basePath := projectName

	// 步骤 2：创建子目录
	dirs := []string{
		"cmd", "config", "internal/app", "internal/database", "internal/models",
		"internal/handlers", "internal/middleware", "internal/repository",
		"internal/services", "pkg/utils", "tests",
	}
	bar := pb.StartNew(len(dirs))
	bar.SetWriter(os.Stdout)
	for _, dir := range dirs {
		if err := os.MkdirAll(filepath.Join(basePath, dir), 0755); err != nil {
			bar.Finish()
			return err
		}
		bar.Increment()
		time.Sleep(50 * time.Millisecond) // 模拟延时，可移除
	}
	bar.Finish()

	// 步骤 3：生成文件
	files := map[string]string{
		"cmd/main.go": fmt.Sprintf(`package main

import (
    "%s/config"
    "%s/internal/app"
)

func main() {
    cfg := config.LoadConfig()
    app.Run(cfg)
}
`, projectName, projectName),

		"config/config.go": `package config

import "github.com/spf13/viper"

type Config struct {
    Port     string ` + "`mapstructure:\"port\"`" + `
    Database string ` + "`mapstructure:\"database\"`" + `
}

func LoadConfig() *Config {
    viper.SetConfigFile("config/config.yaml")
    viper.ReadInConfig()
    cfg := &Config{}
    viper.Unmarshal(cfg)
    return cfg
}
`,

		"config/config.yaml": `port: ":8000"
database: "sqlite://app.db"
`,

		"internal/app/app.go": fmt.Sprintf(`package app

import (
    "github.com/gin-gonic/gin"
    "%s/config"
)

func Run(cfg *config.Config) {
    r := gin.Default()
    RegisterRoutes(r)
    r.Run(cfg.Port)
}
`, projectName),

		"internal/app/routes.go": fmt.Sprintf(`package app

import (
    "github.com/gin-gonic/gin"
    "%s/internal/handlers"
)

func RegisterRoutes(r *gin.Engine) {
    r.GET("/ping", handlers.Ping)
}
`, projectName),

		"internal/handlers/ping.go": `package handlers

import "github.com/gin-gonic/gin"

func Ping(c *gin.Context) {
    c.JSON(200, gin.H{"message": "pong"})
}
`,

		"go.mod": fmt.Sprintf(`module %s

go 1.21

require (
    github.com/gin-gonic/gin v1.9.1
    github.com/spf13/viper v1.19.0
)
`, projectName),
	}

	bar = pb.StartNew(len(files))
	bar.SetWriter(os.Stdout)
	for filePath, content := range files {
		if err := ioutil.WriteFile(filepath.Join(basePath, filePath), []byte(content), 0644); err != nil {
			bar.Finish()
			return err
		}
		bar.Increment()
		time.Sleep(50 * time.Millisecond) // 模拟延时，可移除
	}
	bar.Finish()

	// 步骤 4：初始化 Go 模块
	fmt.Print("Initializing Go module... ")
	if err := os.Chdir(basePath); err != nil {
		fmt.Println("Failed")
		return err
	}
	if err := exec.Command("go", "mod", "tidy").Run(); err != nil {
		fmt.Println("Failed")
		return err
	}
	fmt.Println("Done")

	return nil
}
