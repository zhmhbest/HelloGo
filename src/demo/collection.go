package demo

import "fmt"

func CollectionMap() {
	var kv = make(map[string]string)
	kv["key1"] = "value1"
	kv["key2"] = "value2"
	kv["key3"] = "value3"
	delete(kv, "key2")
	for i, v := range kv {
		fmt.Println(i, "=", v)
	}
	fmt.Print("\n")
}
