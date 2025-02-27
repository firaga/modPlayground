package main

import (
	"fmt"
	"strconv"
)

func roundToTwoNonZeroDecimals(f float64) (float64, error) {
	// 将浮点数转换为字符串
	str := strconv.FormatFloat(f, 'f', -1, 64)

	// 查找小数点的位置
	dotIndex := -1
	for i, c := range str {
		if c == '.' {
			dotIndex = i
			break
		}
	}
	if dotIndex == -1 {
		return 0, fmt.Errorf("the number %f does not have a decimal point", f)
	}

	// 从第一个小数位开始查找前两个非零值
	nonZeroCount := 0
	nonZeroIndex := -1
	for i := dotIndex + 1; i < len(str); i++ {
		if str[i] != '0' && str[i] != '.' {
			nonZeroCount++
			if nonZeroCount == 2 {
				nonZeroIndex = i
				break
			}
		}
	}
	if nonZeroIndex == -1 {
		return 0, fmt.Errorf("less than two non-zero digits found in the decimal part of %f", f)
	}

	// 确定四舍五入的位置
	roundIndex := nonZeroIndex + 1
	if roundIndex >= len(str) {
		return f, nil // 已经是最精确的结果
	}

	// 获取需要四舍五入的部分
	roundDigit, _ := strconv.Atoi(string(str[roundIndex]))

	// 四舍五入
	var roundedStr string
	if roundDigit >= 5 {
		roundedStr = str[:roundIndex-1] + string(str[roundIndex-1]+1)
	} else {
		roundedStr = str[:roundIndex]
	}

	// 转换回浮点数
	roundedF, err := strconv.ParseFloat(roundedStr, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse rounded string: %v", err)
	}

	return roundedF, nil
}

func main() {
	f := 0.000000001369
	fmt.Println(f)
	roundedF, err := roundToTwoNonZeroDecimals(f)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The rounded value of %.10f with two non-zero decimals is: %.10f\n", f, roundedF)
	}
}
