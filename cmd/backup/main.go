package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func backupPostgres(dbConfig map[string]string) error {
	backupFile := fmt.Sprintf("backup_%s.sql", time.Now().Format("20060102_150405"))

	// Формируем команду pg_dump
	cmd := exec.Command(
		"pg_dump",
		"-h", dbConfig["host"],
		"-p", dbConfig["port"],
		"-U", dbConfig["username"],
		"-d", dbConfig["dbname"],
		"-f", backupFile,
	)

	cmd.Env = append(os.Environ(), fmt.Sprintf("PGPASSWORD=%s", dbConfig["password"]))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("ошибка резервного копирования: %w", err)
	}

	fmt.Printf("Бэкап сохранен в %s\n", backupFile)
	return nil
}

func main() {
	dbConfig := map[string]string{
		"host":     "localhost",
		"port":     "54321",
		"username": "question-user",
		"password": "question-password",
		"dbname":   "question",
	}

	if err := backupPostgres(dbConfig); err != nil {
		panic(err)
	}
}
