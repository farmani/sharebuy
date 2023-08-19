package models

import "testing"

func TestNewUser(t *testing.T) {
	tu, err := NewUser("test@gmail.com", "somepass")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if tu.Username == "" {
		t.Errorf("Expected username to be 'test', got %s", tu.Username)
	}

	if tu.Password.Plaintext == nil {
		t.Errorf("Expected password to be set")
	}

	if tu.Password.Hash == nil {
		t.Errorf("Expected password hash to be set")
	}

	if tu.Email != "test@gmail.com" {
		t.Errorf("Expected email to be 'test', got %s", tu.Email)
	}
}
