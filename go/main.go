package main

import (
    "bollinger/bands"
    "bollinger/plot"
    "bollinger/strategies"
    "flag"
    "fmt"
    "sync"
)

func main() {
    args := flag.String("s", "moredown", "suggests investement using given strategy, available: moredown, moreup, uponce, downonce")
    argp := flag.Bool("p", false, "calculates and saves bollinger bands as SYMBOL.png")
    flag.Parse()
    
    symbols := flag.Args()
    
    if len(symbols) == 0 {
        fmt.Println("at least one symbol is required")
        return
    }
    
    if *argp {
        var wg sync.WaitGroup
        wg.Add(len(symbols))
        for _,symbol := range(symbols){
            go func(symbol string){
                all := bands.All(symbol)
                plot.PlotBands(symbol, all)
                fmt.Println("Plot [", symbol, "] OK")
                wg.Done()
            }(symbol)
        }
        wg.Wait()
        return
    } else {
        fn := strategies.MoreDown
        
        switch *args {
            case "moredown": fn = strategies.MoreDown
            case "moreup": fn = strategies.MoreUp
            default: fmt.Println("Invalid strategy: ", *args, "use any of: moredown, moreup, uponce, downonce"); return;
        }
        
        // Run each iteration of this loop as goroutines
        var wg sync.WaitGroup
        wg.Add(len(symbols))
        for _,symbol := range(symbols){
            go func(symbol string){
                all := bands.All(symbol)
                if fn(all) {
                    fmt.Println("Invest [", symbol, "] YES")
                } else {
                    fmt.Println("Invest [", symbol, "] NO")
                }
                wg.Done()
            }(symbol)
        }
        wg.Wait()
    }
}
