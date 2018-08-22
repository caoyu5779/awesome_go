package main

import "testing"

func TestSubstr (t *testing.T){
	test := [] struct{
		s string
		ans int
	}{
		{"abcabcbb",3},

		{"", 0},
		{"b", 1},
		{"abcabcabcd", 4},

		{"我爱拜仁", 4},
	}

	for _,tt := range test{
		actual := lengthOfNonRepeatingSubStr(tt.s)

		if actual != tt.ans{
			t.Errorf("got %d for input %s ;" + "excepted %d", actual, tt.s, tt.ans)
		}
	}
}

func BenchmarkSubstr(b *testing.B){
	s := "我爱拜仁"
	for i := 0; i <13 ; i++{
		s += s
	}
	b.Logf("len(s) = %d", len(s))
	ans := 4
	b.ResetTimer()

	for i := 0; i < b.N; i++{
		actual := lengthOfNonRepeatingSubStr(s)

		if actual != ans{
			b.Errorf("got %d for input %s ;" + "excepted %d", actual, s, ans)
		}
	}

}
