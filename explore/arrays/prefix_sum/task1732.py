def largest_altitude(gain: list[int]) -> int:
    max_alt = 0
    cur_alt = 0

    for num in gain:
        cur_alt += num
        max_alt = max(max_alt, cur_alt)

    return max_alt
