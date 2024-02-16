package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	count := 10000000
	var startTime time.Time
	fmt.Println(startTime)

	wg := sync.WaitGroup{}
	wg.Add(count)
	startTime = time.Now()
	for i := 1; i <= count; i++ {
		wg.Done()
		fmt.Println(i)
	}
	endTime := time.Now()
	wg.Wait()

	fmt.Println(endTime)
	t := endTime.Sub(startTime)
	fmt.Println(t)
}



// package main

// import (
// 	"fmt"
// 	"os"
// 	"sync"
// 	"syscall"
// 	"time"
// )

// func main() {
//     kernel32 := syscall.NewLazyDLL("kernel32.dll")
//     setPriorityClass := kernel32.NewProc("SetPriorityClass")
//     handle := syscall.Handle(os.Getpid())
//     setPriorityClass.Call(uintptr(handle), uintptr(syscall.HIGH_PRIORITY_CLASS))
    
//     count := 10000000
//     var startTime time.Time
//     fmt.Println(startTime)
    
//     wg := sync.WaitGroup{}
//     wg.Add(count)
//     startTime = time.Now()
//     for i := 1; i <= count; i++ {
//         wg.Done()
//         fmt.Println(i)
//     }
//     endTime := time.Now()
//     wg.Wait()
    
//     fmt.Println(endTime)
//     t := endTime.Sub(startTime)
//     fmt.Println(t)
// }
