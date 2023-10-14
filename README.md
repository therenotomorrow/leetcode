LeetCode
========

> [LeetCode problems](https://leetcode.com/problemset/all/) and how I [solved](./pkg/arch) them.

Structure
---------

 - [Internal](./internal) - all completed and refactored solved problems
 - [Again](./pkg/again) - need resolve because of loosing day streak or not enough knowledge
 - [Archive](./pkg/arch) - all solved problems that in progress refactoring

Development
-----------

After adding some code samples run [`code.sh`](./scripts/code.sh)

```shell
./scripts/code.sh
```

You can automate this action by calling `pre`

Testing
-------

```shell
go test -race ./...
```