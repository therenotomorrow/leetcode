def min_sub_array_len(target_sum: int, nums: list[int]) -> int:
    min_len = len(nums) + 1
    left = 0
    cur_sum = 0

    for right, num in enumerate(nums):
        cur_sum += num

        while cur_sum >= target_sum:
            min_len = min(min_len, right - left + 1)
            cur_sum -= nums[left]
            left += 1

    return 0 if min_len > len(nums) else min_len
