package hashmap

import (
	"testing"
	"fmt"
)

var testStrings = []string{"123","321","dev","test","golang","hash","list","map","::::"}

func TestNewHashMap(t *testing.T) {
	hMap := NewHashMap(3)
	for i,e := range testStrings{
		fmt.Println(i,e)
		hMap.Insert(e,i*10)
	}
	fmt.Println(hMap)
}
