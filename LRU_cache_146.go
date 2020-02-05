type LRUCache struct {
    m map[int]*list.Element
    list *list.List
    capacity int 
}

type KeyValue struct {
    key int
    value int
}

func NewLRUCache (capacity int) LRUCache {
    var cache LRUCache
    cache.capacity = capacity
    cache.m = make(map[int]*list.Element)
    cache.list = list.New()
    return cache
}

func (lru *LRUCache) Get(key int) int {
    if e, ok := lru.m[key]; ok {
        lru.list.MoveToFront(e)
        return e.Value.(*KeyValue).value
    } 
    
    return -1
}

func (lru *LRUCache) Put(key, value int) {
    if e, ok := lru.m[key]; ok {
        lru.list.MoveToFront(e)
        e.Value = &KeyValue{key, value}
    } else {
        if lru.list.Len() == lru.capacity {
            e := lru.list.Back()
            delete(lru.m, e.Value.(*KeyValue).key)
            lru.list.Remove(e)
        }

        lru.list.PushFront(&KeyValue{key, value})
        lru.m[key] = lru.list.Front()
    }
}
