package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/DKeshavarz/eventic/internal/delivery/telegram"
	"github.com/DKeshavarz/eventic/internal/delivery/web"
	"github.com/joho/godotenv"
)

func TestConfigStructWithNoFile(t *testing.T) {
	cfg := New()

	if cfg.Telegram == nil || *cfg.Telegram != *(telegram.DefaultConfig()) {
		t.Error("Expected telergarm config to be initialized but got", cfg.Telegram)
	}
	
	if cfg.WebServer == nil || *(cfg.WebServer) != *(web.DefaultConfig()) {
		t.Error("Expected web config to be initialized but got", cfg.WebServer)
	}
}

func TestWithFile(t *testing.T) {
	fileContent := fmt.Sprintf("%s=%s  \n", "TELEGRAM_API_KEY", "myTelegramAPI") +
				   fmt.Sprintf("%s=%s  \n", "WEB_PORT", "2000")

	const fileName = ".env.test"

	tmpFile, err := os.Create(fileName)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(fileName)
	if _, err := tmpFile.Write([]byte(fileContent)); err != nil {
		t.Fatal(err)
	}
	tmpFile.Close()
	err = godotenv.Load(fileName)
	if err != nil {
		t.Fatal(err)
	}

	//*************************************************************//
	cfg := New()
	if cfg.Telegram.APIKey != "myTelegramAPI" {
		t.Errorf("Expected telergarm API to be `%s` but get: `%s`", "myTelegramAPI", cfg.Telegram.APIKey)
	}
	if cfg.WebServer.Port != "2000" {
		t.Errorf("Expected telergarm API to be `%s` but get: `%s`", "2000", cfg.WebServer.Port)
	}
}
