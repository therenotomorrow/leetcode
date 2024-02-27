# ---- two pointers
def reverse_string(chars: list[str]) -> None:
    left = 0
    right = len(chars) - 1

    while left < right:
        chars[left], chars[right] = chars[right], chars[left]
        left += 1
        right -= 1


def sorted_squares(nums: list[int]) -> list[int]:
    left = 0
    right = len(nums) - 1
    point = len(nums) - 1

    ans = [0 for _ in nums]

    while left <= right:
        l_pow = nums[left] * nums[left]
        r_pow = nums[right] * nums[right]

        if l_pow > r_pow:
            ans[point] = l_pow
            left += 1
        else:
            ans[point] = r_pow
            right -= 1

        point -= 1

    return ans


# ---- sliding window


def find_max_average(nums: list[int], window: int) -> float:
    curr = sum(nums[:window])
    ans = curr / window

    for idx, num in enumerate(nums[window:], start=window):
        curr += num - nums[idx - window]

        ans = max(ans, curr / window)

    return ans


def longest_ones(nums: list[int], max_flips: int) -> int:
    left = 0
    flips = 0
    ans = 0

    for right, num in enumerate(nums):
        if not num:
            flips += 1

        while flips > max_flips:
            if not nums[left]:
                flips -= 1

            left += 1

        ans = max(ans, right - left + 1)

    return ans


# ---- prefix sum


def running_sum(nums: list[int]) -> list[int]:
    prefix = [0 for _ in nums]
    prefix[0] = nums[0]

    for idx, num in enumerate(nums[1:], start=1):
        prefix[idx] = prefix[idx - 1] + num

    return prefix


def min_start_value(nums: list[int]) -> int:
    min_val = 0
    prefix = 0

    for num in nums:
        prefix += num
        min_val = min(prefix, min_val)

    return abs(min_val) + 1


def get_averages(nums: list[int], rad: int) -> list[int]:
    prefs = [nums[0]]

    for num in nums[1:]:
        prefs.append(prefs[-1] + num)

    ans = []

    for idx, _ in enumerate(nums):
        if idx < rad or idx + rad >= len(nums):
            mean = -1
        else:
            sum_ = prefs[idx + rad] - prefs[idx - rad] + nums[idx - rad]
            mean = sum_ // (2 * rad + 1)

        ans.append(mean)

    return ans
