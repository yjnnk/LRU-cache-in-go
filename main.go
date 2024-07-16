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

func (lru LRUCache) moveToHead(node *Node) {
	
}

func (lrucCache LRUCache) Get(key string) interface{} {

	if node, exists := lrucCache.cache[key]; exists {
		// deve se tornar a nova cabeça
		lrucCache.moveToHead(node)
		return node.value
	}
	return -1
}

func (lrucCache LRUCache) Set(key string) interface{} {
	return lrucCache.capacity + 1
}

func main() {
	// oi := NewLRUCache(5)

	// fmt.Println(oi.Get())

}
