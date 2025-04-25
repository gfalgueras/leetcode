package longestCommonPrefix14

import "fmt"

func Run() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}) == "fl")
}

func longestCommonPrefix(strs []string) string {
	if strs[0] == "" {
		return ""
	}

	res := ""

	for idx := range strs[0] {

		for i := 1; i < len(strs); i++ {
			if len(strs[i]) == idx || strs[i][idx] != strs[0][idx] {
				return res
			}
		}

		res = res + string(strs[0][idx])
	}

	return res
}
