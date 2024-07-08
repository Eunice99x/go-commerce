package database_test

import (
	"github.com/eunice99x/go-commerce/internal/database"
	"testing"
)

func TestConnection(t *testing.T) {
	err := database.Connect()
	if err != nil {
		t.Errorf("Error connection to database: %v", err)
	}
}
