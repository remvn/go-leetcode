package main

import (
	"fmt"
	"strings"
)

func main() {
	tests := []string{
		"babad",
		"eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee",
		"aacabdkacaa",
		"busislnescsicxpvvysuqgcudefrfjbwwjcchtgqyajdfwvkypfwshnihjdztgmyuuljxgvhdiwphrweyfkbnjgerkmifbirubhseuhrugwrabnjafnbdfjnufdstjbkuwtnpflffaqmjbhssjlnqftgjiglvvequhapasarlkcvbmkwnkuvwktbgfoaxteprobdwswcdyddyvrehvmxrrjiiidatidlpihkbmmruysmhhsncmfdanafdrfpdtfgkglcqpwrrtvacuicohspkounojuziittugpqjyhhkwfnflozbispehrtrnizowrlzcuollagxwtznjwzcumvedjwokueuqktvvouwnsmpxqvvpuwprezrbobrpnwaccwljchdguubjulyilzvmandjjleitweybqkjttschrjjlebnmponvlktzzcdtuybugggcqffkcffpamauvxfbonjrobgpvlyzveiwemmtdvbjciaytvesnocnjrwodtcokgcuoiicxapmrzpkfphjniuvzjrhbnqndfshoduejyktebgdabidxlkstepuwvtrtgbxaeheylicvhrxddijshcvdadxzsccmainyfpfdhqdanfccqkzlmhsfilvoybqojlvbcixjzqpbngdvesuokbxhkomsiqfyukvspqthlzxdnlwthrgaxhtpjzhrugqbfokrdcyurivmzgtynoqfjbafboselxnfupnpqlryvlcxeksirvufepfwczosrrjpudbwqxwldgjyfjhzlzcojxyqjyxxiqvfhjdwtgoqbyeocffnyxhyyiqspnvrpxmrtcnviukrjvpavervvztoxajriuvxqveqsrttjqepvvahywuzwtmgyrzduxfqspeipimyoxmkadrvrdyefekjxcmsmzmtbugyckcbjsrymszftjyllfmoeoylzeahnrxlxpnlvlvzltwnmldi",
	}
	tests = append(tests, strings.Repeat("a", 10000))
	for _, test := range tests {
		fmt.Printf("case %s:\n", test)
		fmt.Println(longestPalindrome(test))
		fmt.Println()
	}
}

func longestPalindrome(s string) string {
	if s == "" || len(s) == 1 {
		return s
	}
	if len(s) == 2 {
		if s[0] == s[1] {
			return s
		}
		return string(s[0])
	}

	longest := string(s[0])
	repeatMap := make(map[byte][]int)

	for i := 0; i < len(s); i++ {
		repeatMap[s[i]] = append(repeatMap[s[i]], i)
	}

	for _, item := range repeatMap {
		if len(item) < 2 {
			continue
		}
		for i := 0; i < len(item); i++ {
			for j := i + 1; j < len(item); j++ {
				checkStr := s[item[i] : item[j]+1]
				if len(checkStr) < len(longest) {
					continue
				}
				// fmt.Println(checkStr)
				if !checkPalindrome(checkStr) {
					continue
				}
				if len(longest) < len(checkStr) {
					longest = checkStr
				}
			}
		}
	}

	return longest
}

func checkPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// Other people's solution
// func longestPalindrome(s string) string {
//
//     if (s == "" || len(s) < 1){return ""}
//
//     start := 0
//     end := 0
//
//     for i := 0; i < len(s); i++{
//         // The reason we have two expand from center:
//         // eg: aabb & aabaa
//         // both of them are valid palindromes
//         // we need a second expand to skip
//		   // the center 'b' character
//         len1 := expandFromCenter(s,i,i)
//         len2 := expandFromCenter(s,i,i+1)
//         len := 0
//
//         if len1 > len2{len=len1}else{len=len2}
//
//         if len > end - start {
//             start = i - ((len - 1) / 2)
//             end = i + (len / 2)
//         }
//
//     }
//
//     return s[start: end + 1]
//
// }
//
// func expandFromCenter(s string, left,right int) int{
//
//     if (s == "" || left > right){return 0}
//
//     for (left >= 0 && right < len(s) && s[left] == s[right]){
//         left--
//         right++
//     }
//
//     return right - left - 1
// }
