package main

type LRUCache struct {
	capacity int
	cache    map[string]*Node
	head     *Node
	tail     *Node
}

type Node struct {
	key   string
	value interface{}
	next  *Node
	prev  *Node
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[string]*Node),
	}
}

func (lru *LRUCache) moveToHead(node *Node) {

	//remover nó: tirá-lo da lista
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		lru.head = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		lru.tail = node.prev
	}

	//adicionar nó: adicioná-lo à lista
	node.prev = nil
	node.next = lru.head

	if lru.head != nil {
		lru.head.prev = node
	}
	lru.head = node

	if lru.tail == nil {
		lru.tail = node
	}
}

func (lrucCache *LRUCache) Get(key string) interface{} {

	if node, exists := lrucCache.cache[key]; exists {
		lrucCache.moveToHead(node)
		return node.value
	}
	return -1
}

func (lru *LRUCache) Set(key string, value interface{}) interface{} {

	return lru.capacity + 1
}

func main() {
	// oi := NewLRUCache(5)

	// fmt.Println(oi.Get())

}
