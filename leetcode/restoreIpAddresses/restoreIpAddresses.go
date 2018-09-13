package restoreIpAddresses

import (
	"fmt"
)

func RestoreIpAddresses(s string) []string {
	n := len(s)

	if n < 4 || n > 12 {
		return []string{}
	}

	res := []string{}
	combination := make([]string, 4)

	var dfs func(int, int)

	dfs = func(idx int, begin int) {
		if idx == 3 {
			//idx =3
			//s[9:]
			temp := s[begin:]
			//true
			if isOk(temp) {
				//combination[3]=s[9:]
				combination[3] = temp
				res = append(res, Ip(combination))
			}

			return
		}

		maxRemain := 3 * (3 - idx)//9 //6 //3
		//1;9;++ begin 0
		//4;10;++ begin 3
		//7;11;++ begin 6
		for end := begin +1 ; end <= n- (3-idx); end++ {
			//3+9=10<12 end = 3
			//4+6=10<12 end = 6
			//7+3=10<12 end = 9
			if end+maxRemain < n {
				continue
			}

			if end-begin > 3 {
				break
			}
			//s[0:3]
			//s[3:6]
			//s[6:9]
			temp := s[begin:end]
			//true
			//true
			//true
			if isOk(temp) {
				//combination[0]=255
				//combination[1]=255
				//combination[2]=255
				combination[idx] = temp
				//dfs(1,3)
				//dfs(2,6)
				//dfs(3,9)
				dfs(idx+1, end)
			}
		}
	}

	dfs(0,0)

	return res
}

func isOk(s string) bool{
	if len(s) > 1 && s[0] == '0' {
		return false
	}

	if len(s) < 3 {
		return true
	}

	switch s[0] {
	case '1':
		return true
	case '2':
		if '0' <= s[1] && s[1] <= '4' {
			return true
		}
		if s[1] <= '5' && '0' <= s[2] && s[2] <= '5' {
			return true
		}
	}

	return false
}

func Ip(s []string) string{
	return fmt.Sprintf("%s.%s.%s.%s", s[0], s[1], s[2], s[3])
}
