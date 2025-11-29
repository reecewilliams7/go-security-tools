package clientcredentials

import "testing"

func TestNewShortUUIDClientIDCreator(t *testing.T) {
	creator := NewShortUUIDClientIDCreator()
	if creator == nil {
		t.Fatal("expected non-nil ShortUUIDClientIDCreator")
	}
}

func TestShortUUIDClientIDCreator_Create(t *testing.T) {
	creator := NewShortUUIDClientIDCreator()

	id, err := creator.Create()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if id == "" {
		t.Error("expected non-empty client ID")
	}
	// ShortUUID is typically 22 characters
	if len(id) < 20 || len(id) > 30 {
		t.Errorf("unexpected ShortUUID length: %d", len(id))
	}
}

func TestShortUUIDClientIDCreator_Create_Uniqueness(t *testing.T) {
	creator := NewShortUUIDClientIDCreator()
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
