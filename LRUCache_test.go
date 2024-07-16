package main

import "testing"

func TestNodeNotFound(t *testing.T) {
	lru := NewLRUCache(2)
	val := lru.Get("chave_aleatoria")

	if(val != -1) {
		t.Errorf("expected -1, got %v", val)
	}

}