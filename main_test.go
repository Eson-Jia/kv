package main

import (
	"fmt"
	"testing"
	"time"
)

func TestKV_Get(t *testing.T) {
	kv := NewKV(time.Second)
	_, ok := kv.Get("not exist")
	if ok {
		t.Errorf("not exist")
	}
}

func TestKV(t *testing.T) {
	kv := NewKV(time.Millisecond)
	kv.Insert("one", 1, time.Second)
	kv.Insert("two", 2, time.Second)
	kv.Insert("three", 3, time.Millisecond*100)
	kv.Delete("two")
	time.Sleep(time.Second * 3)
	if ok := kv.Update("two", 22, time.Second); ok {
		fmt.Println("update success")
	} else {
		fmt.Println("update failed")
	}
	if v, ok := kv.Get("one"); ok {
		fmt.Println("one", v)
	}
	time.Sleep(time.Second * 2)
	kv.Stop()
	time.Sleep(time.Second * 2)
}
