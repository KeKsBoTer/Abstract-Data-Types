package main

import (
	"trees/twoThreeTree"
	"fmt"
)

func main() {
	/*m := map[string]interface{}{
		"Simon": 1,
		"Daniel": "test",
		"Lücie": 564,
		"Gustav": 1339,
		"Heinz": 123,
		"Dieter": 34,
		"Caro": 21,
		"Patrick": 123,
	}*/
	tree := twoThreeTree.NewTree()
	for  i:=0;i<10;i++{
		tree.Insert(i)
	}
	fmt.Println(tree)
}
