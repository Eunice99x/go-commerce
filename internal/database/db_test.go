package database_test

import (
	"gihub.com/eunice99x/go-commerce/src/database"
	"testing"
)

func TestConnection(t *testing.T) {
	err := database.Connect()
	if err != nil {
		t.Errorf("Error connection to database: %v", err)
	}
}
