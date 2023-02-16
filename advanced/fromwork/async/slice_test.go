package async

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var testSlice []int
var mutex = &sync.Mutex{}

func Test_addSlice(t *testing.T) {

	start := time.Now()
	num := 10000

	var wg = &sync.WaitGroup{}

	wg.Add(num)
	for i := 0; i < num; i++ {
		go add(i, wg)
	}
	wg.Wait()

	fmt.Println(len(testSlice), cap(testSlice))

	fmt.Println(time.Since(start))
}

func add(i int, wg *sync.WaitGroup) {
	mutex.Lock()
	defer mutex.Unlock()
	testSlice = append(testSlice, i)
	wg.Done()
}

func Test_addSliceChannel(t *testing.T) {
	start := time.Now()
	num := 10000

	var wg = &sync.WaitGroup{}

	wg.Add(num)

	chan1 := make(chan int)
	for i := 0; i < num; i++ {
		i := i
		go func() {
			chan1 <- i
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(chan1)
	}()

	for v := range chan1 {
		testSlice = append(testSlice, v)
	}

	fmt.Println(len(testSlice), cap(testSlice))

	fmt.Println(time.Since(start))
}

func TestSafeMap(t *testing.T) {
	test := sync.Map{}

	test.Store("test", 22)
	test.Store(1, 333)

	test.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})

	val, _ := test.Load("test")
	fmt.Println(val)
}

func TestSafeSlice(t *testing.T) {
	test := safeSlice{}

	start := time.Now()
	num := 10000

	var wg = &sync.WaitGroup{}

	wg.Add(num)

	for i := 0; i < num; i++ {
		i := i
		go func() {
			test.mutex.Lock()
			test.data = append(test.data, i)
			test.mutex.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(len(test.data), cap(test.data))

	fmt.Println(time.Since(start))
}
