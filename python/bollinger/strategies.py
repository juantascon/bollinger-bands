from . import bands

# each of these classes represent independent investment strategies

# invest if on the last 20 days close price passes at least once over upper band
class UpOnce:
    name = "uponce"
    
    def __init__(self, bands):
        self.bands = bands
    
    def invest(self):
        result = self.bands.all()
        
        for b in result:
            if b.get("close") >= b.get("up"):
                return True
        
        return False

# invest if on the last 20 days close price passes at least once under lower band
class DownOnce:
    name = "downonce"
    
    def __init__(self, bands):
        self.bands = bands
    
    def invest(self):
        result = self.bands.all()
        
        for b in result:
            if b.get("close") <= b.get("down"):
                return True
        
        return False

# invest if on the last 20 days the number of times close price passes
# over the uppper band is greater than the number of times close price passes
# under the lower band
class MoreUp:
    name = "moreup"
    
    def __init__(self, bands):
        self.bands = bands
    
    def invest(self):
        result = self.bands.all()
        
        up = 0
        down = 0
        for b in result:
            if b.get("close") >= b.get("up"):
                up = up+1
                
            if b.get("close") <= b.get("down"):
                down = down+1
        
        #print("up:%d, down: %d"%(up, down))
        
        if up > down:
            return True
        else:
            return False

# invest if on the last 20 days the number of times close price passes
# under the lower band is greater than the number of times close price passes
# over the uppper band
class MoreDown:
    name = "moredown"
    
    def __init__(self, bands):
        self.bands = bands
    
    def invest(self):
        result = self.bands.all()
        
        up = 0
        down = 0
        for b in result:
            if b.get("close") >= b.get("up"):
                up = up+1
                
            if b.get("close") <= b.get("down"):
                down = down+1
        
        #print("up:%d, down: %d"%(up, down))
        
        if down > up:
            return True
        else:
            return False
