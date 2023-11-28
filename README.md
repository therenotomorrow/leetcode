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

|  #   | Problem                                                                                                                                               | Solution                                                                                                                    |
|:----:|:------------------------------------------------------------------------------------------------------------------------------------------------------|:----------------------------------------------------------------------------------------------------------------------------|
|  1   | [Two Sum](https://leetcode.com/problems/two-sum/)                                                                                                     | [`twosum.go`](./internal/solutions/twosum.go)                                                                               |
|  2   | [Add Two Numbers](https://leetcode.com/problems/add-two-numbers/)                                                                                     | [`addtwonumbers.go`](./internal/solutions/addtwonumbers.go)                                                                 |
|  3   | [Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/)                       | [`lengthoflongestsubstring.go`](./internal/solutions/lengthoflongestsubstring.go)                                           |
|  4   | [Median of Two Sorted Arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/)                                                             | [`findmediansortedarrays.go`](./internal/solutions/findmediansortedarrays.go)                                               |
|  7   | [Reverse Integer](https://leetcode.com/problems/reverse-integer/)                                                                                     | [`reverse.go`](./internal/solutions/reverse.go)                                                                             |
|  9   | [Palindrome Number](https://leetcode.com/problems/palindrome-number/)                                                                                 | [`ispalindrome.go`](./internal/solutions/ispalindrome.go)                                                                   |
|  53  | [Maximum Subarray](https://leetcode.com/problems/maximum-subarray/)                                                                                   | [`maxsubarray.go`](./internal/solutions/maxsubarray.go)                                                                     |
|  58  | [Length of Last Word](https://leetcode.com/problems/length-of-last-word/)                                                                             | [`lengthoflastword.go`](./internal/solutions/lengthoflastword.go)                                                           |
|  70  | [Climbing Stairs](https://leetcode.com/problems/climbing-stairs/)                                                                                     | [`climbstairs.go`](./internal/solutions/climbstairs.go)                                                                     |
| 118  | [Pascal's Triangle](https://leetcode.com/problems/pascals-triangle/)                                                                                  | [`generate.go`](./internal/solutions/generate.go)                                                                           |
| 119  | [Pascal's Triangle II](https://leetcode.com/problems/pascals-triangle-ii/)                                                                            | [`getrow.go`](./internal/solutions/getrow.go)                                                                               |
| 198  | [House Robber](https://leetcode.com/problems/house-robber/)                                                                                           | [`rob.go`](./internal/solutions/rob.go)                                                                                     |
| 231  | [Power of Two](https://leetcode.com/problems/power-of-two/)                                                                                           | [`ispoweroftwo.go`](./internal/solutions/ispoweroftwo.go)                                                                   |
| 342  | [Power of Four](https://leetcode.com/problems/power-of-four/)                                                                                         | [`ispoweroffour.go`](./internal/solutions/ispoweroffour.go)                                                                 |
| 346  | [Moving Average from Data Stream](https://leetcode.com/problems/moving-average-from-data-stream/)                                                     | [`movingaverage.go`](./internal/solutions/movingaverage.go)                                                                 |
| 458  | [Poor Pigs](https://leetcode.com/problems/poor-pigs/)                                                                                                 | [`poorpigs.go`](./internal/solutions/poorpigs.go)                                                                           |
| 501  | [Find Mode in Binary Search Tree](https://leetcode.com/problems/find-mode-in-binary-search-tree/)                                                     | [`findmode.go`](./internal/solutions/findmode.go)                                                                           |
| 573  | [Squirrel Simulation](https://leetcode.com/problems/squirrel-simulation/)                                                                             | [`mindistance.go`](./internal/solutions/mindistance.go)                                                                     |
| 624  | [Maximum Distance in Arrays](https://leetcode.com/problems/maximum-distance-in-arrays/)                                                               | [`maxdistance.go`](./internal/solutions/maxdistance.go)                                                                     |
| 779  | [K-th Symbol in Grammar](https://leetcode.com/problems/k-th-symbol-in-grammar/)                                                                       | [`kthgrammar.go`](./internal/solutions/kthgrammar.go)                                                                       |
| 844  | [Backspace String Compare](https://leetcode.com/problems/backspace-string-compare/)                                                                   | [`backspacecompare.go`](./internal/solutions/backspacecompare.go)                                                           |
| 935  | [Knight Dialer](https://leetcode.com/problems/knight-dialer/)                                                                                         | [`knightdialer.go`](./internal/solutions/knightdialer.go)                                                                   |
| 1018 | [Binary Prefix Divisible By 5](https://leetcode.com/problems/binary-prefix-divisible-by-5/)                                                           | [`prefixesdivby5.go`](./internal/solutions/prefixesdivby5.go)                                                               |
| 1057 | [Campus Bikes](https://leetcode.com/problems/campus-bikes/)                                                                                           | [`assignbikes.go`](./internal/solutions/assignbikes.go)                                                                     |
| 1085 | [Sum of Digits in the Minimum Number](https://leetcode.com/problems/sum-of-digits-in-the-minimum-number/)                                             | [`sumofdigits.go`](./internal/solutions/sumofdigits.go)                                                                     |
| 1269 | [Number of Ways to Stay in the Same Place After Some Steps](https://leetcode.com/problems/number-of-ways-to-stay-in-the-same-place-after-some-steps/) | [`numways.go`](./internal/solutions/numways.go)                                                                             |
| 1356 | [Sort Integers by The Number of 1 Bits](https://leetcode.com/problems/sort-integers-by-the-number-of-1-bits/)                                         | [`sortbybits.go`](./internal/solutions/sortbybits.go)                                                                       |
| 1361 | [Validate Binary Tree Nodes](https://leetcode.com/problems/validate-binary-tree-nodes/)                                                               | [`validatebinarytreenodes.go`](./internal/solutions/validatebinarytreenodes.go)                                             |
| 1424 | [Diagonal Traverse II](https://leetcode.com/problems/diagonal-traverse-ii/)                                                                           | [`finddiagonalorder.go`](./internal/solutions/finddiagonalorder.go)                                                         |
| 1441 | [Build an Array With Stack Operations](https://leetcode.com/problems/build-an-array-with-stack-operations/)                                           | [`buildarray.go`](./internal/solutions/buildarray.go)                                                                       |
| 1503 | [Last Moment Before All Ants Fall Out of a Plank](https://leetcode.com/problems/last-moment-before-all-ants-fall-out-of-a-plank/)                     | [`getlastmoment.go`](./internal/solutions/getlastmoment.go)                                                                 |
| 1535 | [Find the Winner of an Array Game](https://leetcode.com/problems/find-the-winner-of-an-array-game/)                                                   | [`getwinner.go`](./internal/solutions/getwinner.go)                                                                         |
| 1561 | [Maximum Number of Coins You Can Get](https://leetcode.com/problems/maximum-number-of-coins-you-can-get/)                                             | [`maxcoins.go`](./internal/solutions/maxcoins.go)                                                                           |
| 1630 | [Arithmetic Subarrays](https://leetcode.com/problems/arithmetic-subarrays/)                                                                           | [`checkarithmeticsubarrays.go`](./internal/solutions/checkarithmeticsubarrays.go)                                           |
| 1685 | [Sum of Absolute Differences in a Sorted Array](https://leetcode.com/problems/sum-of-absolute-differences-in-a-sorted-array/)                         | [`getsumabsolutedifferences.go`](./internal/solutions/getsumabsolutedifferences.go)                                         |
| 1720 | [Decode XORed Array](https://leetcode.com/problems/decode-xored-array/)                                                                               | [`decode.go`](./internal/solutions/decode.go)                                                                               |
| 1743 | [Restore the Array From Adjacent Pairs](https://leetcode.com/problems/restore-the-array-from-adjacent-pairs/)                                         | [`restorearray.go`](./internal/solutions/restorearray.go)                                                                   |
| 1759 | [Count Number of Homogenous Substrings](https://leetcode.com/problems/count-number-of-homogenous-substrings/)                                         | [`counthomogenous.go`](./internal/solutions/counthomogenous.go)                                                             |
| 1814 | [Count Nice Pairs in an Array](https://leetcode.com/problems/count-nice-pairs-in-an-array/)                                                           | [`countnicepairs.go`](./internal/solutions/countnicepairs.go)                                                               |
| 1838 | [Frequency of the Most Frequent Element](https://leetcode.com/problems/frequency-of-the-most-frequent-element/)                                       | [`maxfrequency.go`](./internal/solutions/maxfrequency.go)                                                                   |
| 1845 | [Seat Reservation Manager](https://leetcode.com/problems/seat-reservation-manager/)                                                                   | [`seatmanager.go`](./internal/solutions/seatmanager.go)                                                                     |
| 1846 | [Maximum Element After Decreasing and Rearranging](https://leetcode.com/problems/maximum-element-after-decreasing-and-rearranging/)                   | [`maximumelementafterdecrementingandrearranging.go`](./internal/solutions/maximumelementafterdecrementingandrearranging.go) |
| 1877 | [Minimize Maximum Pair Sum in Array](https://leetcode.com/problems/minimize-maximum-pair-sum-in-array/)                                               | [`minpairsum.go`](./internal/solutions/minpairsum.go)                                                                       |
| 1887 | [Reduction Operations to Make the Array Elements Equal](https://leetcode.com/problems/reduction-operations-to-make-the-array-elements-equal/)         | [`reductionoperations.go`](./internal/solutions/reductionoperations.go)                                                     |
| 1921 | [Eliminate Maximum Number of Monsters](https://leetcode.com/problems/eliminate-maximum-number-of-monsters/)                                           | [`eliminatemaximum.go`](./internal/solutions/eliminatemaximum.go)                                                           |
| 1930 | [Unique Length-3 Palindromic Subsequences](https://leetcode.com/problems/unique-length-3-palindromic-subsequences/)                                   | [`countpalindromicsubsequence.go`](./internal/solutions/countpalindromicsubsequence.go)                                     |
| 1980 | [Find Unique Binary String](https://leetcode.com/problems/find-unique-binary-string/)                                                                 | [`finddifferentbinarystring.go`](./internal/solutions/finddifferentbinarystring.go)                                         |
| 2265 | [Count Nodes Equal to Average of Subtree](https://leetcode.com/problems/count-nodes-equal-to-average-of-subtree/)                                     | [`averageofsubtree.go`](./internal/solutions/averageofsubtree.go)                                                           |
| 2391 | [Minimum Amount of Time to Collect Garbage](https://leetcode.com/problems/minimum-amount-of-time-to-collect-garbage/)                                 | [`garbagecollection.go`](./internal/solutions/garbagecollection.go)                                                         |
| 2433 | [Find The Original Array of Prefix Xor](https://leetcode.com/problems/find-the-original-array-of-prefix-xor/)                                         | [`findarray.go`](./internal/solutions/findarray.go)                                                                         |
| 2785 | [Sort Vowels in a String](https://leetcode.com/problems/sort-vowels-in-a-string/)                                                                     | [`sortvowels.go`](./internal/solutions/sortvowels.go)                                                                       |
| 2849 | [Determine if a Cell Is Reachable at a Given Time](https://leetcode.com/problems/determine-if-a-cell-is-reachable-at-a-given-time/)                   | [`isreachableattime.go`](./internal/solutions/isreachableattime.go)                                                         |

In revision
-----------

1. Longest Palindromic Substring
2. String to Integer (atoi)
3. Container With Most Water
4. Roman to Integer
5. Longest Common Prefix
6. Remove Nth Node From End of List
7. Valid Parentheses
8. Merge Two Sorted Lists
9. Swap Nodes in Pairs
10. Remove Duplicates from Sorted Array
11. Remove Element
12. Find the Index of the First Occurrence in a String
13. Find First and Last Position of Element in Sorted Array
14. Valid Sudoku
15. Rotate Image
16. Group Anagrams
17. Spiral Matrix
18. Plus One
19. Add Binary
20. Simplify Path
21. Remove Duplicates from Sorted List II
22. Remove Duplicates from Sorted List
23. Merge Sorted Array
24. Reverse Linked List II
25. Binary Tree Inorder Traversal
26. Validate Binary Search Tree
27. Same Tree
28. Symmetric Tree
29. Binary Tree Level Order Traversal
30. Binary Tree Zigzag Level Order Traversal
31. Maximum Depth of Binary Tree
32. Convert Sorted Array to Binary Search Tree
33. Minimum Depth of Binary Tree
34. Path Sum
35. Best Time to Buy and Sell Stock
36. Best Time to Buy and Sell Stock II
37. Valid Palindrome
38. Single Number
39. Linked List Cycle
40. LRU Cache
41. Reverse Words in a String
42. Min Stack
43. Find Peak Element
44. Two Sum II - Input Array Is Sorted
45. Rotate Array
46. Reverse Bits
47. Number of 1 Bits
48. Binary Tree Right Side View
49. Remove Linked List Elements
50. Count Primes
51. Isomorphic Strings
52. Reverse Linked List
53. Minimum Size Subarray Sum
54. Contains Duplicate
55. Implement Stack using Queues
56. Summary Ranges
57. Majority Element II
58. Implement Queue using Stacks
59. Palindrome Linked List
60. Lowest Common Ancestor of a Binary Tree
61. Delete Node in a Linked List
62. Product of Array Except Self
63. Sliding Window Maximum
64. Valid Anagram
65. Missing Number
66. Closest Binary Search Tree Value
67. First Bad Version
68. Move Zeroes
69. Find the Duplicate Number
70. Word Pattern
71. Range Sum Query - Immutable
72. Power of Three
73. Odd Even Linked List
74. Increasing Triplet Subsequence
75. Counting Bits
76. Flatten Nested List Iterator
77. Integer Break
78. Reverse String
79. Reverse Vowels of a String
80. Intersection of Two Arrays II
81. Guess Number Higher or Lower
82. Ransom Note
83. Shuffle an Array
84. First Unique Character in a String
85. Find the Difference
86. Is Subsequence
87. Decode String
88. Fizz Buzz
89. Path Sum III
90. Find All Anagrams in a String
91. String Compression
92. Delete Node in a BST
93. Sort Characters By Frequency
94. Hamming Distance
95. Max Consecutive Ones
96. Next Greater Element I
97. Diagonal Traverse
98. Find Largest Value in Each Tree Row
99. Minimum Absolute Difference in BST
100. Diameter of Binary Tree
101. Reverse Words in a String III
102. Subarray Sum Equals K
103. Array Partition
104. Permutation in String
105. Can Place Flowers
106. Maximum Average Subarray I
107. Dota2 Senate
108. Search in a Binary Search Tree
109. Insert into a Binary Search Tree
110. Design HashMap
111. Design Linked List
112. Subarray Product Less Than K
113. Find Pivot Index
114. Asteroid Collision
115. Daily Temperatures
116. Find Smallest Letter Greater Than Target
117. Min Cost Climbing Stairs
118. Largest Number At Least Twice of Others
119. Jewels and Stones
120. Minimum Distance Between BST Nodes
121. Custom Sort String
122. Bus Routes
123. Binary Trees With Factors
124. Peak Index in a Mountain Array
125. Leaf-Similar Trees
126. Middle of the Linked List
127. Online Stock Span
128. Reverse Only Letters
129. Binary Subarrays With Sum
130. Number of Recent Calls
131. Range Sum of BST
132. Validate Stack Sequences
133. Squares of a Sorted Array
134. Max Consecutive Ones III
135. Maximum Difference Between Node and Ancestor
136. Longest Arithmetic Subsequence
137. Remove All Adjacent Duplicates In String
138. Greatest Common Divisor of Strings
139. Find in Mountain Array
140. Print in Order
141. Largest Unique Number
142. N-th Tribonacci Number
143. Snapshot Array
144. Maximum Level Sum of a Binary Tree
145. Maximum Number of Balloons
146. Unique Number of Occurrences
147. Get Equal Substrings Within Budget
148. Count Vowels Permutation
149. Check If It Is a Straight Line
150. Count Number of Nice Subarrays
151. Convert Binary Number in a Linked List to Integer
152. Deepest Leaves Sum
153. Number of Steps to Reduce a Number to Zero
154. Count Negative Numbers in a Sorted Matrix
155. Longest ZigZag Path in a Binary Tree
156. Find Lucky Integer in an Array
157. Minimum Value to Get Positive Step by Step Sum
158. Build Array Where You Can Find The Maximum Exactly K Comparisons
159. Constrained Subsequence Sum
160. Counting Elements
161. Kids With the Greatest Number of Candies
162. Destination City
163. Longest Continuous Subarray With Absolute Diff Less Than or Equal to Limit
164. Count Good Nodes in Binary Tree
165. Maximum Number of Vowels in a Substring of Given Length
166. Max Dot Product of Two Subsequences
167. Running Sum of 1d Array
168. Longest Subarray of 1's After Deleting One Element
169. Path Crossing
170. Can Make Arithmetic Progression From Sequence
171. Number of Good Pairs
172. Make The String Great
173. Determine if Two Strings Are Close
174. Richest Customer Wealth
175. Max Number of K-Sum Pairs
176. Maximum Erasure Value
177. Swapping Nodes in a Linked List
178. Find the Highest Altitude
179. Sum of Unique Elements
180. Merge Strings Alternately
181. Maximum Score of a Good Subarray
182. Check if the Sentence Is Pangram
183. Check if All Characters Have Equal Number of Occurrences
184. Reverse Prefix of Word
185. Minimum Number of Operations to Make Array Continuous
186. Remove Colored Pieces if Both Neighbors are the Same Color
187. Parallel Courses III
188. Reverse Nodes in Even Length Groups
189. K Radius Subarray Averages
190. Delete the Middle Node of a Linked List
191. Maximum Twin Sum of a Linked List
192. Find the Difference of Two Arrays
193. Find Players With Zero or One Losses
194. Add Two Integers
195. Intersection of Multiple Arrays
196. Number of Flowers in Full Bloom
197. Minimum Consecutive Cards to Pick Up
198. Number of Ways to Split Array
199. Successful Pairs of Spells and Potions
200. Max Sum of a Pair With Equal Sum of Digits
201. First Letter to Appear Twice
202. Equal Row and Column Pairs
203. Removing Stars From a String
204. Using a Robot to Print the Lexicographically Smallest String
205. Total Cost to Hire K Workers
206. Design Graph With Shortest Path Calculator
207. Painting the Walls
208. Largest Submatrix With Rearrangements
209. Number of Ways to Divide a Long Corridor
