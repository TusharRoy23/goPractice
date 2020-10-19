package main

import "fmt"

func main() {
	arr := []interface{}{1, []interface{}{2, 3, "null", 4}, []interface{}{"null"}, 5}
	result := dfs(arr)
	fmt.Println(result)
}

func dfs(arr interface{}) []interface{} {
	number := []interface{}{}
	arr2 := arr.([]interface{})

	for _, val := range arr2 {
		if val == nil || val == "null" {
			continue
		}
		if val != nil {
			switch v := val.(type) {
			case int:
				number = append(number, v)
			default:
				number = append(number, dfs(val.([]interface{}))...)
			}
		}
	}

	return number
}
