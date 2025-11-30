package pingpong

import (
	"fmt"
	"sync"
)

//
//Exercise 1 — PingPong: First Steps With Goroutines
//--------------------------------------------------
//
//Background:
//-----------
//Goroutines are lightweight concurrent functions. Channels allow those goroutines
//to communicate by sending and receiving typed values.
//
//In this exercise you'll create two goroutines:
//  - one goroutine that sends the string "ping"
//  - another goroutine that sends the string "pong"
//
//Both should send their values into the SAME channel.
//
//Your task:
//----------
//Implement the function RunPingPong(count int) which:

func RunPingPong(count int) []string {
	ch := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)

	go sendPing(&wg, ch, count)
	go sendPong(&wg, ch, count)

	result := []string{}
	for i := 0; i < count*2; i++ {
		msg := <-ch
		result = append(result, msg)
		fmt.Println(msg)
	}
	wg.Wait()
	return result
}

func sendPing(wg *sync.WaitGroup, ch chan<- string, times int) {
	defer wg.Done()
	for i := 0; i < times; i++ {
		ch <- "ping"
	}
}

func sendPong(wg *sync.WaitGroup, ch chan<- string, times int) {
	defer wg.Done()
	for i := 0; i < times; i++ {
		ch <- "pong"
	}
}

//1. Creates a channel of strings.
//2. Launches:
//      - a goroutine that sends "ping" into the channel 'count' times.
//      - a goroutine that sends "pong" into the channel 'count' times.
//3. Collects EXACTLY 2 * count messages from the channel.
//4. Closes the channel **after** both goroutines have finished sending.
//5. Returns a slice containing ALL received messages (order does NOT matter).
//
//Important rules:
//----------------
//- Use goroutines.
//- Use a channel.
//- DO NOT use time.Sleep anywhere.
//- DO NOT limit receiving order — goroutine scheduling is non-deterministic.
//- You may use a sync.WaitGroup OR channel signalling to know when you're done.
//
//Example return slice for count = 2 (order may vary):
//    ["ping", "pong", "ping", "pong"]
//
//Your function must return all messages.
//
//TODO:
//-----
//Implement RunPingPong(count int) so that the tests pass.
//*/
//
//func RunPingPong(count int) []string {
//	// TODO: implement
//	return nil
//}
//
