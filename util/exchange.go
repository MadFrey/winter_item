package util

func ArrayToString(arr []string) string {
	var result string
	for count, i := range arr {  //遍历数组中所有元素追加成string
		result += i
		if count==len(arr)-1 {
			break
		}
		result+=","
	}
	return result
}

