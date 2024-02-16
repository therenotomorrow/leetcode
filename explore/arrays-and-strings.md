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

Sometimes there will be "three pointers", sometimes the approach with the pointers (1) will not start from 0 or ends
with last element.
Sometimes (2) will use only with one source. The basic idea is to use two pointers to traverse the data structure in a
specific way to achieve a particular goal efficiently.

Sliding window
--------------

A sliding window is actually implemented using two pointers.
Given an array, a subarray is a contiguous section of the array.
A sliding window could be dynamic or fixed size.

1. The problem will either define criteria that make a subarray "valid". A constraint metric and a numeric restriction
   on the constraint metric.
2. The problem will define what makes a subarray better than another or finding the number of valid subarrays.

```python
def sliding_window(nums: list[int], k: int) -> int:
    left = curr = ans = 0

    for right in range(len(nums)):
        curr += nums[right]

        while curr > k:
            curr -= nums[left]
            left += 1

        ans = max(ans, right - left + 1)

    return ans
```

```text
function fn(arr):
    left = 0
    for (int right = 0; right < arr.length; right++):
        Do some logic to "add" element at arr[right] to window

        while WINDOW_IS_INVALID:
            Do some logic to "remove" element at arr[left] from window
            left++

        Do some logic to update the answer
```

#### Notes

The number of valid windows ending at index right is equal to the size of the window, which we know
is `right - left + 1`.
The sliding window approach addresses this inefficiency by intelligently moving a window across the input data, avoiding
unnecessary re-computation of overlapping sub-problems.

Prefix sum
----------

Prefix sum is a technique that can be used on arrays (of numbers) - problem involves sums of a subarray.
The idea is to create an array prefix where `prefix[i]` is the sum of all elements up to the index `i` (inclusive).
Sum of subarray from `i` to `j` inclusive:

1. `prefix[j] - prefix[i - 1]`
2. `prefix[j] - prefix[i] + nums[i]`

```python
def prefix_sum(nums: list[int], queries: list[tuple[int, int]], limit: int) -> list[bool]:
    prefix = [nums[0]]
    for i in range(1, len(nums)):
        prefix.append(nums[i] + prefix[i-1])

    ans = []
    for x, y in queries:
        curr = prefix[y] - prefix[x] + nums[x]  # prefix[j] - prefix[i] + nums[i]
        ans.append(curr < limit)

    return ans
```

```text
prefix = [nums[0]]
for (int i = 1; i < nums.length; i++)
    prefix.append(nums[i] + prefix[prefix.length - 1])
```

#### Notes

Building a prefix sum is a form of pre-processing. Pre-processing is a useful 
strategy in a variety of problems where we store pre-computed data in a data 
structure before running the main logic of our algorithm.

Definitions
-----------

#### Subarrays/substrings

A subarray or substring is a contiguous section of an array or string. 

#### Subsequences

A subsequence is a set of elements of an array/string that 
keeps the same relative order but doesn't need to be contiguous.

#### Subsets

A subset is any set of elements from the original array or string. 
The order doesn't matter and neither do the elements being beside each other.
Subsets that contain the same elements are considered the same.
