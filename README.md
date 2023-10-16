LeetCode
========

> [LeetCode problems](https://leetcode.com/problemset/all/) and how I solved them.

Structure
---------

- [Internal](./internal) - all solved and refactored problems
- [Again](./pkg/again) - need resolve because of loosing day streak or not enough knowledge
- [Archive](./pkg/arch) - all solved problems that in progress refactoring

Development
-----------

After adding some code samples run [`code.sh`](./scripts/code.sh)

```shell
./scripts/code.sh
```

You can automate this action by calling [`pre-commit.sh`](./scripts/pre-commit.sh)

Testing
-------

Coverage tests for CI/CD.

```shell
./scripts/test.sh
```

Fast/Race tests.

```shell
./scripts/test.sh [-fast|-race]
```

Completed problems
------------------

1. [Two Sum](https://leetcode.com/problems/two-sum/): ???
3. [Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/): ???
4. [Median of Two Sorted Arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/): ???
7. [Reverse Integer](https://leetcode.com/problems/reverse-integer/): ???
8. [String to Integer (atoi)](https://leetcode.com/problems/string-to-integer-atoi/): ???
9. [Palindrome Number](https://leetcode.com/problems/palindrome-number/): ???
11. [Container With Most Water](https://leetcode.com/problems/container-with-most-water/): ???
13. [Roman to Integer](https://leetcode.com/problems/roman-to-integer/): ???
14. [Longest Common Prefix](https://leetcode.com/problems/longest-common-prefix/): ???
19. [Remove Nth Node From End of List](https://leetcode.com/problems/remove-nth-node-from-end-of-list/): ???

119. [Pascal's Triangle II](https://leetcode.com/problems/pascals-triangle-ii/): [`getrow.go`](./internal/getrow/getrow.go)
1269. [Number of Ways to Stay in the Same Place After Some Steps](https://leetcode.com/problems/number-of-ways-to-stay-in-the-same-place-after-some-steps/): [`numways`](./internal/numways/numways.go)
