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
| 191  | [Number of 1 Bits](https://leetcode.com/problems/number-of-1-bits/)                                                                                   | [`hammingweight.go`](./internal/solutions/hammingweight.go)                                                                 |
| 198  | [House Robber](https://leetcode.com/problems/house-robber/)                                                                                           | [`rob.go`](./internal/solutions/rob.go)                                                                                     |
| 231  | [Power of Two](https://leetcode.com/problems/power-of-two/)                                                                                           | [`ispoweroftwo.go`](./internal/solutions/ispoweroftwo.go)                                                                   |
| 342  | [Power of Four](https://leetcode.com/problems/power-of-four/)                                                                                         | [`ispoweroffour.go`](./internal/solutions/ispoweroffour.go)                                                                 |
| 346  | [Moving Average from Data Stream](https://leetcode.com/problems/moving-average-from-data-stream/)                                                     | [`movingaverage.go`](./internal/solutions/movingaverage.go)                                                                 |
| 458  | [Poor Pigs](https://leetcode.com/problems/poor-pigs/)                                                                                                 | [`poorpigs.go`](./internal/solutions/poorpigs.go)                                                                           |
| 485  | [Max Consecutive Ones](https://leetcode.com/problems/max-consecutive-ones/)                                                                           | [`findmaxconsecutiveones_1.go`](./internal/solutions/findmaxconsecutiveones_1.go)                                           |
| 487  | [Max Consecutive Ones II](https://leetcode.com/problems/max-consecutive-ones-ii/)                                                                     | [`findmaxconsecutiveones_2.go`](./internal/solutions/findmaxconsecutiveones_2.go)                                           |
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
47. Binary Tree Right Side View
48. Remove Linked List Elements
49. Count Primes
50. Isomorphic Strings
51. Reverse Linked List
52. Minimum Size Subarray Sum
53. Contains Duplicate
54. Implement Stack using Queues
55. Summary Ranges
56. Majority Element II
57. Implement Queue using Stacks
58. Palindrome Linked List
59. Lowest Common Ancestor of a Binary Tree
60. Delete Node in a Linked List
61. Product of Array Except Self
62. Sliding Window Maximum
63. Valid Anagram
64. Missing Number
65. Closest Binary Search Tree Value
66. First Bad Version
67. Move Zeroes
68. Find the Duplicate Number
69. Word Pattern
70. Range Sum Query - Immutable
71. Power of Three
72. Odd Even Linked List
73. Increasing Triplet Subsequence
74. Counting Bits
75. Flatten Nested List Iterator
76. Integer Break
77. Reverse String
78. Reverse Vowels of a String
79. Intersection of Two Arrays II
80. Guess Number Higher or Lower
81. Ransom Note
82. Shuffle an Array
83. First Unique Character in a String
84. Find the Difference
85. Is Subsequence
86. Decode String
87. Fizz Buzz
88. Path Sum III
89. Find All Anagrams in a String
90. String Compression
91. Delete Node in a BST
92. Sort Characters By Frequency
93. Hamming Distance
94. Next Greater Element I
95. Diagonal Traverse
96. Find Largest Value in Each Tree Row
97. Minimum Absolute Difference in BST
98. Diameter of Binary Tree
99. Reverse Words in a String III
100. Subarray Sum Equals K
101. Array Partition
102. Permutation in String
103. Can Place Flowers
104. Maximum Average Subarray I
105. Dota2 Senate
106. Search in a Binary Search Tree
107. Insert into a Binary Search Tree
108. Design HashMap
109. Design Linked List
110. Subarray Product Less Than K
111. Find Pivot Index
112. Asteroid Collision
113. Daily Temperatures
114. Find Smallest Letter Greater Than Target
115. Min Cost Climbing Stairs
116. Largest Number At Least Twice of Others
117. Jewels and Stones
118. Minimum Distance Between BST Nodes
119. Custom Sort String
120. Bus Routes
121. Binary Trees With Factors
122. Peak Index in a Mountain Array
123. Leaf-Similar Trees
124. Middle of the Linked List
125. Online Stock Span
126. Reverse Only Letters
127. Binary Subarrays With Sum
128. Number of Recent Calls
129. Range Sum of BST
130. Validate Stack Sequences
131. Squares of a Sorted Array
132. Max Consecutive Ones III
133. Maximum Difference Between Node and Ancestor
134. Longest Arithmetic Subsequence
135. Remove All Adjacent Duplicates In String
136. Greatest Common Divisor of Strings
137. Find in Mountain Array
138. Print in Order
139. Largest Unique Number
140. N-th Tribonacci Number
141. Snapshot Array
142. Maximum Level Sum of a Binary Tree
143. Maximum Number of Balloons
144. Unique Number of Occurrences
145. Get Equal Substrings Within Budget
146. Count Vowels Permutation
147. Check If It Is a Straight Line
148. Count Number of Nice Subarrays
149. Convert Binary Number in a Linked List to Integer
150. Deepest Leaves Sum
151. Number of Steps to Reduce a Number to Zero
152. Count Negative Numbers in a Sorted Matrix
153. Longest ZigZag Path in a Binary Tree
154. Find Lucky Integer in an Array
155. Minimum Value to Get Positive Step by Step Sum
156. Build Array Where You Can Find The Maximum Exactly K Comparisons
157. Constrained Subsequence Sum
158. Counting Elements
159. Kids With the Greatest Number of Candies
160. Destination City
161. Longest Continuous Subarray With Absolute Diff Less Than or Equal to Limit
162. Count Good Nodes in Binary Tree
163. Maximum Number of Vowels in a Substring of Given Length
164. Max Dot Product of Two Subsequences
165. Running Sum of 1d Array
166. Longest Subarray of 1's After Deleting One Element
167. Path Crossing
168. Can Make Arithmetic Progression From Sequence
169. Number of Good Pairs
170. Make The String Great
171. Determine if Two Strings Are Close
172. Richest Customer Wealth
173. Max Number of K-Sum Pairs
174. Maximum Erasure Value
175. Swapping Nodes in a Linked List
176. Find the Highest Altitude
177. Sum of Unique Elements
178. Merge Strings Alternately
179. Maximum Score of a Good Subarray
180. Check if the Sentence Is Pangram
181. Check if All Characters Have Equal Number of Occurrences
182. Reverse Prefix of Word
183. Minimum Number of Operations to Make Array Continuous
184. Remove Colored Pieces if Both Neighbors are the Same Color
185. Parallel Courses III
186. Reverse Nodes in Even Length Groups
187. K Radius Subarray Averages
188. Delete the Middle Node of a Linked List
189. Maximum Twin Sum of a Linked List
190. Find the Difference of Two Arrays
191. Find Players With Zero or One Losses
192. Add Two Integers
193. Intersection of Multiple Arrays
194. Number of Flowers in Full Bloom
195. Minimum Consecutive Cards to Pick Up
196. Number of Ways to Split Array
197. Successful Pairs of Spells and Potions
198. Max Sum of a Pair With Equal Sum of Digits
199. First Letter to Appear Twice
200. Equal Row and Column Pairs
201. Removing Stars From a String
202. Using a Robot to Print the Lexicographically Smallest String
203. Total Cost to Hire K Workers
204. Design Graph With Shortest Path Calculator
205. Painting the Walls
206. Largest Submatrix With Rearrangements
207. Number of Ways to Divide a Long Corridor
208. Minimum One Bit Operations to Make Integers Zero
