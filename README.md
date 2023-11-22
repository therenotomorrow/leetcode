LeetCode
========

[LeetCode problems](https://leetcode.com/problemset/all/) and how
I [solved them](https://leetcode.com/therenotomorrow/).

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

```shell
./scripts/test.sh [-fast|-race|-cover]
```

Completed explores
------------------

* [The LeetCode Beginner's Guide](https://leetcode.com/explore/featured/card/the-leetcode-beginners-guide/)
* [Array and String. Introduction to Data Structure](https://leetcode.com/explore/learn/card/array-and-string/)
* [Top Interview Questions. Easy Collection](https://leetcode.com/explore/featured/card/top-interview-questions-easy/)

Completed problems
------------------

|  #   | Problem                                                                                                                                               | Solution                                                                          |
|:----:|:------------------------------------------------------------------------------------------------------------------------------------------------------|:----------------------------------------------------------------------------------|
|  1   | [Two Sum](https://leetcode.com/problems/two-sum/)                                                                                                     | [`twosum.go`](./internal/solutions/twosum.go)                                     |
|  2   | [Add Two Numbers](https://leetcode.com/problems/add-two-numbers/)                                                                                     | [`addtwonumbers.go`](./internal/solutions/addtwonumbers.go)                       |
|  3   | [Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/)                       | [`lengthoflongestsubstring.go`](./internal/solutions/lengthoflongestsubstring.go) |
|  4   | [Median of Two Sorted Arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/)                                                             | [`findmediansortedarrays.go`](./internal/solutions/findmediansortedarrays.go)     |
|  7   | [Reverse Integer](https://leetcode.com/problems/reverse-integer/)                                                                                     | [`reverse.go`](./internal/solutions/reverse.go)                                   |
|  9   | [Palindrome Number](https://leetcode.com/problems/palindrome-number/)                                                                                 | [`ispalindrome.go`](./internal/solutions/ispalindrome.go)                         |
|  53  | [Maximum Subarray](https://leetcode.com/problems/maximum-subarray/)                                                                                   | [`maxsubarray.go`](./internal/solutions/maxsubarray.go)                           |
|  58  | [Length of Last Word](https://leetcode.com/problems/length-of-last-word/)                                                                             | [`lengthoflastword.go`](./internal/solutions/lengthoflastword.go)                 |
|  70  | [Climbing Stairs](https://leetcode.com/problems/climbing-stairs/)                                                                                     | [`climbstairs.go`](./internal/solutions/climbstairs.go)                           |
| 118  | [Pascal's Triangle](https://leetcode.com/problems/pascals-triangle/)                                                                                  | [`generate.go`](./internal/solutions/generate.go)                                 |
| 119  | [Pascal's Triangle II](https://leetcode.com/problems/pascals-triangle-ii/)                                                                            | [`getrow.go`](./internal/solutions/getrow.go)                                     |
| 198  | [House Robber](https://leetcode.com/problems/house-robber/)                                                                                           | [`rob.go`](./internal/solutions/rob.go)                                           |
| 231  | [Power of Two](https://leetcode.com/problems/power-of-two/)                                                                                           | [`ispoweroftwo.go`](./internal/solutions/ispoweroftwo.go)                         |
| 342  | [Power of Four](https://leetcode.com/problems/power-of-four/)                                                                                         | [`ispoweroffour.go`](./internal/solutions/ispoweroffour.go)                       |
| 458  | [Poor Pigs](https://leetcode.com/problems/poor-pigs/)                                                                                                 | [`poorpigs.go`](./internal/solutions/poorpigs.go)                                 |
| 501  | [Find Mode in Binary Search Tree](https://leetcode.com/problems/find-mode-in-binary-search-tree/)                                                     | [`findmode.go`](./internal/solutions/findmode.go)                                 |
| 779  | [K-th Symbol in Grammar](https://leetcode.com/problems/k-th-symbol-in-grammar/)                                                                       | [`kthgrammar.go`](./internal/solutions/kthgrammar.go)                             |
| 844  | [Backspace String Compare](https://leetcode.com/problems/backspace-string-compare/)                                                                   | [`backspacecompare.go`](./internal/solutions/backspacecompare.go)                 |
| 1018 | [Binary Prefix Divisible By 5](https://leetcode.com/problems/binary-prefix-divisible-by-5/)                                                           | [`prefixesdivby5.go`](./internal/solutions/prefixesdivby5.go)                     |
| 1085 | [Sum of Digits in the Minimum Number](https://leetcode.com/problems/sum-of-digits-in-the-minimum-number/)                                             | [`sumofdigits.go`](./internal/solutions/sumofdigits.go)                           |
| 1269 | [Number of Ways to Stay in the Same Place After Some Steps](https://leetcode.com/problems/number-of-ways-to-stay-in-the-same-place-after-some-steps/) | [`numways.go`](./internal/solutions/numways.go)                                   |
| 1356 | [Sort Integers by The Number of 1 Bits](https://leetcode.com/problems/sort-integers-by-the-number-of-1-bits/)                                         | [`sortbybits.go`](./internal/solutions/sortbybits.go)                             |
| 1361 | [Validate Binary Tree Nodes](https://leetcode.com/problems/validate-binary-tree-nodes/)                                                               | [`validatebinarytreenodes.go`](./internal/solutions/validatebinarytreenodes.go)   |
| 1441 | [Build an Array With Stack Operations](https://leetcode.com/problems/build-an-array-with-stack-operations/)                                           | [`buildarray.go`](./internal/solutions/buildarray.go)                             |
| 1503 | [Last Moment Before All Ants Fall Out of a Plank](https://leetcode.com/problems/last-moment-before-all-ants-fall-out-of-a-plank/)                     | [`getlastmoment.go`](./internal/solutions/getlastmoment.go)                       |
| 1535 | [Find the Winner of an Array Game](https://leetcode.com/problems/find-the-winner-of-an-array-game/)                                                   | [`getwinner.go`](./internal/solutions/getwinner.go)                               |
| 1720 | [Decode XORed Array](https://leetcode.com/problems/decode-xored-array/)                                                                               | [`decode.go`](./internal/solutions/decode.go)                                     |
| 1743 | [Restore the Array From Adjacent Pairs](https://leetcode.com/problems/restore-the-array-from-adjacent-pairs/)                                         | [`restorearray.go`](./internal/solutions/restorearray.go)                         |
| 1759 | [Count Number of Homogenous Substrings](https://leetcode.com/problems/count-number-of-homogenous-substrings/)                                         | [`counthomogenous.go`](./internal/solutions/counthomogenous.go)                   |
| 1838 | [Frequency of the Most Frequent Element](https://leetcode.com/problems/frequency-of-the-most-frequent-element/)                                       | [`maxfrequency.go`](./internal/solutions/maxfrequency.go)                         |
| 1845 | [Seat Reservation Manager](https://leetcode.com/problems/seat-reservation-manager/)                                                                   | [`seatmanager.go`](./internal/solutions/seatmanager.go)                           |
| 1887 | [Reduction Operations to Make the Array Elements Equal](https://leetcode.com/problems/reduction-operations-to-make-the-array-elements-equal/)         | [`reductionoperations.go`](./internal/solutions/reductionoperations.go)           |
| 1921 | [Eliminate Maximum Number of Monsters](https://leetcode.com/problems/eliminate-maximum-number-of-monsters/)                                           | [`eliminatemaximum.go`](./internal/solutions/eliminatemaximum.go)                 |
| 2265 | [Count Nodes Equal to Average of Subtree](https://leetcode.com/problems/count-nodes-equal-to-average-of-subtree/)                                     | [`averageofsubtree.go`](./internal/solutions/averageofsubtree.go)                 |
| 2433 | [Find The Original Array of Prefix Xor](https://leetcode.com/problems/find-the-original-array-of-prefix-xor/)                                         | [`findarray.go`](./internal/solutions/findarray.go)                               |
| 2785 | [Sort Vowels in a String](https://leetcode.com/problems/sort-vowels-in-a-string/)                                                                     | [`sortvowels.go`](./internal/solutions/sortvowels.go)                             |
| 2849 | [Determine if a Cell Is Reachable at a Given Time](https://leetcode.com/problems/determine-if-a-cell-is-reachable-at-a-given-time/)                   | [`isreachableattime.go`](./internal/solutions/isreachableattime.go)               |

5. Longest Palindromic Substring
8. String to Integer (atoi)
11. Container With Most Water
13. Roman to Integer
14. Longest Common Prefix
19. Remove Nth Node From End of List
20. Valid Parentheses
21. Merge Two Sorted Lists
24. Swap Nodes in Pairs
26. Remove Duplicates from Sorted Array
27. Remove Element
28. Find the Index of the First Occurrence in a String
34. Find First and Last Position of Element in Sorted Array
36. Valid Sudoku
48. Rotate Image
49. Group Anagrams
54. Spiral Matrix
66. Plus One
67. Add Binary
71. Simplify Path
82. Remove Duplicates from Sorted List II
83. Remove Duplicates from Sorted List
88. Merge Sorted Array
92. Reverse Linked List II
94. Binary Tree Inorder Traversal
98. Validate Binary Search Tree
100. Same Tree
101. Symmetric Tree
102. Binary Tree Level Order Traversal
103. Binary Tree Zigzag Level Order Traversal
104. Maximum Depth of Binary Tree
108. Convert Sorted Array to Binary Search Tree
111. Minimum Depth of Binary Tree
112. Path Sum
121. Best Time to Buy and Sell Stock
122. Best Time to Buy and Sell Stock II
125. Valid Palindrome
136. Single Number
141. Linked List Cycle
146. LRU Cache
151. Reverse Words in a String
155. Min Stack
162. Find Peak Element
167. Two Sum II - Input Array Is Sorted
189. Rotate Array
190. Reverse Bits
191. Number of 1 Bits
199. Binary Tree Right Side View
203. Remove Linked List Elements
204. Count Primes
205. Isomorphic Strings
206. Reverse Linked List
209. Minimum Size Subarray Sum
217. Contains Duplicate
225. Implement Stack using Queues
228. Summary Ranges
229. Majority Element II
232. Implement Queue using Stacks
234. Palindrome Linked List
236. Lowest Common Ancestor of a Binary Tree
237. Delete Node in a Linked List
238. Product of Array Except Self
239. Sliding Window Maximum
242. Valid Anagram
268. Missing Number
270. Closest Binary Search Tree Value
278. First Bad Version
283. Move Zeroes
287. Find the Duplicate Number
290. Word Pattern
303. Range Sum Query - Immutable
326. Power of Three
328. Odd Even Linked List
334. Increasing Triplet Subsequence
338. Counting Bits
341. Flatten Nested List Iterator
343. Integer Break
344. Reverse String
345. Reverse Vowels of a String
350. Intersection of Two Arrays II
374. Guess Number Higher or Lower
383. Ransom Note
384. Shuffle an Array
387. First Unique Character in a String
389. Find the Difference
392. Is Subsequence
394. Decode String
412. Fizz Buzz
437. Path Sum III
438. Find All Anagrams in a String
443. String Compression
450. Delete Node in a BST
451. Sort Characters By Frequency
461. Hamming Distance
485. Max Consecutive Ones
496. Next Greater Element I
498. Diagonal Traverse
515. Find Largest Value in Each Tree Row
530. Minimum Absolute Difference in BST
543. Diameter of Binary Tree
557. Reverse Words in a String III
560. Subarray Sum Equals K
561. Array Partition
567. Permutation in String
605. Can Place Flowers
643. Maximum Average Subarray I
649. Dota2 Senate
700. Search in a Binary Search Tree
701. Insert into a Binary Search Tree
706. Design HashMap
707. Design Linked List
713. Subarray Product Less Than K
724. Find Pivot Index
735. Asteroid Collision
739. Daily Temperatures
744. Find Smallest Letter Greater Than Target
746. Min Cost Climbing Stairs
747. Largest Number At Least Twice of Others
771. Jewels and Stones
783. Minimum Distance Between BST Nodes
791. Custom Sort String
815. Bus Routes
823. Binary Trees With Factors
852. Peak Index in a Mountain Array
872. Leaf-Similar Trees
876. Middle of the Linked List
901. Online Stock Span
917. Reverse Only Letters
930. Binary Subarrays With Sum
933. Number of Recent Calls
938. Range Sum of BST
946. Validate Stack Sequences
977. Squares of a Sorted Array
1004. Max Consecutive Ones III
1026. Maximum Difference Between Node and Ancestor
1027. Longest Arithmetic Subsequence
1047. Remove All Adjacent Duplicates In String
1071. Greatest Common Divisor of Strings
1095. Find in Mountain Array
1114. Print in Order
1133. Largest Unique Number
1137. N-th Tribonacci Number
1146. Snapshot Array
1161. Maximum Level Sum of a Binary Tree
1189. Maximum Number of Balloons
1207. Unique Number of Occurrences
1208. Get Equal Substrings Within Budget
1220. Count Vowels Permutation
1232. Check If It Is a Straight Line
1248. Count Number of Nice Subarrays
1290. Convert Binary Number in a Linked List to Integer
1302. Deepest Leaves Sum
1342. Number of Steps to Reduce a Number to Zero
1351. Count Negative Numbers in a Sorted Matrix
1372. Longest ZigZag Path in a Binary Tree
1394. Find Lucky Integer in an Array
1413. Minimum Value to Get Positive Step by Step Sum
1420. Build Array Where You Can Find The Maximum Exactly K Comparisons
1425. Constrained Subsequence Sum
1426. Counting Elements
1431. Kids With the Greatest Number of Candies
1436. Destination City
1438. Longest Continuous Subarray With Absolute Diff Less Than or Equal to Limit
1448. Count Good Nodes in Binary Tree
1456. Maximum Number of Vowels in a Substring of Given Length
1458. Max Dot Product of Two Subsequences
1480. Running Sum of 1d Array
1493. Longest Subarray of 1's After Deleting One Element
1496. Path Crossing
1502. Can Make Arithmetic Progression From Sequence
1512. Number of Good Pairs
1544. Make The String Great
1657. Determine if Two Strings Are Close
1672. Richest Customer Wealth
1679. Max Number of K-Sum Pairs
1695. Maximum Erasure Value
1721. Swapping Nodes in a Linked List
1732. Find the Highest Altitude
1748. Sum of Unique Elements
1768. Merge Strings Alternately
1793. Maximum Score of a Good Subarray
1832. Check if the Sentence Is Pangram
1941. Check if All Characters Have Equal Number of Occurrences
2000. Reverse Prefix of Word
2009. Minimum Number of Operations to Make Array Continuous
2038. Remove Colored Pieces if Both Neighbors are the Same Color
2050. Parallel Courses III
2074. Reverse Nodes in Even Length Groups
2090. K Radius Subarray Averages
2095. Delete the Middle Node of a Linked List
2130. Maximum Twin Sum of a Linked List
2215. Find the Difference of Two Arrays
2225. Find Players With Zero or One Losses
2235. Add Two Integers
2248. Intersection of Multiple Arrays
2251. Number of Flowers in Full Bloom
2260. Minimum Consecutive Cards to Pick Up
2270. Number of Ways to Split Array
2300. Successful Pairs of Spells and Potions
2342. Max Sum of a Pair With Equal Sum of Digits
2351. First Letter to Appear Twice
2352. Equal Row and Column Pairs
2390. Removing Stars From a String
2434. Using a Robot to Print the Lexicographically Smallest String
2462. Total Cost to Hire K Workers
2642. Design Graph With Shortest Path Calculator
2742. Painting the Walls
