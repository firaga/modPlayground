package main

func f() int {
	i := 5
	defer func() {
		i++
	}()
	return i
}

func f1() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func f4() (r int) {
	defer func() {
		r = r + 5
	}()
	//r=1;
	//r+=5;
	//ret
	return 1
}

func main() {
	println(f())
	println(f1())
	println(f2())
	println(f3())
	//改f3的变量拷贝为变量引用,
	println(f4())
}

//作者：尼不要逗了
//链接：https://zhuanlan.zhihu.com/p/63354092
//来源：知乎
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
