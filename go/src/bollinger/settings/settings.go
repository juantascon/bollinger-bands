package settings

import (
    "time"
)

// number of days to calculate SMA
const SMA_DAYS = 20

// number of days to display on graph
const GRAPH_DAYS = 20

// standard deviation multiplier, default 2
// more info: http://www.great-trades.com/Help/bollinger%20bands%20calculation.htm
const STANDARD_DEVIATIONS = 2

// start date
var TODAY = time.Now()
