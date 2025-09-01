package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/Kosfedev/learn_go/internal/config"
)

const (
	backupPath = "db/backups/"
)

func backupPostgres() error {
	backupFile := fmt.Sprintf("%sbackup_%s.sql", backupPath, time.Now().Format("20060102_150405"))

	cmd := exec.Command("pg_dump", "-f", backupFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("ошибка резервного копирования: %w", err)
	}

	fmt.Printf("Бэкап сохранен в %s\n", backupFile)
	return nil
}

// TODO: дублирует app.go
func initConfig() error {
	mode := "dev"
	if os.Getenv("MODE") == "prod" {
		mode = "prod"
	}

	err := config.Load(fmt.Sprintf(".env.%s", mode))
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := initConfig()
	if err != nil {
		log.Fatal(err)
	}

	for {
		if err = backupPostgres(); err != nil {
			log.Fatal(err)
		}

		time.Sleep(1 * time.Minute)
	}
}
