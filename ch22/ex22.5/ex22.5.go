package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["이화랑"] = "서울시 광진구"
	m["이화랑1"] = "서울시 광진구"
	m["이화랑2"] = "서울시 광진구"
	m["이화랑3"] = "서울시 광진구"
	m["이화랑4"] = "서울시 광진구"

	for key, value := range m {
		fmt.Println(key, value)
	}

	delete(m, "이화랑2")
	v, ok := m["이화랑2"]
	if !ok {
		fmt.Println("Deleted Element")
		return
	}
	fmt.Println("Element:", v)

}
