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
|  35  | [Search Insert Position](https://leetcode.com/problems/search-insert-position/)                                                                       | [`searchinsert.go`](./internal/solutions/searchinsert.go)                                                                   |
|  53  | [Maximum Subarray](https://leetcode.com/problems/maximum-subarray/)                                                                                   | [`maxsubarray.go`](./internal/solutions/maxsubarray.go)                                                                     |
|  58  | [Length of Last Word](https://leetcode.com/problems/length-of-last-word/)                                                                             | [`lengthoflastword.go`](./internal/solutions/lengthoflastword.go)                                                           |
|  69  | [Sqrt(x)](https://leetcode.com/problems/sqrtx/)                                                                                                       | [`mysqrt.go`](./internal/solutions/mysqrt.go)                                                                               |
|  70  | [Climbing Stairs](https://leetcode.com/problems/climbing-stairs/)                                                                                     | [`climbstairs.go`](./internal/solutions/climbstairs.go)                                                                     |
|  94  | [Binary Tree Inorder Traversal](https://leetcode.com/problems/binary-tree-inorder-traversal/)                                                         | [`inordertraversal.go`](./internal/solutions/inordertraversal.go)                                                           |
| 118  | [Pascal's Triangle](https://leetcode.com/problems/pascals-triangle/)                                                                                  | [`generate.go`](./internal/solutions/generate.go)                                                                           |
| 119  | [Pascal's Triangle II](https://leetcode.com/problems/pascals-triangle-ii/)                                                                            | [`getrow.go`](./internal/solutions/getrow.go)                                                                               |
| 144  | [Binary Tree Preorder Traversal](https://leetcode.com/problems/binary-tree-preorder-traversal/)                                                       | [`preordertraversal.go`](./internal/solutions/preordertraversal.go)                                                         |
| 145  | [Binary Tree Postorder Traversal](https://leetcode.com/problems/binary-tree-postorder-traversal/)                                                     | [`postordertraversal.go`](./internal/solutions/postordertraversal.go)                                                       |
| 163  | [Missing Ranges](https://leetcode.com/problems/missing-ranges/)                                                                                       | [`findmissingranges.go`](./internal/solutions/findmissingranges.go)                                                         |
| 169  | [Majority Element](https://leetcode.com/problems/majority-element/)                                                                                   | [`majorityelement.go`](./internal/solutions/majorityelement.go)                                                             |
| 191  | [Number of 1 Bits](https://leetcode.com/problems/number-of-1-bits/)                                                                                   | [`hammingweight.go`](./internal/solutions/hammingweight.go)                                                                 |
| 198  | [House Robber](https://leetcode.com/problems/house-robber/)                                                                                           | [`rob.go`](./internal/solutions/rob.go)                                                                                     |
| 202  | [Happy Number](https://leetcode.com/problems/happy-number/)                                                                                           | [`ishappy.go`](./internal/solutions/ishappy.go)                                                                             |
| 231  | [Power of Two](https://leetcode.com/problems/power-of-two/)                                                                                           | [`ispoweroftwo.go`](./internal/solutions/ispoweroftwo.go)                                                                   |
| 342  | [Power of Four](https://leetcode.com/problems/power-of-four/)                                                                                         | [`ispoweroffour.go`](./internal/solutions/ispoweroffour.go)                                                                 |
| 346  | [Moving Average from Data Stream](https://leetcode.com/problems/moving-average-from-data-stream/)                                                     | [`movingaverage.go`](./internal/solutions/movingaverage.go)                                                                 |
| 359  | [Logger Rate Limiter](https://leetcode.com/problems/logger-rate-limiter/)                                                                             | [`logger.go`](./internal/solutions/logger.go)                                                                               |
| 458  | [Poor Pigs](https://leetcode.com/problems/poor-pigs/)                                                                                                 | [`poorpigs.go`](./internal/solutions/poorpigs.go)                                                                           |
| 485  | [Max Consecutive Ones](https://leetcode.com/problems/max-consecutive-ones/)                                                                           | [`findmaxconsecutiveones_1.go`](./internal/solutions/findmaxconsecutiveones_1.go)                                           |
| 487  | [Max Consecutive Ones II](https://leetcode.com/problems/max-consecutive-ones-ii/)                                                                     | [`findmaxconsecutiveones_2.go`](./internal/solutions/findmaxconsecutiveones_2.go)                                           |
| 501  | [Find Mode in Binary Search Tree](https://leetcode.com/problems/find-mode-in-binary-search-tree/)                                                     | [`findmode.go`](./internal/solutions/findmode.go)                                                                           |
| 525  | [Contiguous Array](https://leetcode.com/problems/contiguous-array/)                                                                                   | [`findmaxlength.go`](./internal/solutions/findmaxlength.go)                                                                 |
| 573  | [Squirrel Simulation](https://leetcode.com/problems/squirrel-simulation/)                                                                             | [`mindistance.go`](./internal/solutions/mindistance.go)                                                                     |
| 606  | [Construct String from Binary Tree](https://leetcode.com/problems/construct-string-from-binary-tree/)                                                 | [`tree2str.go`](./internal/solutions/tree2str.go)                                                                           |
| 624  | [Maximum Distance in Arrays](https://leetcode.com/problems/maximum-distance-in-arrays/)                                                               | [`maxdistance.go`](./internal/solutions/maxdistance.go)                                                                     |
| 779  | [K-th Symbol in Grammar](https://leetcode.com/problems/k-th-symbol-in-grammar/)                                                                       | [`kthgrammar.go`](./internal/solutions/kthgrammar.go)                                                                       |
| 844  | [Backspace String Compare](https://leetcode.com/problems/backspace-string-compare/)                                                                   | [`backspacecompare.go`](./internal/solutions/backspacecompare.go)                                                           |
| 867  | [Transpose Matrix](https://leetcode.com/problems/transpose-matrix/)                                                                                   | [`transpose.go`](./internal/solutions/transpose.go)                                                                         |
| 935  | [Knight Dialer](https://leetcode.com/problems/knight-dialer/)                                                                                         | [`knightdialer.go`](./internal/solutions/knightdialer.go)                                                                   |
| 1018 | [Binary Prefix Divisible By 5](https://leetcode.com/problems/binary-prefix-divisible-by-5/)                                                           | [`prefixesdivby5.go`](./internal/solutions/prefixesdivby5.go)                                                               |
| 1057 | [Campus Bikes](https://leetcode.com/problems/campus-bikes/)                                                                                           | [`assignbikes.go`](./internal/solutions/assignbikes.go)                                                                     |
| 1085 | [Sum of Digits in the Minimum Number](https://leetcode.com/problems/sum-of-digits-in-the-minimum-number/)                                             | [`sumofdigits.go`](./internal/solutions/sumofdigits.go)                                                                     |
| 1099 | [Two Sum Less Than K](https://leetcode.com/problems/two-sum-less-than-k/)                                                                             | [`twosumlessthank.go`](./internal/solutions/twosumlessthank.go)                                                             |
| 1120 | [Maximum Average Subtree](https://leetcode.com/problems/maximum-average-subtree/)                                                                     | [`maximumaveragesubtree.go`](./internal/solutions/maximumaveragesubtree.go)                                                 |
| 1160 | [Find Words That Can Be Formed by Characters](https://leetcode.com/problems/find-words-that-can-be-formed-by-characters/)                             | [`countcharacters.go`](./internal/solutions/countcharacters.go)                                                             |
| 1266 | [Minimum Time Visiting All Points](https://leetcode.com/problems/minimum-time-visiting-all-points/)                                                   | [`mintimetovisitallpoints.go`](./internal/solutions/mintimetovisitallpoints.go)                                             |
| 1269 | [Number of Ways to Stay in the Same Place After Some Steps](https://leetcode.com/problems/number-of-ways-to-stay-in-the-same-place-after-some-steps/) | [`numways.go`](./internal/solutions/numways.go)                                                                             |
| 1287 | [Element Appearing More Than 25% In Sorted Array](https://leetcode.com/problems/element-appearing-more-than-25-in-sorted-array/)                      | [`findspecialinteger.go`](./internal/solutions/findspecialinteger.go)                                                       |
| 1356 | [Sort Integers by The Number of 1 Bits](https://leetcode.com/problems/sort-integers-by-the-number-of-1-bits/)                                         | [`sortbybits.go`](./internal/solutions/sortbybits.go)                                                                       |
| 1361 | [Validate Binary Tree Nodes](https://leetcode.com/problems/validate-binary-tree-nodes/)                                                               | [`validatebinarytreenodes.go`](./internal/solutions/validatebinarytreenodes.go)                                             |
| 1424 | [Diagonal Traverse II](https://leetcode.com/problems/diagonal-traverse-ii/)                                                                           | [`finddiagonalorder.go`](./internal/solutions/finddiagonalorder.go)                                                         |
| 1441 | [Build an Array With Stack Operations](https://leetcode.com/problems/build-an-array-with-stack-operations/)                                           | [`buildarray.go`](./internal/solutions/buildarray.go)                                                                       |
| 1464 | [Maximum Product of Two Elements in an Array](https://leetcode.com/problems/maximum-product-of-two-elements-in-an-array/)                             | [`maxproduct.go`](./internal/solutions/maxproduct.go)                                                                       |
| 1503 | [Last Moment Before All Ants Fall Out of a Plank](https://leetcode.com/problems/last-moment-before-all-ants-fall-out-of-a-plank/)                     | [`getlastmoment.go`](./internal/solutions/getlastmoment.go)                                                                 |
| 1535 | [Find the Winner of an Array Game](https://leetcode.com/problems/find-the-winner-of-an-array-game/)                                                   | [`getwinner.go`](./internal/solutions/getwinner.go)                                                                         |
| 1561 | [Maximum Number of Coins You Can Get](https://leetcode.com/problems/maximum-number-of-coins-you-can-get/)                                             | [`maxcoins.go`](./internal/solutions/maxcoins.go)                                                                           |
| 1630 | [Arithmetic Subarrays](https://leetcode.com/problems/arithmetic-subarrays/)                                                                           | [`checkarithmeticsubarrays.go`](./internal/solutions/checkarithmeticsubarrays.go)                                           |
| 1662 | [Check If Two String Arrays are Equivalent](https://leetcode.com/problems/check-if-two-string-arrays-are-equivalent/)                                 | [`arraystringsareequal.go`](./internal/solutions/arraystringsareequal.go)                                                   |
| 1685 | [Sum of Absolute Differences in a Sorted Array](https://leetcode.com/problems/sum-of-absolute-differences-in-a-sorted-array/)                         | [`getsumabsolutedifferences.go`](./internal/solutions/getsumabsolutedifferences.go)                                         |
| 1688 | [Count of Matches in Tournament](https://leetcode.com/problems/count-of-matches-in-tournament/)                                                       | [`numberofmatches.go`](./internal/solutions/numberofmatches.go)                                                             |
| 1716 | [Calculate Money in Leetcode Bank](https://leetcode.com/problems/calculate-money-in-leetcode-bank/)                                                   | [`totalmoney.go`](./internal/solutions/totalmoney.go)                                                                       |
| 1720 | [Decode XORed Array](https://leetcode.com/problems/decode-xored-array/)                                                                               | [`decode.go`](./internal/solutions/decode.go)                                                                               |
| 1743 | [Restore the Array From Adjacent Pairs](https://leetcode.com/problems/restore-the-array-from-adjacent-pairs/)                                         | [`restorearray.go`](./internal/solutions/restorearray.go)                                                                   |
| 1759 | [Count Number of Homogenous Substrings](https://leetcode.com/problems/count-number-of-homogenous-substrings/)                                         | [`counthomogenous.go`](./internal/solutions/counthomogenous.go)                                                             |
| 1814 | [Count Nice Pairs in an Array](https://leetcode.com/problems/count-nice-pairs-in-an-array/)                                                           | [`countnicepairs.go`](./internal/solutions/countnicepairs.go)                                                               |
| 1822 | [Sign of the Product of an Array](https://leetcode.com/problems/sign-of-the-product-of-an-array/)                                                     | [`arraysign.go`](./internal/solutions/arraysign.go)                                                                         |
| 1838 | [Frequency of the Most Frequent Element](https://leetcode.com/problems/frequency-of-the-most-frequent-element/)                                       | [`maxfrequency.go`](./internal/solutions/maxfrequency.go)                                                                   |
| 1845 | [Seat Reservation Manager](https://leetcode.com/problems/seat-reservation-manager/)                                                                   | [`seatmanager.go`](./internal/solutions/seatmanager.go)                                                                     |
| 1846 | [Maximum Element After Decreasing and Rearranging](https://leetcode.com/problems/maximum-element-after-decreasing-and-rearranging/)                   | [`maximumelementafterdecrementingandrearranging.go`](./internal/solutions/maximumelementafterdecrementingandrearranging.go) |
| 1877 | [Minimize Maximum Pair Sum in Array](https://leetcode.com/problems/minimize-maximum-pair-sum-in-array/)                                               | [`minpairsum.go`](./internal/solutions/minpairsum.go)                                                                       |
| 1887 | [Reduction Operations to Make the Array Elements Equal](https://leetcode.com/problems/reduction-operations-to-make-the-array-elements-equal/)         | [`reductionoperations.go`](./internal/solutions/reductionoperations.go)                                                     |
| 1903 | [Largest Odd Number in String](https://leetcode.com/problems/largest-odd-number-in-string/)                                                           | [`largestoddnumber.go`](./internal/solutions/largestoddnumber.go)                                                           |
| 1921 | [Eliminate Maximum Number of Monsters](https://leetcode.com/problems/eliminate-maximum-number-of-monsters/)                                           | [`eliminatemaximum.go`](./internal/solutions/eliminatemaximum.go)                                                           |
| 1930 | [Unique Length-3 Palindromic Subsequences](https://leetcode.com/problems/unique-length-3-palindromic-subsequences/)                                   | [`countpalindromicsubsequence.go`](./internal/solutions/countpalindromicsubsequence.go)                                     |
| 1980 | [Find Unique Binary String](https://leetcode.com/problems/find-unique-binary-string/)                                                                 | [`finddifferentbinarystring.go`](./internal/solutions/finddifferentbinarystring.go)                                         |
| 2259 | [Remove Digit From Number to Maximize Result](https://leetcode.com/problems/remove-digit-from-number-to-maximize-result/)                             | [`removedigit.go`](./internal/solutions/removedigit.go)                                                                     |
| 2264 | [Largest 3-Same-Digit Number in String](https://leetcode.com/problems/largest-3-same-digit-number-in-string/)                                         | [`largestgoodinteger.go`](./internal/solutions/largestgoodinteger.go)                                                       |
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
25. Validate Binary Search Tree
26. Same Tree
27. Symmetric Tree
28. Binary Tree Level Order Traversal
29. Binary Tree Zigzag Level Order Traversal
30. Maximum Depth of Binary Tree
31. Convert Sorted Array to Binary Search Tree
32. Minimum Depth of Binary Tree
33. Path Sum
34. Best Time to Buy and Sell Stock
35. Best Time to Buy and Sell Stock II
36. Valid Palindrome
37. Single Number
38. Linked List Cycle
39. LRU Cache
40. Reverse Words in a String
41. Min Stack
42. Find Peak Element
43. Two Sum II - Input Array Is Sorted
44. Rotate Array
45. Reverse Bits
46. Binary Tree Right Side View
47. Remove Linked List Elements
48. Count Primes
49. Isomorphic Strings
50. Reverse Linked List
51. Minimum Size Subarray Sum
52. Contains Duplicate
53. Implement Stack using Queues
54. Summary Ranges
55. Majority Element II
56. Implement Queue using Stacks
57. Palindrome Linked List
58. Lowest Common Ancestor of a Binary Tree
59. Delete Node in a Linked List
60. Product of Array Except Self
61. Sliding Window Maximum
62. Valid Anagram
63. Missing Number
64. Closest Binary Search Tree Value
65. First Bad Version
66. Move Zeroes
67. Find the Duplicate Number
68. Word Pattern
69. Range Sum Query - Immutable
70. Power of Three
71. Odd Even Linked List
72. Increasing Triplet Subsequence
73. Counting Bits
74. Flatten Nested List Iterator
75. Integer Break
76. Reverse String
77. Reverse Vowels of a String
78. Intersection of Two Arrays II
79. Guess Number Higher or Lower
80. Ransom Note
81. Shuffle an Array
82. First Unique Character in a String
83. Find the Difference
84. Is Subsequence
85. Decode String
86. Fizz Buzz
87. Path Sum III
88. Find All Anagrams in a String
89. String Compression
90. Delete Node in a BST
91. Sort Characters By Frequency
92. Hamming Distance
93. Next Greater Element I
94. Diagonal Traverse
95. Find Largest Value in Each Tree Row
96. Minimum Absolute Difference in BST
97. Diameter of Binary Tree
98. Reverse Words in a String III
99. Subarray Sum Equals K
100. Array Partition
101. Permutation in String
102. Can Place Flowers
103. Maximum Average Subarray I
104. Dota2 Senate
105. Search in a Binary Search Tree
106. Insert into a Binary Search Tree
107. Design HashMap
108. Design Linked List
109. Subarray Product Less Than K
110. Find Pivot Index
111. Asteroid Collision
112. Daily Temperatures
113. Find Smallest Letter Greater Than Target
114. Min Cost Climbing Stairs
115. Largest Number At Least Twice of Others
116. Jewels and Stones
117. Minimum Distance Between BST Nodes
118. Custom Sort String
119. Bus Routes
120. Binary Trees With Factors
121. Peak Index in a Mountain Array
122. Leaf-Similar Trees
123. Middle of the Linked List
124. Online Stock Span
125. Reverse Only Letters
126. Binary Subarrays With Sum
127. Number of Recent Calls
128. Range Sum of BST
129. Validate Stack Sequences
130. Squares of a Sorted Array
131. Max Consecutive Ones III
132. Maximum Difference Between Node and Ancestor
133. Longest Arithmetic Subsequence
134. Remove All Adjacent Duplicates In String
135. Greatest Common Divisor of Strings
136. Find in Mountain Array
137. Print in Order
138. Largest Unique Number
139. N-th Tribonacci Number
140. Snapshot Array
141. Maximum Level Sum of a Binary Tree
142. Maximum Number of Balloons
143. Unique Number of Occurrences
144. Get Equal Substrings Within Budget
145. Count Vowels Permutation
146. Check If It Is a Straight Line
147. Count Number of Nice Subarrays
148. Convert Binary Number in a Linked List to Integer
149. Deepest Leaves Sum
150. Number of Steps to Reduce a Number to Zero
151. Count Negative Numbers in a Sorted Matrix
152. Longest ZigZag Path in a Binary Tree
153. Find Lucky Integer in an Array
154. Minimum Value to Get Positive Step by Step Sum
155. Build Array Where You Can Find The Maximum Exactly K Comparisons
156. Constrained Subsequence Sum
157. Counting Elements
158. Kids With the Greatest Number of Candies
159. Destination City
160. Longest Continuous Subarray With Absolute Diff Less Than or Equal to Limit
161. Count Good Nodes in Binary Tree
162. Maximum Number of Vowels in a Substring of Given Length
163. Max Dot Product of Two Subsequences
164. Running Sum of 1d Array
165. Longest Subarray of 1's After Deleting One Element
166. Path Crossing
167. Can Make Arithmetic Progression From Sequence
168. Number of Good Pairs
169. Make The String Great
170. Determine if Two Strings Are Close
171. Richest Customer Wealth
172. Max Number of K-Sum Pairs
173. Maximum Erasure Value
174. Swapping Nodes in a Linked List
175. Find the Highest Altitude
176. Sum of Unique Elements
177. Merge Strings Alternately
178. Maximum Score of a Good Subarray
179. Check if the Sentence Is Pangram
180. Check if All Characters Have Equal Number of Occurrences
181. Reverse Prefix of Word
182. Minimum Number of Operations to Make Array Continuous
183. Remove Colored Pieces if Both Neighbors are the Same Color
184. Parallel Courses III
185. Reverse Nodes in Even Length Groups
186. K Radius Subarray Averages
187. Delete the Middle Node of a Linked List
188. Maximum Twin Sum of a Linked List
189. Find the Difference of Two Arrays
190. Find Players With Zero or One Losses
191. Add Two Integers
192. Intersection of Multiple Arrays
193. Number of Flowers in Full Bloom
194. Minimum Consecutive Cards to Pick Up
195. Number of Ways to Split Array
196. Successful Pairs of Spells and Potions
197. Max Sum of a Pair With Equal Sum of Digits
198. First Letter to Appear Twice
199. Equal Row and Column Pairs
200. Removing Stars From a String
201. Using a Robot to Print the Lexicographically Smallest String
202. Total Cost to Hire K Workers
203. Design Graph With Shortest Path Calculator
204. Painting the Walls
205. Largest Submatrix With Rearrangements
206. Number of Ways to Divide a Long Corridor
207. Minimum One Bit Operations to Make Integers Zero
