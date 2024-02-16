# ---- two pointers

def reverse_string(s: list[str]) -> None:
    left = 0
    right = len(s) - 1

    while left < right:
        s[left], s[right] = s[right], s[left]
        left += 1
        right -= 1


def sorted_squares(nums: list[int]) -> list[int]:
    left = 0
    right = len(nums) - 1
    point = len(nums) - 1

    ans = [0] * len(nums)

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

def find_max_average(nums: list[int], k: int) -> float:
    curr = sum(nums[:k])
    ans = curr / k

    for i in range(k, len(nums)):
        curr += nums[i] - nums[i - k]

        ans = max(ans, curr / k)

    return ans


def longest_ones(nums: list[int], k: int) -> int:
    left = flips = ans = 0

    for right in range(len(nums)):
        if not nums[right]:
            flips += 1

        while flips > k:
            if not nums[left]:
                flips -= 1

            left += 1

        ans = max(ans, right - left + 1)

    return ans


# ---- prefix sum

def running_sum(nums: list[int]) -> list[int]:
    prefix = [0] * len(nums)
    prefix[0] = nums[0]

    for i in range(1, len(nums)):
        prefix[i] = prefix[i - 1] + nums[i]

    return prefix


def min_start_value(nums: list[int]) -> int:
    val = 0
    prefix = 0

    for num in nums:
        prefix += num
        val = min(prefix, val)

    return abs(val) + 1


def get_averages(nums: list[int], k: int) -> list[int]:
    prefix = [nums[0]]

    for num in nums[1:]:
        prefix.append(prefix[-1] + num)

    ans = []

    for i in range(len(nums)):
        if i < k or i + k >= len(nums):
            ans.append(-1)
        else:
            ans.append((prefix[i + k] - prefix[i - k] + nums[i - k]) // (2 * k + 1))

    return ans
