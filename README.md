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

1269. [Number of Ways to Stay in the Same Place After Some Steps](https://leetcode.com/problems/number-of-ways-to-stay-in-the-same-place-after-some-steps/): [`numways`](./internal/numways)
