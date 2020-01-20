type LRUCache struct {
    capacity int
    list *list.List
    m map[int]*list.Element
}

type KeyVal struct {
    key int
    val int
}

func Constructor(capacity int) LRUCache {
    var cache LRUCache
    cache.capacity = capacity
    cache.list = list.New()
    cache.m = make(map[int]*list.Element)
    return cache
}

func (this *LRUCache) Get(key int) int {
    le, ok := this.m[key]
    if !ok {
        return -1
    }
    this.list.MoveToFront(le)
    return le.Value.(*KeyVal).val
}

func (this *LRUCache) Put(key int, value int)  {
    _, ok := this.m[key]
    if ok {
	le := this.m[key]
	this.list.MoveToFront(le)
	le.Value = &KeyVal{key, value}
    } else {
	if this.list.Len() == this.capacity {
	    le := this.list.Back()
	    mkey := le.Value.(*KeyVal).key
	    delete(this.m, mkey)
	    this.list.Remove(le)
	}
	this.list.PushFront(&KeyVal{key, value})
	this.m[key] = this.list.Front()
    }
}
