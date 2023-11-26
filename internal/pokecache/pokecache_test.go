package pokeacache

import "testing"

// test to see we're never return nil
func TestCreateCache(t *testing.T) {
	cache := NewCache()
	// if there is no key value pairs
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

// test adding of key value pairs
func TestAddGetCache(t *testing.T) {
	cache := NewCache()

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
