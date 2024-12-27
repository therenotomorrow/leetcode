def reverse_only_letters(text: str) -> str:
    left = 0
    right = len(text) - 1
    chars: list[str] = list(text)

    while left < right:
        if not chars[left].isalpha():
            left += 1
            continue

        if not chars[right].isalpha():
            right -= 1
            continue

        chars[left], chars[right] = chars[right], chars[left]  # noqa:WPS414
        left += 1
        right -= 1

    return ''.join(chars)
