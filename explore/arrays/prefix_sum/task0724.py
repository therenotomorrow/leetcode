def pivot_index(nums: list[int]) -> int:
    left = 0
    right = sum(nums)

    for idx, num in enumerate(nums):
        right -= num

        if left == right:
            return idx

        left += num

    return -1
