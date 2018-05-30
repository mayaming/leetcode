package main

import "fmt"

type CacheEntry struct {
	Key, Value int
	Prev, Next *CacheEntry
}

type LRUCache struct {
	Capacity int
	Entries map[int]*CacheEntry
	MostRecent, LeastRecent *CacheEntry
}

func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{capacity, make(map[int]*CacheEntry), nil, nil}
	lruCache.MostRecent = &CacheEntry{-1, -1, nil, nil}
	lruCache.LeastRecent = &CacheEntry{-1, -1, nil, nil}
	lruCache.MostRecent.Next = lruCache.LeastRecent
	lruCache.LeastRecent.Prev = lruCache.MostRecent
	return lruCache
}

func (this *LRUCache) visit(entry *CacheEntry) {
	prev := entry.Prev
	next := entry.Next
	prev.Next = next
	next.Prev = prev
	entry.Next = this.MostRecent.Next
	entry.Prev = this.MostRecent
	this.MostRecent.Next.Prev = entry
	this.MostRecent.Next = entry
}

func (this *LRUCache) removeLRU() {
	target := this.LeastRecent.Prev
	if target != this.MostRecent {
		target.Prev.Next = this.LeastRecent
		this.LeastRecent.Prev = target.Prev
		delete(this.Entries, target.Key)
	}
}

func (this *LRUCache) Get(key int) int {
	entry, exists := this.Entries[key]
	if exists {
		this.visit(entry)
		return entry.Value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int)  {
	entry, exists := this.Entries[key]
	if exists {
		entry.Value = value
		this.visit(entry)
	} else {
		entry = &CacheEntry{key, value, nil, nil}
		this.Entries[key] = entry
		entry.Prev = this.MostRecent
		entry.Next = this.MostRecent.Next
		entry.Next.Prev = entry
		this.MostRecent.Next = entry
		if this.Capacity == 0 {
			this.removeLRU()
		} else {
			this.Capacity -= 1
		}
	}
}

func (this *LRUCache) curEntries() {
	for cur := this.MostRecent.Next; cur != this.LeastRecent; cur = cur.Next {
		fmt.Printf("{%d: %d}, ", cur.Key, cur.Value)
	}
	fmt.Println()
}

func main() {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.curEntries()
	obj.Put(2, 2)
	obj.curEntries()
	fmt.Println(obj.Get(1)) // returns 1
	obj.curEntries()
	obj.Put(3, 3)
	obj.curEntries()
	fmt.Println(obj.Get(2)) // returns -1 (not found)
	obj.curEntries()
	obj.Put(4, 4)
	obj.curEntries()
	fmt.Println(obj.Get(1)) // returns -1 (not found)
	obj.curEntries()
	fmt.Println(obj.Get(3)) // returns 3
	obj.curEntries()
	fmt.Println(obj.Get(4)) // returns 4
	obj.curEntries()
}
/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */