Two pointers
------------

#### (1) Use two pointers in loop from different edges on the same iterable object.

```python
def two_pointers(nums: list[int], target: int) -> bool:
    left = 0
    right = len(nums) - 1

    while left < right:
        curr = nums[left] + nums[right]

        if curr == target:
            return True

        if curr > target:
            right -= 1
        else:
            left += 1

    return False
```

```text
function fn(arr):
    left = 0
    right = arr.length - 1

    while left < right:
        Do some logic here depending on the problem
        Do some more logic here to decide on one of the following:
            1. left++
            2. right--
            3. Both left++ and right--
```

#### (2) Move along both inputs simultaneously until all elements have been checked.

```python
def two_pointers(arr1: list[int], arr2: list[int]) -> list[int]:
    ans = []
    i = j = 0

    while i < len(arr1) and j < len(arr2):
        if arr1[i] < arr2[j]:
            ans.append(arr1[i])
            i += 1
        else:
            ans.append(arr2[j])
            j += 1

    while i < len(arr1):
        ans.append(arr1[i])
        i += 1

    while j < len(arr2):
        ans.append(arr2[j])
        j += 1

    return ans
```

```text
function fn(arr1, arr2):
    i = j = 0
    
    while i < arr1.length AND j < arr2.length:
        Do some logic here depending on the problem
        Do some more logic here to decide on one of the following:
            1. i++
            2. j++
            3. Both i++ and j++

    // Step 4: make sure both iterables are exhausted
    // Note that only one of these loops would run
    while i < arr1.length:
        Do some logic here depending on the problem
        i++

    while j < arr2.length:
        Do some logic here depending on the problem
        j++
```

#### Notes

Sometimes there will be "three pointers", sometimes the approach with the pointers (1) will not start from 0 or ends with last element. 
Sometimes (2) will use only with one source. The basic idea is to use two pointers to traverse the data structure in a specific way to achieve a particular goal efficiently.
