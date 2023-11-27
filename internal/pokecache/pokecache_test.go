package pokeacache

import (
	"testing"
	"time"
)

// test to see we're never return nil
func TestCreateCache(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)
	// if there is no key value pairs
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

// test adding of key value pairs
func TestAddGetCache(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{inputKey: "key1", inputVal: []byte("val1")},
		{inputKey: "key2", inputVal: []byte("val2")},
		{inputKey: "key3", inputVal: []byte("val3")},
	}

	for _, cas := range cases {
		cache.Add(cas.inputKey, cas.inputVal)
		actual, ok := cache.Get(cas.inputKey)
		if !ok {
			t.Errorf("%s does not exist", cas.inputKey)
			continue
		}
		if string(actual) != string(cas.inputVal) {
			t.Errorf("%s value doesn't match %s", string(actual), string(cas.inputVal))
		}
	}
}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	keyOne := "key1"
	cache.Add(keyOne, []byte("hello"))

	//needs to wait longer than the interval
	time.Sleep(interval + time.Millisecond)

	_, ok := cache.Get(keyOne)
	if ok {
		t.Errorf("%s did not get reaped", keyOne)
	}
}
