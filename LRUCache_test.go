package main

import (
	"testing"
)

func TestNodeNotFound(t *testing.T) {
	lru := NewLRUCache(2)
	val := lru.Get("chave_aleatoria")

	if val != -1 {
		t.Errorf("expected -1, got %v", val)
	}
}

func TestNodeFound(t *testing.T) {
	lru := NewLRUCache(2)
	lru.cache["123"] = &Node{
		key:   "123",
		value: "meu_valor",
	}

	val := lru.Get("123")

	if val != "meu_valor" {
		t.Errorf("expected meu_valor, got %v", val)
	}

	if lru.head != lru.cache["123"] {
		t.Errorf("expected meu_valor, got %v", lru.head)
	}
}

func TestSetCache(t *testing.T) {
	lru := NewLRUCache(2)
	lru.Set("key1", "value1")
	lru.Set("key2", "value2")

	val := lru.Get("key1")
	val2 := lru.Get("key2")

	if val != "value1" {
		t.Error("wrong value")
	}

	if val2 != "value2" {
		t.Error("wrong value")
	}

	if lru.head.value != "value2" {
		t.Error("error wrong head")
	} 
}
