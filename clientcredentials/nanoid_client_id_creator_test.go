package clientcredentials

import (
	"strings"
	"testing"
)

func TestNewNanoidClientIDCreator(t *testing.T) {
	creator := NewNanoidClientIDCreator()
	if creator == nil {
		t.Fatal("expected non-nil NanoidClientIDCreator")
	}
}

func TestNanoidClientIDCreator_Create(t *testing.T) {
	creator := NewNanoidClientIDCreator()

	id, err := creator.Create()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if id == "" {
		t.Error("expected non-empty ID")
	}
}

func TestNanoidClientIDCreator_Create_Length(t *testing.T) {
	creator := NewNanoidClientIDCreator()

	id, err := creator.Create()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(id) != nanoidSize {
		t.Errorf("expected length %d, got %d", nanoidSize, len(id))
	}
}

func TestNanoidClientIDCreator_Create_URLSafeAlphabet(t *testing.T) {
	creator := NewNanoidClientIDCreator()

	for i := range 100 {
		id, err := creator.Create()
		if err != nil {
			t.Fatalf("unexpected error on iteration %d: %v", i, err)
		}
		for _, ch := range id {
			if !strings.ContainsRune(nanoidAlphabet, ch) {
				t.Errorf("ID contains character outside nanoid alphabet: %q (full ID: %s)", ch, id)
			}
		}
	}
}

func TestNanoidClientIDCreator_Create_Uniqueness(t *testing.T) {
	creator := NewNanoidClientIDCreator()
	seen := make(map[string]bool)

	for i := range 100 {
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
