def is_vowel(let: str) -> bool:
    return let in {'a', 'e', 'i', 'o', 'u'}


def max_vowels(text: str, window: int) -> int:
    left = 0
    max_cnt = sum(map(is_vowel, text[:window]))
    cur_cnt = max_cnt

    for let in text[window:]:
        cur_cnt -= is_vowel(text[left])
        cur_cnt += is_vowel(let)
        max_cnt = max(max_cnt, cur_cnt)
        left += 1

    return max_cnt
