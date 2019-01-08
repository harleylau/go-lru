package lru

import (
	"container/list"
	"reflect"
	"testing"
)

func TestLRU_Clear(t *testing.T) {
	type fields struct {
		size      int
		cacheList *list.List
		items     map[interface{}]*list.Element
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &LRU{
				size:      tt.fields.size,
				cacheList: tt.fields.cacheList,
				items:     tt.fields.items,
			}
			c.Clear()
		})
	}
}

func TestLRU_PUT(t *testing.T) {
	type fields struct {
		size      int
		cacheList *list.List
		items     map[interface{}]*list.Element
	}
	type args struct {
		key   interface{}
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &LRU{
				size:      tt.fields.size,
				cacheList: tt.fields.cacheList,
				items:     tt.fields.items,
			}
			c.PUT(tt.args.key, tt.args.value)
		})
	}
}

func TestLRU_Get(t *testing.T) {
	type fields struct {
		size      int
		cacheList *list.List
		items     map[interface{}]*list.Element
	}
	type args struct {
		key interface{}
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantValue interface{}
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &LRU{
				size:      tt.fields.size,
				cacheList: tt.fields.cacheList,
				items:     tt.fields.items,
			}
			gotValue, err := c.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("LRU.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("LRU.Get() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestLRU_DEL(t *testing.T) {
	type fields struct {
		size      int
		cacheList *list.List
		items     map[interface{}]*list.Element
	}
	type args struct {
		key interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &LRU{
				size:      tt.fields.size,
				cacheList: tt.fields.cacheList,
				items:     tt.fields.items,
			}
			if got := c.DEL(tt.args.key); got != tt.want {
				t.Errorf("LRU.DEL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLRU_Len(t *testing.T) {
	type fields struct {
		size      int
		cacheList *list.List
		items     map[interface{}]*list.Element
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &LRU{
				size:      tt.fields.size,
				cacheList: tt.fields.cacheList,
				items:     tt.fields.items,
			}
			if got := c.Len(); got != tt.want {
				t.Errorf("LRU.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}
