// [Timers](timers) are for when you want to do
// something once in the future - _tickers_ are for when
// you want to do something repeatedly（反复地） at regular（有规律的）
// intervals（间隔）. Here's an example of a ticker that ticks
// periodically（定期地） until we stop it.

package main

import "time"
import "fmt"

func main() {

    // Tickers use a similar（类似的） mechanism（机制） to timers: a
    // channel that is sent values. Here we'll use the
    // `range` builtin on（内置在） the channel to iterate over（遍历）
    // the values as they arrive every 500ms.
    ticker := time.NewTicker(time.Millisecond * 500)
    go func() {
        for t := range ticker.C {
            fmt.Println("Tick at", t)
        }
    }()

    // Tickers can be stopped like timers. Once a ticker
    // is stopped it won't receive any more values on its
    // channel. We'll stop ours after 1600ms.
    time.Sleep(time.Millisecond * 1600)
    ticker.Stop()
    fmt.Println("Ticker stopped")
}
