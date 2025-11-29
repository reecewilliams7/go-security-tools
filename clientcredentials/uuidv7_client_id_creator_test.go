package clientcredentials

import (
	"strings"
	"testing"
)

func TestNewUUIDv7ClientIDCreator(t *testing.T) {
	creator := NewUUIDv7ClientIDCreator()
	if creator == nil {
		t.Fatal("expected non-nil UUIDv7ClientIDCreator")
	}
}

func TestUUIDv7ClientIDCreator_Create(t *testing.T) {
	creator := NewUUIDv7ClientIDCreator()

	id, err := creator.Create()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if id == "" {
		t.Error("expected non-empty client ID")
	}
	// UUIDv7 format should be 36 characters (xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx)
	if len(id) != 36 {
		t.Errorf("expected UUID length 36, got %d", len(id))
	}
	// Check UUID format with dashes
	parts := strings.Split(id, "-")
	if len(parts) != 5 {
		t.Errorf("expected 5 UUID parts, got %d", len(parts))
	}
}

func TestUUIDv7ClientIDCreator_Create_Uniqueness(t *testing.T) {
	creator := NewUUIDv7ClientIDCreator()
	seen := make(map[string]bool)

	for i := 0; i < 100; i++ {
		id, err := creator.Create()
		if err != nil {
			t.Fatalf("unexpected error on iteration %d: %v", i, err)
		}
		if seen[id] {
			t.Errorf("duplicate ID generated: %s", id)
		}
		seen[id] = true
	}
}
