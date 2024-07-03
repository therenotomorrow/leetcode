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

Counting
--------

```python
import collections

def counting(s: str, k: int) -> int:
    ans = 0
    dic = collections.defaultdict(int)
    left = 0

    for right in range(len(s)):
        dic[s[right]] += 1

        while len(dic) > k:
            dic[s[left]] -= 1

            if dic[s[left]] == 0:
                del dic[s[left]]

            left += 1
        
        ans = max(ans, right - left + 1)
    
    return ans
```

#### Notes

Anytime you need to count anything, think about using a hash map to do it.
