import collections

# ---- checking for existence


def check_if_pangram(sentence: str) -> bool:
    uniq = set()
    size = 26

    for letter in sentence:
        uniq.add(letter)

        if len(uniq) == size:
            return True

    return False


def missing_number(nums: list[int]) -> int:
    uniq = set(nums)

    for num in range(len(nums) + 1):
        if num not in uniq:
            return num

    return -1


def count_elements(arr: list[int]) -> int:
    uniq = set(arr)
    cnt = 0

    for num in arr:
        if num + 1 in uniq:
            cnt += 1

    return cnt


# ---- counting


def find_winners(matches: list[list[int]]) -> list[list[int]]:
    cnt: dict[int, int] = collections.defaultdict(int)

    for winner, loser in matches:
        cnt.setdefault(winner, 0)
        cnt[loser] -= 1

    ans: list[list[int]] = [[], []]

    for player, score in cnt.items():
        match score:
            case 0:
                ans[0].append(player)
            case -1:
                ans[1].append(player)

    ans[0].sort()
    ans[1].sort()

    return ans


def largest_unique_number(nums: list[int]) -> int:
    ans = -1
    cnt: dict[int, int] = collections.defaultdict(int)

    for number in nums:
        cnt[number] += 1

    for num, times in cnt.items():
        if times > 1:
            continue

        ans = max(ans, num)

    return ans


def max_number_of_balloons(text: str) -> int:  # noqa:WPS231
    cnt: dict[str, int] = collections.defaultdict(int)

    for char in text:
        cnt[char] += 1

    matches = [0, 0, 0, 0, 0]

    for let, times in cnt.items():
        match let:
            case 'a':
                idx = 0
            case 'b':
                idx = 1
            case 'n':
                idx = 2
            case 'o':
                idx = 3
                times //= 2
            case 'l':
                idx = 4
                times //= 2
            case _:
                continue

        matches[idx] += times

    return min(matches)


def find_max_length(nums: list[int]) -> int:
    ans = 0
    sum_ = 0

    cnt = collections.defaultdict(int)
    cnt[0] = -1  # we don't know the start index

    for idx, num in enumerate(nums):
        if num:
            sum_ += 1
        else:
            sum_ -= 1

        if cnt.get(sum_):
            ans = max(ans, idx - cnt[sum_])
        else:
            cnt[sum_] = idx

    return ans
