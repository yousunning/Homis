package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings" // 문자열과 특정문자 또는 부분문자열을 찾는 데 사용되는 함수들을 제공한다.
	// 문자열을 조작하고 검색, 문자열간의 비교, 분리, 결합 등을 수행하는 함수 제공.
)

func main() { // os.Stdin 은 프로그램이 사용자의 키보드 입력을 받을 수 있게 하는 표준 방법
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("간단한 계산기 입니다.")
	fmt.Println("연산을 입력하세요 (예 : 5+3):")

	for { //Scan() 표준 라이브러리에 포함된'bufio'패키지 일부 (입출력을 제공하는 타입.)
		fmt.Print(" ==> ")
		scanner.Scan()
		text := scanner.Text()

		if text == "exit" {
			break
		}

		parts := strings.Fields(text)
		if len(parts) != 3 {
			fmt.Println("올바른 형식으로 입력해주세요 (예 : 5+3).")
			continue
		}
		//사용자가 입력한 문자열에서 숫자를 추출하고, 이를 실수형으로 변환하는데 사용한다.
		//'strconv'패키지의'parsefloat'함수를 사용하여 문자열을 'float64' 타입의 실수로 변환한다.
		//parts[1]은 parts 배열의 두 번째 요소를 의미
		// op := parts[1]는 사용자가 입력한 수식에서 연산자를 추출하여 op 변수에 저장하는 역할한다.

		a, err1 := strconv.ParseFloat(parts[0], 64)
		b, err2 := strconv.ParseFloat(parts[2], 64)
		op := parts[1]
		// op새로운 변수 선언하고 parts[1]의 값을 op 에 할당

		if err1 != nil || err2 != nil {
			fmt.Println("숫지를 올바르게 입력해주세요.")
			continue // 다시 처음으로 돌아감
		}
		var result float64
		switch op {
		case "+":
			result = a + b
		case "-":
			result = a - b
		case "*":
			result = a * b
		case "/":
			if b == 0 {
				fmt.Println("0으로 나눌 수 없습니다.")
				continue
			}
			result = a / b
		default:
			fmt.Println("알 수 없는 연산자 이므로, 재확인 후 입력해주세요." +
				"사용가능한 연산자는 +,-,*,/ 입니다.")
			continue
		} //%.2f: 부동소수점 숫자를 출력하는 포맷 지정자
		fmt.Printf("결과 : %.2f", result)
	}

}

//<디버그 콘솔창>
// Type 'dlv help' for list of commands.
// 간단한 계산기 입니다.
// 연산을 입력하세요 (예 : 5+3):
//  ==> 
// 5+3
// 8
// 8*6
// 48
// 9*9
// 81
