## idea
### normal way Kadane's Algorithm
```
Initialize:
    result = INT_MIN
    max_ending_here = 0

Loop for each element of the array
  (a) max_ending_here = max_ending_here + a[i]
  (b) if(result < max_ending_here)
            result = max_ending_here
  (c) if(max_ending_here < 0)
            max_ending_here = 0
return result
```