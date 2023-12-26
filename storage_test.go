package storage

import (
	"fmt"
	"testing"
	"time"
)

func TestKeyValueStore_SetGetDelete(t *testing.T) {
	s := New()

	// SET/GET
	s.Set("testKey", "testValue", 1*time.Second)
	value, err := s.Get("testKey")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if value != "testValue" {
		t.Errorf("Expected value 'testValue', got '%s'", value)
	}

	// GET after TTL expired
	time.Sleep(2 * time.Second)
	_, err = s.Get("testKey")
	if err == nil {
		t.Errorf("Expected error due to TTL expiration, got nil")
	}

	// DELETE
	s.Set("deleteKey", "deleteValue", 1*time.Minute)
	s.Delete("deleteKey")
	_, err = s.Get("deleteKey")
	if err == nil {
		t.Errorf("Expected error for deleted key, got nil")
	}
}
func TestKeyValueStore_Performance(t *testing.T) {
	s := New()

	for i := 0; i < 1000000; i++ {
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		s.Set(key, value, 10*time.Minute)
	}

	for i := 0; i < 1000000; i++ {
		key := fmt.Sprintf("key%d", i)
		s.Get(key)
	}
}
