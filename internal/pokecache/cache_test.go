package pokecache

import (
	"testing"
	"fmt"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "hello",
			val: []byte("world"),
		},
		{
			key: "foo",
			val: []byte("bar"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			actual, exists := cache.Get(c.key)
			if !exists {
				t.Errorf("expected key %q to exist", c.key)
				t.Fail()
			}
			if string(actual) != string(c.val) {
				t.Errorf("case %d: expected value %q, got %q", i, c.val, actual)
				t.Fail()
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5 * time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("hello", []byte("world"))

	_, ok := cache.Get("hello")
	if !ok {
		t.Errorf("expected key to exist")
		t.Fail()
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("hello")
	if ok {
		t.Errorf("expected key to be reaped")
		t.Fail()
	}
}