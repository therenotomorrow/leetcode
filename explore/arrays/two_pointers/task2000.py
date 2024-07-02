def reverse_prefix(word: str, ch: str) -> str:
    until = -1

    for idx, let in enumerate(word):
        if ch == let:
            until = idx
            break

    if until == -1:
        return word

    prefix = word[until::-1]
    until += 1
    suffix = word[until::]

    return prefix + suffix
