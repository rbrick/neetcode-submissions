class MedianFinder:

    def __init__(self):
        self.min_heap, self.max_heap = [], []
        
    def addNum(self, num: int) -> None:
        if len(self.min_heap) == len(self.max_heap): 
            # try a pushpop
            heapq.heappush_max(self.max_heap, heapq.heappushpop(self.min_heap, num))
        else:
            heapq.heappush(self.min_heap, heapq.heappushpop_max(self.max_heap, num))
        
    def findMedian(self) -> float:
        if len(self.min_heap) == len(self.max_heap): return (self.min_heap[0] + self.max_heap[0]) / 2
        return self.max_heap[0]