from .lib import ystockquote
from . import settings

from datetime import date, timedelta, datetime
import math

class Bands:
    def __init__(self, symbol):
        self.symbol = symbol
        self.prices = []
        
    # given a start date and a number of days it
    # calculates the "floor" considering weekdays only
    def _weekdays_interval_past(self, start_date, days):
        start = start_date
        while start.weekday() >= 5:
            start -= timedelta(days=1)
        
        end = start
        while days > 0:
            end -= timedelta(days=1)
            if end.weekday() < 5:
                days -= 1
        
        return start, end
    
    def fetch(self):
        # calculate interval dates, considers only weekdays
        _start, _end = self._weekdays_interval_past(settings.TODAY, settings.GRAPH_DAYS)
        _tmp, _end = self._weekdays_interval_past(_end, settings.BOLLINGER_DAYS)
        self.start_date = _start
        self.end_date = _end
        
        # call yahoo finance api
        historical = ystockquote.get_historical_prices(self.symbol, str(self.end_date), str(self.start_date))
        
        # remove all prices from previous calculations
        del self.prices[:]
        
        # sort first, historical is a regular dict
        for k in sorted(historical.keys(), reverse=True): 
            # close value sorted by date
            close = float(historical[k]["Close"])
            
            #prices[0] = today, prices[1] = yesterday
            self.prices.append(close)
    
    # calculate bands for one day
    def one(self, index):
        values = self.prices[index:index+settings.BOLLINGER_DAYS]
        
        # simple moving average
        sma = math.fsum(values) / len(values)
        
        squares = [ math.pow((v-sma),2) for v in values ]
        dev = math.sqrt( (math.fsum(squares) / len(squares)) )
        
        # upper band
        up = sma + (settings.STANDARD_DEVIATIONS * dev)
        
        # lower band
        down = sma - (settings.STANDARD_DEVIATIONS * dev)
        
        return { "close": values[0], "up": up, "middle": sma, "down": down, "index": index }
    
    # calculate bands for all days
    def all(self):
        return [ self.one(day) for day in range(0, settings.GRAPH_DAYS) ]
