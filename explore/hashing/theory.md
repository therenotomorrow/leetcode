Checking for existence
----------------------

```python
def checking_for_existence(nums: list[int], target: int) -> list[int]:
    dic = {}

    for i in range(len(nums)):
        num = nums[i]
        complement = target - num
        if complement in dic: # This operation is O(1)!
            return [i, dic[complement]]
        
        dic[num] = i
    
    return [-1, -1]
```

#### Notes

Anytime find algorithm running `if ... in ...`, then consider using a hash map or set to store elements.
`value % x` is the modulo operation, and makes sure the final converted value will be in the range `[0, x - 1]`.
