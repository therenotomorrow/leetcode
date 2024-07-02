def equal_substring(text: str, tmpl: str, max_cost: int) -> int:
    cur_cost = 0
    max_len = 0
    left = 0

    for right, let in enumerate(text):
        cur_cost += abs(ord(let) - ord(tmpl[right]))

        while cur_cost > max_cost:
            cur_cost -= abs(ord(text[left]) - ord(tmpl[left]))
            left += 1

        max_len = max(max_len, right - left + 1)

    return max_len
