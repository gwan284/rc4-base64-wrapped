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

func sprinkleIndices(s string) []int {
	rand.Seed(time.Now().UnixNano())

	n :=  rand.Intn(len(s) / 2) + 1;
	idx := make([]int, 1, n)
	for i := 0; i < n; i++ {
		v := rand.Intn(n - 2) + 1
		if !contains(idx, v) {
			idx = append(idx, v)
		}
	}
	return idx
}

func sprinkle(s, ch string) string {
	pos := sprinkleIndices(s)
	sort.Ints(pos)

	var buffer bytes.Buffer

	npos := 0
	for i := range pos {
		buffer.WriteString(s[npos - i : pos[i] + i])
		buffer.WriteString(ch)
		npos = buffer.Len()
	}
	sprinkled := buffer.String()
	left := len(s) - (len(sprinkled) - len(pos))

	return sprinkled + s[len(s) - left:]
}

func unsprinkle(s, ch string) string {
	return strings.Replace(s, ch, "", -1)
}
