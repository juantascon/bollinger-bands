package strategies

// several investment strategies

import (
    "bollinger/bands"
)

// invest if on the last 20 days the number of times close price passes
// under the lower band is greater than the number of times close price passes
// over the uppper band
func MoreDown(all []bands.Band) (bool) {
    up := 0
    down := 0
    
    for _,b := range(all){
        if b.Close >= b.Up {
            up++
        } else if b.Close <= b.Down{
            down++
        }
    }
    
    if down > up {
        return true
    } else {
        return false
    }
}

// invest if on the last 20 days the number of times close price passes
// over the uppper band is greater than the number of times close price passes
// under the lower band
func MoreUp(all []bands.Band) (bool) {
    up := 0
    down := 0
    
    for _,b := range(all){
        if b.Close >= b.Up {
            up++
        } else if b.Close <= b.Down{
            down++
        }
    }
    
    if up > down {
        return true
    } else {
        return false
    }
}

// invest if on the last 20 days close price passes at least once over upper band
func UpOnce(all []bands.Band) (bool) {
    for _,b := range(all){
        if b.Close >= b.Up {
            return true
        }
    }
    return false
}

// invest if on the last 20 days close price passes at least once under lower band
func DownOnce(all []bands.Band) (bool) {
    for _,b := range(all){
        if b.Close <= b.Down {
            return true
        }
    }
    return false
}
