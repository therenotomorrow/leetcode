def reverse_words(text: str) -> str:
    space = -1
    chars: list[str] = list(text)

    for idx in range(len(chars) + 1):
        if idx == len(text) or chars[idx].isspace():
            start = space + 1
            end = idx - 1

            while start < end:
                chars[start], chars[end] = chars[end], chars[start]
                start += 1
                end -= 1

            space = idx

    return ''.join(chars)
