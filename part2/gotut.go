package main

// Part 22&23 Channels & Buffering and Iteration over Channels

// var wg sync.WaitGroup

// func foo(c chan int, someValue int) {
// 	defer wg.Done()
// 	c <- someValue * 5
// }

// func main() {
// 	fooVal := make(chan int, 10)
// 	// go foo(fooVal, 5)
// 	// go foo(fooVal, 3)

// 	// // v1 := <-fooVal
// 	// // v2 := <-fooVal

// 	// v1, v2 := <-fooVal, <- fooVal

// 	// fmt.Println(v1, v2)

// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go foo(fooVal, i)
// 	}

// 	wg.Wait()
// 	close(fooVal)

// 	for item := range fooVal {
// 		fmt.Println(item)
// 	}

// 	//time.Sleep(time.Second * 2)

// }

// // Part 21 Panic and Recover

// var wg sync.WaitGroup

// func cleanup() {
// 	defer wg.Done()
// 	if r := recover(); r != nil {
// 		fmt.Println("Recovered in clearnup: ", r)
// 	}
// }

// func say(s string) {
// 	defer cleanup()
// 	for i := 0; i < 3; i++ {
// 		time.Sleep(time.Millisecond * 100)
// 		fmt.Println(s)
// 		if i == 2 {
// 			panic("Oh dear, a 2")
// 		}
// 	}
// }

// func main() {
// 	wg.Add(1)
// 	go say("Hey")
// 	wg.Add(1)
// 	go say("There")
// 	wg.Wait()
// }

// PART 20 Defer

// func foo() {
// 	// defer fmt.Println("Done!")
// 	// defer fmt.Println("Are we done?")
// 	// fmt.Println("Doing some stuff, who knows what?")

// 	for i := 0; i < 5; i++ {
// 		defer fmt.Println(i)
// 	}
// }

// func main() {
// 	foo()
// }

// PART 18,19, 20 Goroutine, Goroutine Synchronization

// var wg sync.WaitGroup

// func say(s string) {
// 	defer wg.Done()
// 	for i := 0; i < 3; i++ {
// 		time.Sleep(time.Millisecond * 100)
// 		fmt.Println(s)
// 	}
// }

// func main() {
// 	wg.Add(1)
// 	go say("Hey")
// 	wg.Add(1)
// 	go say("There")
// 	wg.Wait()

// 	// wg.Add(1)
// 	// go say("Hi")
// 	// wg.Wait()

// 	//time.Sleep(time.Second )
// }
