package main

import (
	"testing"

	"github.com/amirmarcel/dotenv-linter/diff"
	"github.com/amirmarcel/dotenv-linter/parser"
)

func TestParseEnvFile(t *testing.T) {
	keys, err := parser.ParseEnvFile("testdata/.env.example")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !keys["DATABASE_URL"] {
		t.Error("expected DATABASE_URL to be present")
	}
	if !keys["PORT"] {
		t.Error("expected PORT to be present")
	}
	if len(keys) != 3 {
		t.Errorf("expected 3 keys, got %d", len(keys))
	}
}

func TestDiffKeys(t *testing.T) {
	example := map[string]bool{
		"DATABASE_URL": true,
		"PORT": true,
		"SECRET_KEY": true,
	}
	local := map[string]bool{
		"DATABASE_URL": true,
		"PORT": true,
		"DEBUG_MODE": true,
	}

	missingInLocal, missingInExample := diff.DiffKeys(example, local)

	if len(missingInLocal) != 1 || missingInLocal[0] != "SECRET_KEY" {
		t.Errorf("expected SECRET_KEY missing in local, got %v", missingInLocal)
	}
	if len(missingInExample) != 1 || missingInExample[0] != "DEBUG_MODE" {
		t.Errorf("expected DEBUG_MODE missing in local, got %v", missingInExample)
	}
}