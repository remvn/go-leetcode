package main

import (
	"fmt"
	"strings"
)

func main() {
	tests := []string{
		"/home/",
		"/../",
		"/home//foo/",
		"/home/user/../../usr/local/bin",
		"/a/./b/../../c/",
	}

	for _, path := range tests {
		fmt.Println(path)
		fmt.Println(simplifyPath(path))
		fmt.Println()
	}
}

func simplifyPath(path string) string {
	arr := strings.FieldsFunc(path, func(r rune) bool {
		return r == '/'
	})
	// fmt.Println(arr)

	resultArr := []string{}
	for i := 0; i < len(arr); i++ {
		if arr[i] == "." {
			continue
		}
		if arr[i] == ".." {
			if len(resultArr) > 0 {
				resultArr = resultArr[:len(resultArr)-1]
			}
		} else {
			resultArr = append(resultArr, arr[i])
		}
	}

	return "/" + strings.Join(resultArr, "/")
}
