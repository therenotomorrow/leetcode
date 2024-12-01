import collections


def first_star(list1: list[int], list2: list[int]) -> int:
    list1.sort()
    list2.sort()

    return sum(abs(list2[idx] - left) for idx, left in enumerate(list1))


def second_star(list1: list[int], list2: list[int]) -> int:
    count = collections.Counter(list2)

    return sum(num * count.get(num, 0) for num in list1)
