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

func (lru *LRUCache) removeNode(node *Node) {
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
}

func (lru *LRUCache) addNode(node *Node) {
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

func (lru *LRUCache) moveToHead(node *Node) {

	//remover nó: tirá-lo da lista
	lru.removeNode(node)

	//adicionar nó: adicioná-lo à lista
	lru.addNode(node)
}

func (lrucCache *LRUCache) Get(key string) interface{} {

	if node, exists := lrucCache.cache[key]; exists {
		lrucCache.moveToHead(node)
		return node.value
	}
	return -1
}

func (c *LRUCache) removeTail() *Node {
	if c.tail == nil {
		return nil
	}
	node := c.tail
	c.removeNode(node)
	return node
}

func (lru *LRUCache) Set(key string, value interface{}) interface{} {

	if node, found := lru.cache[key]; found {
		node.value = value
		lru.moveToHead(node)
		return nil
	}

	newNode := &Node{
		key:   key,
		value: value,
	}

	lru.cache[key] = newNode

	newNode.prev = nil
	newNode.next = lru.head

	if lru.head != nil {
		lru.head.prev = newNode
	}
	lru.head = newNode

	if lru.tail == nil {
		lru.tail = newNode
	}

	if len(lru.cache) > lru.capacity {
		tail := lru.removeTail()
		delete(lru.cache, tail.key)
	}

	return nil
}

func main() {
	// cac := NewLRUCache(5)
}
