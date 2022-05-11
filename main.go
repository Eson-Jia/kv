package main

import (
	"fmt"
	"time"
)

//需要设计一个kv缓存，缓存中的数据会过期。现在要求完成缓存的增删改查功能
// 不需要支持并发访问

type Value struct {
	value  int
	expire time.Time //s
}

type KV struct {
	store map[string]Value
	ctx   chan struct{}
}

func NewKV(expireFrame time.Duration) *KV {
	store := make(map[string]Value)
	ctx := make(chan struct{})
	go func() {
		deleteList := make([]string, 0)
		tick := time.Tick(expireFrame)
		defer fmt.Printf("clear")
		for {
			select {
			case <-ctx:
				fmt.Println("KV clear")
				return
			case <-tick:
				now := time.Now()
				for k, v := range store {
					if now.After(v.expire) {
						deleteList = append(deleteList, k)
					}
				}
				for _, k := range deleteList {
					delete(store, k)
					fmt.Printf("delete expired key %s\n", k)
				}
				deleteList = nil
			}

		}
	}()
	kv := KV{
		store: store,
		ctx:   ctx,
	}
	return &kv
}

func (kv *KV) Stop() {
	close(kv.ctx)
}
func (kv *KV) Insert(k string, v int, expire time.Duration) {
	kv.store[k] = Value{
		value:  v,
		expire: time.Now().Add(expire),
	}
}
func (kv *KV) Update(k string, v int, expire time.Duration) bool {
	if _, ok := kv.store[k]; ok {
		kv.Insert(k, v, expire)
		return true
	}
	return false
}
func (kv *KV) Delete(k string) bool {
	if _, ok := kv.store[k]; ok {
		delete(kv.store, k)
		return true
	}
	return false
}
func (kv *KV) Get(k string) (int, bool) {
	if v, ok := kv.store[k]; ok {
		return v.value, true
	}
	return 0, false
}

func main() {

}
