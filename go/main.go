package main

import (
    "bollinger/bands"
    "bollinger/plot"
    "bollinger/strategies"
    "flag"
    "fmt"
    "sync"
)

// Calculate investment suggestion using strategy on each given symbol
func parallelSuggest(symbols []string, strategy string){
    // default strategy
    fn := strategies.MoreDown
    
    switch strategy {
        case "moredown": fn = strategies.MoreDown
        case "moreup": fn = strategies.MoreUp
        case "uponce": fn = strategies.UpOnce
        case "downonce": fn = strategies.DownOnce
        default: fmt.Println("Invalid strategy: ", strategy, "use any of: moredown, moreup, uponce, downonce"); return;
    }
    
    // process each symbol in parallel
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
    
    // wait for execution of all goroutines
    wg.Wait()
}

// Executes plot on each given symbol
func parallelPlot(symbols []string){
    
    // process each symbol in parallel
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
    
    // wait for execution of all goroutines
    wg.Wait()
}

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
        parallelPlot(symbols)
    } else {
        parallelSuggest(symbols, *args)
    }
}
