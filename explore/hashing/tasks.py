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
