package main

import (
	"container/list"
	"fmt"
)

/**
运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。

获取数据 get(key) - 如果关键字 (key) 存在于缓存中，则获取关键字的值（总是正数），否则返回 -1。
写入数据 put(key, value) - 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字/值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。

进阶:你是否可以在 O(1) 时间复杂度内完成这两种操作？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/lru-cache
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type LRUCache struct {
	capacity int
	lru      *list.List
	items    map[int]*list.Element
}

type entry struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		lru:      list.New(),
		items:    map[int]*list.Element{},
	}
}

func (this *LRUCache) Get(key int) int {
	if element, found := this.items[key]; found {
		this.lru.MoveToFront(element)
		return element.Value.(entry).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	putEntry := entry{
		key:   key,
		value: value,
	}
	if element, found := this.items[key]; found {
		element.Value = putEntry
		this.lru.MoveToFront(element)
		return
	}
	this.items[key] = this.lru.PushFront(putEntry)

	if this.lru.Len() > this.capacity {
		deleteEntry := this.lru.Remove(this.lru.Back()).(entry)
		delete(this.items, deleteEntry.key)
	}
	return
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	lruCache := Constructor(3)
	lruCache.Put(1, 1)
	lruCache.Put(2, 2)
	lruCache.Put(3, 3)
	lruCache.Put(4, 4)
	fmt.Println(lruCache.Get(1))
}
