package main

import (
	"trees/twoThreeTree"
	"fmt"
	"math/rand"
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
	values := make([]int,200)
	for  i:=0;i<200;i++{
		ran := rand.Intn(500)-100
		values[i]=ran
		tree.Insert(i)
	}
	fmt.Println(tree)
	fmt.Printf("%d, %d",len(values),tree.Length())
}
