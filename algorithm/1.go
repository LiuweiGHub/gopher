package main

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"
)

type ListNode struct {
	Next *ListNode
	Val  int
}

/**
var s []int{1,2,3,4,5} 往slice添加元素后，容量变成多少  // 10  扩容原理   先是翻倍 没毛病
x = s[3:5] 此时x的len以及cap分别是多少  // 2   5(2) cap为啥是2   通过切片s初始化x1
x = append(x, 6)此时输出s是多少  []int{1,2,3,4,5}
x = append(x, 7,8,9) 此时输出s是多少  []int{1,2,3,4,5}
*/

var ErrClosedPipe = errors.New("xxxxxxxxxxxx")

const n int = 6
const m = 8

const (
	x    = 9
	t    = false
	e, r = 1 << iota, 10 + iota
	_, _
	g, h
	i, j
)

var (
	a int
	b = 2
	c = int64(9)
	d []byte
	f = 3.14
)

func demo1() {

	var m = [...]int{1, 3, 4, 5, 6, 7}

	for i, v := range m {
		go func(i, v int) {

			time.Sleep(3 * time.Second)
			fmt.Println(i, v)
		}(i, v)
	}

	time.Sleep(time.Second * 10)
}

type Config struct {
	path string
}

func (c *Config) Path() string {
	if c == nil {
		return "/home/user"
	}
	return c.path
}

func loop() {
start:
	for {
		fmt.Println("Start")
		s := []int{1, 2, 3}
		for _, v := range s {
			fmt.Println(v)
			time.Sleep(time.Second * 1)
			// break/continue + label 方式 语义更丰富
			continue start
		}
	}
}

func deadloop() {
	for {
	}
}

func spawnGroup(n int, f func(args ...interface{}), args ...interface{}) chan struct{} {
	c := make(chan struct{})
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			name := fmt.Sprintf("work-%d", i)
			f(args...)
			println(name, "done")
			wg.Done()
		}(i)
	}

	go func() {
		wg.Wait()
		c <- struct{}{}
	}()

	return c
}

func worker(args ...interface{}) {
	if len(args) == 0 {
		return
	}

	interval, ok := args[0].(int)
	if !ok {
		return
	}

	time.Sleep(time.Second * (time.Duration(interval)))
}

func main() {

	done := spawnGroup(5, worker, 3)
	println("spawn a group of wokers")
	<-done
	println("group workers done")

	return

	go deadloop()
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("I'm get scheduled.")
	}

	var c1 *Config
	var c2 = &Config{path: "/export"}
	fmt.Println(c1.Path(), c2.Path())
	return
	demo1()
	arr1 := [...]int{1, 2, 3}
	sli := []int{1, 2, 3}
	fmt.Printf("%T %T\n", arr1, sli)

	switch x, y := 1, 2; x + y {

	case 3, 5, 6:
		fmt.Println(x + y)
		fallthrough
	case 4:
		fmt.Println(x + y)
		// 跳到下个case 不用命中case也可执行
		fallthrough
	default:
		fmt.Println("default")

	}
	return

	{
		{
			a := 1
			if true {
				println(a)
			}
		}
	}

	n0, n1 := 1, 2
	n0, n1 = n0+n1, n0

	fmt.Println(n0, n1)
	return

	arr := [8]int{1, 2, 3, 4, 5, 6, 7, 8}

	s11 := arr[1:6] // 数组获取切片 左右都闭

	s22 := s11[2:3] // 切片获取切片 左开右闭
	fmt.Print(s11, s22)

	var m = map[string]int{
		"a": 1,
		"b": 2,
	}

	m1 := make(map[string]string, 3)
	m1["a"] = "a"
	m1["b"] = "c"

	fmt.Print(m, m1)

	str := `好雨知时节，
当春乃发生.
随风潜入夜，
润物细无声.`

	fmt.Printf("%T\n", m1)
	fmt.Println(str)

	return

	r := longestPalindrome("aaasdssdsa")
	fmt.Println(r)
	return
	for i, v := range []int{1, 2, 3} {
		fmt.Println(i, v)
	}

	return

	for _, v := range strconv.Itoa(-123) {
		fmt.Println(string(v))
	}
	res := longestPalindrome("abba")
	fmt.Println(res)

	return
	var s = []int{1, 2, 3, 4, 5}
	//s = append(s, 6)
	x1 := s[3:5]
	x1 = append(x1, 6)
	x1 = append(x1, 7, 8, 9)
	fmt.Println(len(x1), cap(x1))
	fmt.Println(s)

	var sm sync.Map
	//sync.Map set
	sm.Store("greece", 97)
	sm.Store("london", "100")
	sm.Store("eqypt", 123)

	//sync.Map get
	fmt.Println(sm.Load("london"))

	//sync.Map delete
	sm.Delete("london")

	//sync.Map 遍历
	sm.Range(func(k, v interface{}) bool {
		fmt.Println("interate:", k, v)
		return true
	})

	return

	x := -2147483412
	r = reverse(x)
	fmt.Println(r)
	return

	//res = findMedianSortedArrays([]int{2,1}, []int{3,4})
	//fmt.Println(res)
}

func longestPalindrome(s string) int {

	m := make(map[rune]int, len(s))
	for _, v := range s {
		m[v] += 1
	}

	maxLen := 0
	flag := true
	for _, v := range m {
		if v%2 == 0 {
			maxLen += v
		}

		if v%2 != 0 {
			if v == 1 && flag {
				maxLen += 1
				flag = false
			} else {
				if flag {
					maxLen += v
					flag = false
				} else {
					maxLen += v - 1
				}
			}
		}

	}
	return maxLen

}

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	bs := []byte(s)

	var nbs string
	for _, v := range bs {

		if unicode.IsLetter([]rune(string(v))[0]) {
			nbs += strings.ToLower(string(v))
		}
		if unicode.IsNumber([]rune(string(v))[0]) {
			nbs += string(v)
		}
	}

	i := 0
	j := len(s) - 1
	for i < j {
		if bs[i] != bs[j] {
			return false
		}
		i++
		j--
	}
	return true
}

//func longestPalindrome(s string) string {
//
//      bs := []byte(s)
//      if len(bs) < 2 {
//              return s
//      }
//      maxLen := 1
//      begin := 0
//      for i := 0; i < len(bs) - 1; i++ {
//              for j := i + 1; j < len(bs); j++ {
//                      if j-i+1 > maxLen && validPalindrome(bs, i, j) {
//                              maxLen = j - i + 1
//                              begin = i
//                      }
//              }
//      }
//      return string(bs[begin : begin+maxLen])
//}

func validPalindrome(bs []byte, i int, j int) bool {
	for i < j {
		if bs[i] != bs[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {

	var tail *ListNode
	carry := 0

	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		//fmt.Println(sum)
		sum, carry = sum%10, sum/10 //因为sum是一个[0-18]的整数  所以除以10只能要么是0 要么是1
		//fmt.Println(carry)
		//fmt.Println(sum, carry)

		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}

	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return
}

func reverse(x int) int {
	if x >= math.MaxInt32 || x <= math.MinInt32 {
		return 0
	}
	fmt.Println(x)
	z := []string{}
	newS := ""
	// 栈 + 标记正负号的falg
	flag := false
	//转成字符串遍历
	s := strconv.Itoa(x)

	for _, v := range s {
		if string(v) == "-" {
			flag = true
			continue
		}
		z = append(z, string(v))
	}

	l := len(z)
	// 又忘了
	if flag {
		newS += "-"
	}
	for i := l - 1; i >= 0; i-- {
		newS += z[i]
		fmt.Println(i, z, newS)
	}
	//字符串转成int返回

	n, _ := strconv.Atoi(newS)
	return n
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	for _, v := range "bbbb" {
		fmt.Println(string(v)) // string（）强制转换
	}
	var res float64
	if len(nums1) == 0 && len(nums2) == 0 {
		return res
	}
	num := append(nums1, nums2...)

	fmt.Println(num)
	sort.Ints(num)
	fmt.Println(num)

	l := len(num)
	// num排序
	// 奇偶判断
	if l%2 == 0 {
		fmt.Println(num[l/2])
		fmt.Println(num[l/2-1])
		res = float64(num[l/2]+num[l/2-1]) / 2
	} else {
		res = float64(num[l/2])
	}

	return res
}

func lengthOfLongestSubstring(s string) int {

	right, left := 0, 0
	window := s[left:right]
	res := 0

	fmt.Println(len(s))
	if len(s) == 0 {
		return res
	}
	if len(s) == 1 {
		return 1
	}

	for ; right < len(s); right++ {
		window = s[left:right]
		fmt.Println(window)
		if index := strings.IndexByte(window, s[right]); index != -1 {
			left += index + 1
		}
		res = int(math.Max(float64(res), float64(len(window))))
	}
	return res
}
