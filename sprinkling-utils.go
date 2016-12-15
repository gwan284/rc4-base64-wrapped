package rc4_base64_wrapped

import (
	"math/rand"
	"time"
	"bytes"
	"sort"
	"strings"
)

func contains(idx []int, v int) bool {
	for _, elem := range idx {
		if elem == v {
			return true
		}
	}

	return false
}

func sprinklingIndices(l int) []int {
	rand.Seed(time.Now().UnixNano())

	n :=  rand.Intn(l / 3) + 1;
	idx := make([]int, 1, n)
	for i := 0; i < n; i++ {
		v := rand.Intn(l - 2) + 1
		if !contains(idx, v) {
			idx = append(idx, v)
		}
	}

	sort.Ints(idx)
	return idx
}

func sprinkle(s, ch string) string {
	idx := sprinklingIndices(len(s))

	var buffer bytes.Buffer

	buffer.WriteString(s[ : idx[0]])
	for i := 1; i < len(idx); i++ {
		l := idx[i-1]
		r := idx[i]
		buffer.WriteString(s[l : r])
		buffer.WriteString(ch)
	}

	sprinkled := buffer.String()
	left := s[idx[len(idx) -1] : ]

	return sprinkled + left
}

func unsprinkle(s, ch string) string {
	return strings.Replace(s, ch, "", -1)
}
