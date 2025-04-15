class NumArray:
    def __init__(self, nums: list[int]):
        self._prefix = [0]
        for num in nums:
            self._prefix.append(num + self._prefix[-1])

    def sum_range(self, left: int, right: int) -> int:
        return self._prefix[right + 1] - self._prefix[left]
