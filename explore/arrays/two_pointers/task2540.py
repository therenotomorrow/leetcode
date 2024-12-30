def get_common(nums1: list[int], nums2: list[int]) -> int:
    i1 = 0
    i2 = 0

    while i1 < len(nums1) and i2 < len(nums2):
        if nums1[i1] > nums2[i2]:
            i2 += 1
        elif nums2[i2] > nums1[i1]:
            i1 += 1
        else:
            return nums1[i1]

    return -1
