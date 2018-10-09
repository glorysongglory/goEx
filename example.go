package main

import (
	"errors"
	"fmt"
	"math"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	fmt.Println("----------------------------")

	var a1 string = "initial"
	fmt.Println(a1)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "short"
	fmt.Println(f)

	fmt.Println("----------------------------")

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}
	fmt.Println("----------------------------")

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	fmt.Println("----------------------------")
	i = 2
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's not weekend")
	default:
		fmt.Println("it's weekend")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")
	}

	fmt.Println("----------------------------")

	var arr [5]int
	fmt.Println("set:", arr)
	fmt.Println("get:", arr[4])
	fmt.Println("len:", len(arr))

	barr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dc1:", barr)

	var twoD [2][3]int
	for tmpi := 0; tmpi < 2; tmpi++ {
		for tmpj := 0; tmpj < 3; tmpj++ {
			twoD[tmpi][tmpj] = tmpi + tmpj;
		}
	}
	fmt.Println("2d:", twoD)

	fmt.Println("----------------------------")

	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0], s[1], s[2] = "a", "b", "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd", s)

	carr := make([]string, len(s))
	copy(carr, s)
	fmt.Println("cpy:", carr)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	tarr := []string{"g", "h", "i"}
	fmt.Println("dcl:", tarr)

	twoDArr := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoDArr[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoDArr[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoDArr)

	fmt.Println("----------------------------")

	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	//当从一个 map 中取值时，可选的第二返回值指示这个键是在这个 map 中。这可以用来消除键不存在和键有零值，像 0 或者 "" 而产生的歧义
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	fmt.Println("----------------------------")

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

	fmt.Println("----------------------------")
	fmt.Println("1+2=", plus(1, 2))

	fmt.Println("----------------------------")
	aa, bb := vals()
	fmt.Println(aa)
	fmt.Println(bb)

	_, cc := vals()
	fmt.Println(cc)

	fmt.Println("----------------------------")

	sumf(1, 2)
	sumf(1, 2, 3)
	num := []int{1, 2, 3, 4}
	sumf(num...)

	fmt.Println("----------------------------")

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())

	fmt.Println("----------------------------")
	fmt.Println(fact(7))

	fmt.Println("----------------------------")
	i = 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr", i)

	fmt.Println("pointer:", &i)
	fmt.Println("----------------------------")
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40})

	sp := person{name: "Sean", age: 50}
	fmt.Println(sp.name)

	sp1 := &sp
	fmt.Println(sp1.age)

	sp1.age = 51
	fmt.Println(sp.age)
	fmt.Println("----------------------------")
	r := rect{width: 10, height: 5}
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())

	rr := recti{width: 3, height: 4}
	ccc := circlei{radius: 5}

	measure(rr)
	measure(ccc)

	fmt.Println("----------------------------")
	for _, i := range [] int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed", e)
		} else {
			fmt.Println("f1 worked", r)
		}
	}

	for _, i := range [] int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed", e)
		} else {
			fmt.Println("f2 worked", r)
		}
	}

	_, ee := f2(42)
	if ae, ok := ee.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
	fmt.Println("----------------------------")
	fc("direct")

	go fc("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	//等待输入
	//var input string
	//fmt.Scanln(&input)
	fmt.Println("done")

	fmt.Println("----------------------------")
	messages := make(chan string)
	go func() { messages <- "ping" }()
	msg := <-messages
	fmt.Println(msg)

	fmt.Println("----------------------------")
	bufmessages := make(chan string, 2)
	bufmessages <- "buffed1"
	bufmessages <- "buffed2"

	fmt.Println(<-bufmessages)
	fmt.Println(<-bufmessages)

	fmt.Println("----------------------------")
	done := make(chan bool, 1)
	go worker(done)

	<-done

	fmt.Println("----------------------------")
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)

	fmt.Println("----------------------------")
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second * 3)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	fmt.Println("----------------------------")
	cc1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		cc1 <- "result 1"
	}()

	select {
	case res := <-cc1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	cc2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		cc2 <- "result 2"
	}()

	select {
	case res := <-cc2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}

	fmt.Println("----------------------------")
	messages1 := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages1:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg11 := "hi"
	select {
	case messages1 <- msg11:
		fmt.Println("sent message", msg11)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages1:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

	fmt.Println("----------------------------")
	jobs := make(chan int, 5)
	done1 := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done1 <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("send job", j)
	}
	close(jobs)
	fmt.Println("send all jobs")

	<-done1

	fmt.Println("----------------------------")

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)
	for elem := range queue {
		fmt.Println(elem)
	}

	fmt.Println("----------------------------")
	timer1 := time.NewTimer(time.Second * 2)

	<-timer1.C
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	fmt.Println("----------------------------")
	ticker := time.NewTicker(time.Microsecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at ", t)
		}
	}()
	time.Sleep(time.Microsecond * 16000)
	ticker.Stop()
	fmt.Println("Ticker stopped")

	fmt.Println("----------------------------")
	jobs1 := make(chan int, 100)
	results1 := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker1(w, jobs1, results1)
	}

	for j := 1; j <= 9; j++ {
		jobs1 <- j
	}
	close(jobs1)
	for a := 1; a <= 9; a++ {
		<-results1
	}
	fmt.Println("----------------------------")

	request1 := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		request1 <- i
	}
	close(request1)

	limiter := time.Tick(time.Millisecond * 200)
	for req := range request1 {
		<-limiter
		fmt.Println("request ", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)
	for i := 1; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}

	fmt.Println("----------------------------")

	var ops uint64 = 0
	for i:=0;i<50;i++{
		go func() {
			atomic.AddUint64(&ops,1)
			runtime.Gosched()
		}()
	}
	time.Sleep(time.Second)
	opsFinal:=atomic.LoadUint64(&ops)
	fmt.Println("ops",opsFinal)

}

func worker1(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker ", id, " procession job ", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func plus(a int, b int) int {
	return a + b
}

func vals() (int, int) {
	return 3, 7
}

func sumf(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

type person struct {
	name string
	age  int
}

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.height * r.width
}

func (r *rect) perim() int {
	return 2*r.width + 2*r.height
}

type geometry interface {
	area() float64
	perim() float64
}

type recti struct {
	width, height float64
}

type circlei struct {
	radius float64
}

func (r recti) area() float64 {
	return r.width * r.height
}

func (r recti) perim() float64 {
	return 2*r.height + 2*r.height
}

func (c circlei) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circlei) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func fc(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

//https://books.studygolang.com/gobyexample/errors/
