import operator


def move_zeroes(nums: list[int]) -> None:
    idx = 0

    for num in filter(operator.truth, nums):
        nums[idx] = num
        idx += 1

    while idx < len(nums):
        nums[idx] = 0
        idx += 1
