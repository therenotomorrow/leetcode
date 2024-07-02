class NumArray:
    def __init__(self, nums: list[int]):
        self._prefixes = [0]
        for num in nums:
            self._prefixes.append(num + self._prefixes[-1])

    def sum_range(self, left: int, right: int) -> int:
        return self._prefixes[right + 1] - self._prefixes[left]
