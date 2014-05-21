package strategies

import (
    "bollinger/bands"
)

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

func UpOnce(all []bands.Band) (bool) {
    for _,b := range(all){
        if b.Close >= b.Up {
            return true
        }
    }
    return false
}

func DownOnce(all []bands.Band) (bool) {
    for _,b := range(all){
        if b.Close <= b.Down {
            return true
        }
    }
    return false
}
