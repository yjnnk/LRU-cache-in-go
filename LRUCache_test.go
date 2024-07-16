package main

import "testing"

func TestNodeNotFound(t *testing.T) {
	lru := NewLRUCache(2)
	val := lru.Get("chave_aleatoria")

	if(val != -1) {
		t.Errorf("expected -1, got %v", val)
	}
}

func TestNodeFound(t *testing.T) {
	lru := NewLRUCache(2)
	lru.cache["123"] = &Node{
		key: "123",
		value: "meu_valor",
	}

	val := lru.Get("123")

	if(val != "meu_valor") {
		t.Errorf("expected meu_valor, got %v", val)
	}
}