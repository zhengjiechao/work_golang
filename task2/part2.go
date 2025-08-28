package task2

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup

var lock sync.Mutex

// 指针
// 1.1
func Pointer1(num *int) {
	*num += 10
}

// 1.2
func Pointer2(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		(*arr)[i] *= 2
	}
}

// 2.Goroutine
// 2.1
func Goroutine1() {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i += 2 {
			fmt.Println(i, "--协程1")
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			fmt.Println(i, "==协程2")
		}
	}()

	wg.Wait()
}

// 2.2
func Goroutine2(arr []func()) {
	wg.Add(len(arr))
	for _, f := range arr {
		go func() {
			defer wg.Done()
			start := time.Now().UnixMilli()
			f()
			fmt.Printf("耗时：%v 毫秒\n", time.Now().UnixMilli()-start)
		}()
	}
	wg.Wait()
}

// 3.面向对象
// 3.1
type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
}

func (r Rectangle) Area() {
	fmt.Println("我是Rectangle的Area")
}

func (r Rectangle) Perimeter() {
	fmt.Println("我是Rectangle的Perimeter")
}

type Circle struct {
}

func (r Circle) Area() {
	fmt.Println("我是Circle的Area")
}

func (r Circle) Perimeter() {
	fmt.Println("我是Circle的Perimeter")
}

// 3.2
type Person struct {
	Name string
	Age  int
}

type Employee struct {
	EmployeeID int
	Person
}

func (e Employee) PrintInfo() {
	fmt.Println("EmployeeID:", e.EmployeeID)
	fmt.Println("Name:", e.Name)
	fmt.Println("Age:", e.Age)
}

// 4.Channel
// 4.1
func Channel1() {
	ch := make(chan int, 5)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			ch <- i
			fmt.Println("发送一个：", i)
		}
		close(ch)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println("接收到一个：", v)
		}
	}()

	wg.Wait()
}

// 4.2
func Channel2() {
	ch := make(chan int, 10)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			ch <- i
			fmt.Println("生产者发送一个：", i)
		}
		close(ch)
	}()

	wg.Add(1)
	go func() {
		wg.Done()
		for v := range ch {
			fmt.Println("消费者接收到一个：", v)
		}
	}()

	wg.Wait()
}

// 5.锁机制
// 5.1
func Lock1() {
	num := 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				lock.Lock()
				num += 1
				lock.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(num)
}

// 5.2
func Lock2() {
	var num int64 = 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&num, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println(num)
}
