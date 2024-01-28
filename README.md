LeetCode
========

[LeetCode problems](https://leetcode.com/problemset/all/) and how
I [solved them](https://leetcode.com/therenotomorrow/).

Structure
---------

- [Internal](./internal) - all solved and refactored problems
- [Again](legacy/again) - need resolve because of loosing day streak or not enough knowledge
- [Archive](legacy/arch) - all solved problems that in progress refactoring

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

|  #   | Problem                                                                                                                                                                     | Solution                                                                                                                    |
|:----:|:----------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:----------------------------------------------------------------------------------------------------------------------------|
|  1   | [Two Sum](https://leetcode.com/problems/two-sum/description/)                                                                                                               | [`twosum_1.go`](./internal/solutions/twosum_1.go)                                                                           |
|  2   | [Add Two Numbers](https://leetcode.com/problems/add-two-numbers/description/)                                                                                               | [`addtwonumbers.go`](./internal/solutions/addtwonumbers.go)                                                                 |
|  3   | [Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/description/)                                 | [`lengthoflongestsubstring.go`](./internal/solutions/lengthoflongestsubstring.go)                                           |
|  4   | [Median of Two Sorted Arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/description/)                                                                       | [`findmediansortedarrays.go`](./internal/solutions/findmediansortedarrays.go)                                               |
|  7   | [Reverse Integer](https://leetcode.com/problems/reverse-integer/description/)                                                                                               | [`reverse.go`](./internal/solutions/reverse.go)                                                                             |
|  9   | [Palindrome Number](https://leetcode.com/problems/palindrome-number/description/)                                                                                           | [`ispalindrome.go`](./internal/solutions/ispalindrome.go)                                                                   |
|  35  | [Search Insert Position](https://leetcode.com/problems/search-insert-position/description/)                                                                                 | [`searchinsert.go`](./internal/solutions/searchinsert.go)                                                                   |
|  53  | [Maximum Subarray](https://leetcode.com/problems/maximum-subarray/description/)                                                                                             | [`maxsubarray.go`](./internal/solutions/maxsubarray.go)                                                                     |
|  58  | [Length of Last Word](https://leetcode.com/problems/length-of-last-word/description/)                                                                                       | [`lengthoflastword.go`](./internal/solutions/lengthoflastword.go)                                                           |
|  62  | [Unique Paths](https://leetcode.com/problems/unique-paths/description/)                                                                                                     | [`uniquepaths.go`](./internal/solutions/uniquepaths.go)                                                                     |
|  69  | [Sqrt(x)](https://leetcode.com/problems/sqrtx/description/)                                                                                                                 | [`mysqrt.go`](./internal/solutions/mysqrt.go)                                                                               |
|  70  | [Climbing Stairs](https://leetcode.com/problems/climbing-stairs/description/)                                                                                               | [`climbstairs.go`](./internal/solutions/climbstairs.go)                                                                     |
|  94  | [Binary Tree Inorder Traversal](https://leetcode.com/problems/binary-tree-inorder-traversal/description/)                                                                   | [`inordertraversal.go`](./internal/solutions/inordertraversal.go)                                                           |
| 110  | [Balanced Binary Tree](https://leetcode.com/problems/balanced-binary-tree/description/)                                                                                     | [`isbalanced.go`](./internal/solutions/isbalanced.go)                                                                       |
| 118  | [Pascal's Triangle](https://leetcode.com/problems/pascals-triangle/description/)                                                                                            | [`generate.go`](./internal/solutions/generate.go)                                                                           |
| 119  | [Pascal's Triangle II](https://leetcode.com/problems/pascals-triangle-ii/description/)                                                                                      | [`getrow.go`](./internal/solutions/getrow.go)                                                                               |
| 144  | [Binary Tree Preorder Traversal](https://leetcode.com/problems/binary-tree-preorder-traversal/description/)                                                                 | [`preordertraversal.go`](./internal/solutions/preordertraversal.go)                                                         |
| 145  | [Binary Tree Postorder Traversal](https://leetcode.com/problems/binary-tree-postorder-traversal/description/)                                                               | [`postordertraversal.go`](./internal/solutions/postordertraversal.go)                                                       |
| 157  | [Read N Characters Given Read4](https://leetcode.com/problems/read-n-characters-given-read4/description/)                                                                   | [`solution.go`](./internal/solutions/solution.go)                                                                           |
| 159  | [Longest Substring with At Most Two Distinct Characters](https://leetcode.com/problems/longest-substring-with-at-most-two-distinct-characters/description/)                 | [`lengthoflongestsubstringtwodistinct.go`](./internal/solutions/lengthoflongestsubstringtwodistinct.go)                     |
| 160  | [Intersection of Two Linked Lists](https://leetcode.com/problems/intersection-of-two-linked-lists/description/)                                                             | [`getIntersectionNode.go`](./internal/solutions/getIntersectionNode.go)                                                     |
| 163  | [Missing Ranges](https://leetcode.com/problems/missing-ranges/description/)                                                                                                 | [`findmissingranges.go`](./internal/solutions/findmissingranges.go)                                                         |
| 168  | [Excel Sheet Column Title](https://leetcode.com/problems/excel-sheet-column-title/description/)                                                                             | [`converttotitle.go`](./internal/solutions/converttotitle.go)                                                               |
| 169  | [Majority Element](https://leetcode.com/problems/majority-element/description/)                                                                                             | [`majorityelement.go`](./internal/solutions/majorityelement.go)                                                             |
| 171  | [Excel Sheet Column Number](https://leetcode.com/problems/excel-sheet-column-number/description/)                                                                           | [`titletonumber.go`](./internal/solutions/titletonumber.go)                                                                 |
| 191  | [Number of 1 Bits](https://leetcode.com/problems/number-of-1-bits/description/)                                                                                             | [`hammingweight.go`](./internal/solutions/hammingweight.go)                                                                 |
| 198  | [House Robber](https://leetcode.com/problems/house-robber/description/)                                                                                                     | [`rob.go`](./internal/solutions/rob.go)                                                                                     |
| 202  | [Happy Number](https://leetcode.com/problems/happy-number/description/)                                                                                                     | [`ishappy.go`](./internal/solutions/ishappy.go)                                                                             |
| 219  | [Contains Duplicate II](https://leetcode.com/problems/contains-duplicate-ii/description/)                                                                                   | [`containsnearbyduplicate.go`](./internal/solutions/containsnearbyduplicate.go)                                             |
| 222  | [Count Complete Tree Nodes](https://leetcode.com/problems/count-complete-tree-nodes/description/)                                                                           | [`countnodes.go`](./internal/solutions/countnodes.go)                                                                       |
| 226  | [Invert Binary Tree](https://leetcode.com/problems/invert-binary-tree/description/)                                                                                         | [`inverttree.go`](./internal/solutions/inverttree.go)                                                                       |
| 231  | [Power of Two](https://leetcode.com/problems/power-of-two/description/)                                                                                                     | [`ispoweroftwo.go`](./internal/solutions/ispoweroftwo.go)                                                                   |
| 242  | [Valid Anagram](https://leetcode.com/problems/valid-anagram/description/)                                                                                                   | [`isanagram.go`](./internal/solutions/isanagram.go)                                                                         |
| 243  | [Shortest Word Distance](https://leetcode.com/problems/shortest-word-distance/description/)                                                                                 | [`shortestdistance.go`](./internal/solutions/shortestdistance.go)                                                           |
| 246  | [Strobogrammatic Number](https://leetcode.com/problems/strobogrammatic-number/description/)                                                                                 | [`isstrobogrammatic.go`](./internal/solutions/isstrobogrammatic.go)                                                         |
| 258  | [Add Digits](https://leetcode.com/problems/add-digits/description/)                                                                                                         | [`adddigits.go`](./internal/solutions/adddigits.go)                                                                         |
| 296  | [Best Meeting Point](https://leetcode.com/problems/best-meeting-point/description/)                                                                                         | [`mintotaldistance.go`](./internal/solutions/mintotaldistance.go)                                                           |
| 300  | [Longest Increasing Subsequence](https://leetcode.com/problems/longest-increasing-subsequence/description/)                                                                 | [`lengthoflis.go`](./internal/solutions/lengthoflis.go)                                                                     |
| 342  | [Power of Four](https://leetcode.com/problems/power-of-four/description/)                                                                                                   | [`ispoweroffour.go`](./internal/solutions/ispoweroffour.go)                                                                 |
| 346  | [Moving Average from Data Stream](https://leetcode.com/problems/moving-average-from-data-stream/description/)                                                               | [`movingaverage.go`](./internal/solutions/movingaverage.go)                                                                 |
| 359  | [Logger Rate Limiter](https://leetcode.com/problems/logger-rate-limiter/description/)                                                                                       | [`logger.go`](./internal/solutions/logger.go)                                                                               |
| 367  | [Valid Perfect Square](https://leetcode.com/problems/valid-perfect-square/description/)                                                                                     | [`isperfectsquare.go`](./internal/solutions/isperfectsquare.go)                                                             |
| 380  | [Insert Delete GetRandom O(1)](https://leetcode.com/problems/insert-delete-getrandom-o1/description/)                                                                       | [`randomizedset.go`](./internal/solutions/randomizedset.go)                                                                 |
| 455  | [Assign Cookies](https://leetcode.com/problems/assign-cookies/description/)                                                                                                 | [`findcontentchildren.go`](./internal/solutions/findcontentchildren.go)                                                     |
| 458  | [Poor Pigs](https://leetcode.com/problems/poor-pigs/description/)                                                                                                           | [`poorpigs.go`](./internal/solutions/poorpigs.go)                                                                           |
| 485  | [Max Consecutive Ones](https://leetcode.com/problems/max-consecutive-ones/description/)                                                                                     | [`findmaxconsecutiveones_1.go`](./internal/solutions/findmaxconsecutiveones_1.go)                                           |
| 487  | [Max Consecutive Ones II](https://leetcode.com/problems/max-consecutive-ones-ii/description/)                                                                               | [`findmaxconsecutiveones_2.go`](./internal/solutions/findmaxconsecutiveones_2.go)                                           |
| 501  | [Find Mode in Binary Search Tree](https://leetcode.com/problems/find-mode-in-binary-search-tree/description/)                                                               | [`findmode.go`](./internal/solutions/findmode.go)                                                                           |
| 509  | [Fibonacci Number](https://leetcode.com/problems/fibonacci-number/description/)                                                                                             | [`fib.go`](./internal/solutions/fib.go)                                                                                     |
| 525  | [Contiguous Array](https://leetcode.com/problems/contiguous-array/description/)                                                                                             | [`findmaxlength.go`](./internal/solutions/findmaxlength.go)                                                                 |
| 539  | [Minimum Time Difference](https://leetcode.com/problems/minimum-time-difference/description/)                                                                               | [`findmindifference.go`](./internal/solutions/findmindifference.go)                                                         |
| 572  | [Subtree of Another Tree](https://leetcode.com/problems/subtree-of-another-tree/description/)                                                                               | [`issubtree.go`](./internal/solutions/issubtree.go)                                                                         |
| 573  | [Squirrel Simulation](https://leetcode.com/problems/squirrel-simulation/description/)                                                                                       | [`mindistance.go`](./internal/solutions/mindistance.go)                                                                     |
| 576  | [Out of Boundary Paths](https://leetcode.com/problems/out-of-boundary-paths/description/)                                                                                   | [`findpaths.go`](./internal/solutions/findpaths.go)                                                                         |
| 606  | [Construct String from Binary Tree](https://leetcode.com/problems/construct-string-from-binary-tree/description/)                                                           | [`tree2str.go`](./internal/solutions/tree2str.go)                                                                           |
| 624  | [Maximum Distance in Arrays](https://leetcode.com/problems/maximum-distance-in-arrays/description/)                                                                         | [`maxdistance.go`](./internal/solutions/maxdistance.go)                                                                     |
| 645  | [Set Mismatch](https://leetcode.com/problems/set-mismatch/description/)                                                                                                     | [`finderrornums.go`](./internal/solutions/finderrornums.go)                                                                 |
| 661  | [Image Smoother](https://leetcode.com/problems/image-smoother/description/)                                                                                                 | [`imagesmoother.go`](./internal/solutions/imagesmoother.go)                                                                 |
| 704  | [Binary Search](https://leetcode.com/problems/binary-search/description/)                                                                                                   | [`search.go`](./internal/solutions/search.go)                                                                               |
| 746  | [Min Cost Climbing Stairs](https://leetcode.com/problems/min-cost-climbing-stairs/description/)                                                                             | [`mincostclimbingstairs.go`](./internal/solutions/mincostclimbingstairs.go)                                                 |
| 779  | [K-th Symbol in Grammar](https://leetcode.com/problems/k-th-symbol-in-grammar/description/)                                                                                 | [`kthgrammar.go`](./internal/solutions/kthgrammar.go)                                                                       |
| 844  | [Backspace String Compare](https://leetcode.com/problems/backspace-string-compare/description/)                                                                             | [`backspacecompare.go`](./internal/solutions/backspacecompare.go)                                                           |
| 867  | [Transpose Matrix](https://leetcode.com/problems/transpose-matrix/description/)                                                                                             | [`transpose.go`](./internal/solutions/transpose.go)                                                                         |
| 872  | [Leaf-Similar Trees](https://leetcode.com/problems/leaf-similar-trees/description/)                                                                                         | [`leafsimilar.go`](./internal/solutions/leafsimilar.go)                                                                     |
| 905  | [Sort Array By Parity](https://leetcode.com/problems/sort-array-by-parity/description/)                                                                                     | [`sortarraybyparity.go`](./internal/solutions/sortarraybyparity.go)                                                         |
| 907  | [Sum of Subarray Minimums](https://leetcode.com/problems/sum-of-subarray-minimums/description/)                                                                             | [`sumsubarraymins.go`](./internal/solutions/sumsubarraymins.go)                                                             |
| 931  | [Minimum Falling Path Sum](https://leetcode.com/problems/minimum-falling-path-sum/description/)                                                                             | [`minfallingpathsum.go`](./internal/solutions/minfallingpathsum.go)                                                         |
| 935  | [Knight Dialer](https://leetcode.com/problems/knight-dialer/description/)                                                                                                   | [`knightdialer.go`](./internal/solutions/knightdialer.go)                                                                   |
| 938  | [Range Sum of BST](https://leetcode.com/problems/range-sum-of-bst/description/)                                                                                             | [`rangesumbst.go`](./internal/solutions/rangesumbst.go)                                                                     |
| 1002 | [Find Common Characters](https://leetcode.com/problems/find-common-characters/description/)                                                                                 | [`commonchars.go`](./internal/solutions/commonchars.go)                                                                     |
| 1018 | [Binary Prefix Divisible By 5](https://leetcode.com/problems/binary-prefix-divisible-by-5/description/)                                                                     | [`prefixesdivby5.go`](./internal/solutions/prefixesdivby5.go)                                                               |
| 1026 | [Maximum Difference Between Node and Ancestor](https://leetcode.com/problems/maximum-difference-between-node-and-ancestor/description/)                                     | [`maxancestordiff.go`](./internal/solutions/maxancestordiff.go)                                                             |
| 1051 | [Height Checker](https://leetcode.com/problems/height-checker/description/)                                                                                                 | [`heightchecker.go`](./internal/solutions/heightchecker.go)                                                                 |
| 1057 | [Campus Bikes](https://leetcode.com/problems/campus-bikes/description/)                                                                                                     | [`assignbikes_1.go`](./internal/solutions/assignbikes_1.go)                                                                 |
| 1066 | [Campus Bikes II](https://leetcode.com/problems/campus-bikes-ii/description/)                                                                                               | [`assignbikes_2.go`](./internal/solutions/assignbikes_2.go)                                                                 |
| 1085 | [Sum of Digits in the Minimum Number](https://leetcode.com/problems/sum-of-digits-in-the-minimum-number/description/)                                                       | [`sumofdigits.go`](./internal/solutions/sumofdigits.go)                                                                     |
| 1099 | [Two Sum Less Than K](https://leetcode.com/problems/two-sum-less-than-k/description/)                                                                                       | [`twosumlessthank.go`](./internal/solutions/twosumlessthank.go)                                                             |
| 1119 | [Remove Vowels from a String](https://leetcode.com/problems/remove-vowels-from-a-string/description/)                                                                       | [`removevowels.go`](./internal/solutions/removevowels.go)                                                                   |
| 1120 | [Maximum Average Subtree](https://leetcode.com/problems/maximum-average-subtree/description/)                                                                               | [`maximumaveragesubtree.go`](./internal/solutions/maximumaveragesubtree.go)                                                 |
| 1137 | [N-th Tribonacci Number](https://leetcode.com/problems/n-th-tribonacci-number/description/)                                                                                 | [`tribonacci.go`](./internal/solutions/tribonacci.go)                                                                       |
| 1143 | [Longest Common Subsequence](https://leetcode.com/problems/longest-common-subsequence/description/)                                                                         | [`longestcommonsubsequence.go`](./internal/solutions/longestcommonsubsequence.go)                                           |
| 1155 | [Number of Dice Rolls With Target Sum](https://leetcode.com/problems/number-of-dice-rolls-with-target-sum/description/)                                                     | [`numrollstotarget.go`](./internal/solutions/numrollstotarget.go)                                                           |
| 1160 | [Find Words That Can Be Formed by Characters](https://leetcode.com/problems/find-words-that-can-be-formed-by-characters/description/)                                       | [`countcharacters.go`](./internal/solutions/countcharacters.go)                                                             |
| 1165 | [Single-Row Keyboard](https://leetcode.com/problems/single-row-keyboard/description/)                                                                                       | [`calculatetime.go`](./internal/solutions/calculatetime.go)                                                                 |
| 1167 | [Minimum Cost to Connect Sticks](https://leetcode.com/problems/minimum-cost-to-connect-sticks/description/)                                                                 | [`connectsticks.go`](./internal/solutions/connectsticks.go)                                                                 |
| 1207 | [Unique Number of Occurrences](https://leetcode.com/problems/unique-number-of-occurrences/description/)                                                                     | [`uniqueoccurrences.go`](./internal/solutions/uniqueoccurrences.go)                                                         |
| 1216 | [Valid Palindrome III](https://leetcode.com/problems/valid-palindrome-iii/description/)                                                                                     | [`isvalidpalindrome.go`](./internal/solutions/isvalidpalindrome.go)                                                         |
| 1219 | [Path with Maximum Gold](https://leetcode.com/problems/path-with-maximum-gold/description/)                                                                                 | [`getmaximumgold.go`](./internal/solutions/getmaximumgold.go)                                                               |
| 1239 | [Maximum Length of a Concatenated String with Unique Characters](https://leetcode.com/problems/maximum-length-of-a-concatenated-string-with-unique-characters/description/) | [`maxlength.go`](./internal/solutions/maxlength.go)                                                                         |
| 1243 | [Array Transformation](https://leetcode.com/problems/array-transformation/description/)                                                                                     | [`transformarray.go`](./internal/solutions/transformarray.go)                                                               |
| 1266 | [Minimum Time Visiting All Points](https://leetcode.com/problems/minimum-time-visiting-all-points/description/)                                                             | [`mintimetovisitallpoints.go`](./internal/solutions/mintimetovisitallpoints.go)                                             |
| 1269 | [Number of Ways to Stay in the Same Place After Some Steps](https://leetcode.com/problems/number-of-ways-to-stay-in-the-same-place-after-some-steps/description/)           | [`numways.go`](./internal/solutions/numways.go)                                                                             |
| 1287 | [Element Appearing More Than 25% In Sorted Array](https://leetcode.com/problems/element-appearing-more-than-25-in-sorted-array/description/)                                | [`findspecialinteger.go`](./internal/solutions/findspecialinteger.go)                                                       |
| 1295 | [Find Numbers with Even Number of Digits](https://leetcode.com/problems/find-numbers-with-even-number-of-digits/description/)                                               | [`findnumbers.go`](./internal/solutions/findnumbers.go)                                                                     |
| 1347 | [Minimum Number of Steps to Make Two Strings Anagram](https://leetcode.com/problems/minimum-number-of-steps-to-make-two-strings-anagram/description/)                       | [`minsteps.go`](./internal/solutions/minsteps.go)                                                                           |
| 1356 | [Sort Integers by The Number of 1 Bits](https://leetcode.com/problems/sort-integers-by-the-number-of-1-bits/description/)                                                   | [`sortbybits.go`](./internal/solutions/sortbybits.go)                                                                       |
| 1360 | [Number of Days Between Two Dates](https://leetcode.com/problems/number-of-days-between-two-dates/description/)                                                             | [`daysbetweendates.go`](./internal/solutions/daysbetweendates.go)                                                           |
| 1361 | [Validate Binary Tree Nodes](https://leetcode.com/problems/validate-binary-tree-nodes/description/)                                                                         | [`validatebinarytreenodes.go`](./internal/solutions/validatebinarytreenodes.go)                                             |
| 1422 | [Maximum Score After Splitting a String](https://leetcode.com/problems/maximum-score-after-splitting-a-string/description/)                                                 | [`maxscore.go`](./internal/solutions/maxscore.go)                                                                           |
| 1424 | [Diagonal Traverse II](https://leetcode.com/problems/diagonal-traverse-ii/description/)                                                                                     | [`finddiagonalorder.go`](./internal/solutions/finddiagonalorder.go)                                                         |
| 1428 | [Leftmost Column with at Least a One](https://leetcode.com/problems/leftmost-column-with-at-least-a-one/description/)                                                       | [`leftmostcolumnwithone.go`](./internal/solutions/leftmostcolumnwithone.go)                                                 |
| 1436 | [Destination City](https://leetcode.com/problems/destination-city/description/)                                                                                             | [`destcity.go`](./internal/solutions/destcity.go)                                                                           |
| 1441 | [Build an Array With Stack Operations](https://leetcode.com/problems/build-an-array-with-stack-operations/description/)                                                     | [`buildarray.go`](./internal/solutions/buildarray.go)                                                                       |
| 1457 | [Pseudo-Palindromic Paths in a Binary Tree](https://leetcode.com/problems/pseudo-palindromic-paths-in-a-binary-tree/description/)                                           | [`pseudopalindromicpaths.go`](./internal/solutions/pseudopalindromicpaths.go)                                               |
| 1464 | [Maximum Product of Two Elements in an Array](https://leetcode.com/problems/maximum-product-of-two-elements-in-an-array/description/)                                       | [`maxproduct.go`](./internal/solutions/maxproduct.go)                                                                       |
| 1470 | [Shuffle the Array](https://leetcode.com/problems/shuffle-the-array/description/)                                                                                           | [`shuffle.go`](./internal/solutions/shuffle.go)                                                                             |
| 1496 | [Path Crossing](https://leetcode.com/problems/path-crossing/description/)                                                                                                   | [`ispathcrossing.go`](./internal/solutions/ispathcrossing.go)                                                               |
| 1503 | [Last Moment Before All Ants Fall Out of a Plank](https://leetcode.com/problems/last-moment-before-all-ants-fall-out-of-a-plank/description/)                               | [`getlastmoment.go`](./internal/solutions/getlastmoment.go)                                                                 |
| 1535 | [Find the Winner of an Array Game](https://leetcode.com/problems/find-the-winner-of-an-array-game/description/)                                                             | [`getwinner.go`](./internal/solutions/getwinner.go)                                                                         |
| 1561 | [Maximum Number of Coins You Can Get](https://leetcode.com/problems/maximum-number-of-coins-you-can-get/description/)                                                       | [`maxcoins.go`](./internal/solutions/maxcoins.go)                                                                           |
| 1578 | [Minimum Time to Make Rope Colorful](https://leetcode.com/problems/minimum-time-to-make-rope-colorful/description/)                                                         | [`mincost.go`](./internal/solutions/mincost.go)                                                                             |
| 1582 | [Special Positions in a Binary Matrix](https://leetcode.com/problems/special-positions-in-a-binary-matrix/description/)                                                     | [`numspecial.go`](./internal/solutions/numspecial.go)                                                                       |
| 1624 | [Largest Substring Between Two Equal Characters](https://leetcode.com/problems/largest-substring-between-two-equal-characters/description/)                                 | [`maxlengthbetweenequalcharacters.go`](./internal/solutions/maxlengthbetweenequalcharacters.go)                             |
| 1630 | [Arithmetic Subarrays](https://leetcode.com/problems/arithmetic-subarrays/description/)                                                                                     | [`checkarithmeticsubarrays.go`](./internal/solutions/checkarithmeticsubarrays.go)                                           |
| 1637 | [Widest Vertical Area Between Two Points Containing No Points](https://leetcode.com/problems/widest-vertical-area-between-two-points-containing-no-points/description/)     | [`maxwidthofverticalarea.go`](./internal/solutions/maxwidthofverticalarea.go)                                               |
| 1657 | [Determine if Two Strings Are Close](https://leetcode.com/problems/determine-if-two-strings-are-close/description/)                                                         | [`closestrings.go`](./internal/solutions/closestrings.go)                                                                   |
| 1662 | [Check If Two String Arrays are Equivalent](https://leetcode.com/problems/check-if-two-string-arrays-are-equivalent/description/)                                           | [`arraystringsareequal.go`](./internal/solutions/arraystringsareequal.go)                                                   |
| 1685 | [Sum of Absolute Differences in a Sorted Array](https://leetcode.com/problems/sum-of-absolute-differences-in-a-sorted-array/description/)                                   | [`getsumabsolutedifferences.go`](./internal/solutions/getsumabsolutedifferences.go)                                         |
| 1688 | [Count of Matches in Tournament](https://leetcode.com/problems/count-of-matches-in-tournament/description/)                                                                 | [`numberofmatches.go`](./internal/solutions/numberofmatches.go)                                                             |
| 1704 | [Determine if String Halves Are Alike](https://leetcode.com/problems/determine-if-string-halves-are-alike/description/)                                                     | [`halvesarealike.go`](./internal/solutions/halvesarealike.go)                                                               |
| 1716 | [Calculate Money in Leetcode Bank](https://leetcode.com/problems/calculate-money-in-leetcode-bank/description/)                                                             | [`totalmoney.go`](./internal/solutions/totalmoney.go)                                                                       |
| 1720 | [Decode XORed Array](https://leetcode.com/problems/decode-xored-array/description/)                                                                                         | [`decode.go`](./internal/solutions/decode.go)                                                                               |
| 1743 | [Restore the Array From Adjacent Pairs](https://leetcode.com/problems/restore-the-array-from-adjacent-pairs/description/)                                                   | [`restorearray.go`](./internal/solutions/restorearray.go)                                                                   |
| 1758 | [Minimum Changes To Make Alternating Binary String](https://leetcode.com/problems/minimum-changes-to-make-alternating-binary-string/description/)                           | [`minoperations_1.go`](./internal/solutions/minoperations_1.go)                                                             |
| 1759 | [Count Number of Homogenous Substrings](https://leetcode.com/problems/count-number-of-homogenous-substrings/description/)                                                   | [`counthomogenous.go`](./internal/solutions/counthomogenous.go)                                                             |
| 1814 | [Count Nice Pairs in an Array](https://leetcode.com/problems/count-nice-pairs-in-an-array/description/)                                                                     | [`countnicepairs.go`](./internal/solutions/countnicepairs.go)                                                               |
| 1822 | [Sign of the Product of an Array](https://leetcode.com/problems/sign-of-the-product-of-an-array/description/)                                                               | [`arraysign.go`](./internal/solutions/arraysign.go)                                                                         |
| 1838 | [Frequency of the Most Frequent Element](https://leetcode.com/problems/frequency-of-the-most-frequent-element/description/)                                                 | [`maxfrequency.go`](./internal/solutions/maxfrequency.go)                                                                   |
| 1845 | [Seat Reservation Manager](https://leetcode.com/problems/seat-reservation-manager/description/)                                                                             | [`seatmanager.go`](./internal/solutions/seatmanager.go)                                                                     |
| 1846 | [Maximum Element After Decreasing and Rearranging](https://leetcode.com/problems/maximum-element-after-decreasing-and-rearranging/description/)                             | [`maximumelementafterdecrementingandrearranging.go`](./internal/solutions/maximumelementafterdecrementingandrearranging.go) |
| 1877 | [Minimize Maximum Pair Sum in Array](https://leetcode.com/problems/minimize-maximum-pair-sum-in-array/description/)                                                         | [`minpairsum.go`](./internal/solutions/minpairsum.go)                                                                       |
| 1887 | [Reduction Operations to Make the Array Elements Equal](https://leetcode.com/problems/reduction-operations-to-make-the-array-elements-equal/description/)                   | [`reductionoperations.go`](./internal/solutions/reductionoperations.go)                                                     |
| 1897 | [Redistribute Characters to Make All Strings Equal](https://leetcode.com/problems/redistribute-characters-to-make-all-strings-equal/description/)                           | [`makeequal.go`](./internal/solutions/makeequal.go)                                                                         |
| 1903 | [Largest Odd Number in String](https://leetcode.com/problems/largest-odd-number-in-string/description/)                                                                     | [`largestoddnumber.go`](./internal/solutions/largestoddnumber.go)                                                           |
| 1913 | [Maximum Product Difference Between Two Pairs](https://leetcode.com/problems/maximum-product-difference-between-two-pairs/description/)                                     | [`maxproductdifference.go`](./internal/solutions/maxproductdifference.go)                                                   |
| 1921 | [Eliminate Maximum Number of Monsters](https://leetcode.com/problems/eliminate-maximum-number-of-monsters/description/)                                                     | [`eliminatemaximum.go`](./internal/solutions/eliminatemaximum.go)                                                           |
| 1930 | [Unique Length-3 Palindromic Subsequences](https://leetcode.com/problems/unique-length-3-palindromic-subsequences/description/)                                             | [`countpalindromicsubsequence.go`](./internal/solutions/countpalindromicsubsequence.go)                                     |
| 1961 | [Check If String Is a Prefix of Array](https://leetcode.com/problems/check-if-string-is-a-prefix-of-array/description/)                                                     | [`isprefixstring.go`](./internal/solutions/isprefixstring.go)                                                               |
| 1980 | [Find Unique Binary String](https://leetcode.com/problems/find-unique-binary-string/description/)                                                                           | [`finddifferentbinarystring.go`](./internal/solutions/finddifferentbinarystring.go)                                         |
| 2125 | [Number of Laser Beams in a Bank](https://leetcode.com/problems/number-of-laser-beams-in-a-bank/description/)                                                               | [`numberofbeams.go`](./internal/solutions/numberofbeams.go)                                                                 |
| 2225 | [Find Players With Zero or One Losses](https://leetcode.com/problems/find-players-with-zero-or-one-losses/description/)                                                     | [`findwinners.go`](./internal/solutions/findwinners.go)                                                                     |
| 2244 | [Minimum Rounds to Complete All Tasks](https://leetcode.com/problems/minimum-rounds-to-complete-all-tasks/description/)                                                     | [`minimumrounds.go`](./internal/solutions/minimumrounds.go)                                                                 |
| 2259 | [Remove Digit From Number to Maximize Result](https://leetcode.com/problems/remove-digit-from-number-to-maximize-result/description/)                                       | [`removedigit.go`](./internal/solutions/removedigit.go)                                                                     |
| 2264 | [Largest 3-Same-Digit Number in String](https://leetcode.com/problems/largest-3-same-digit-number-in-string/description/)                                                   | [`largestgoodinteger.go`](./internal/solutions/largestgoodinteger.go)                                                       |
| 2265 | [Count Nodes Equal to Average of Subtree](https://leetcode.com/problems/count-nodes-equal-to-average-of-subtree/description/)                                               | [`averageofsubtree.go`](./internal/solutions/averageofsubtree.go)                                                           |
| 2353 | [Design a Food Rating System](https://leetcode.com/problems/design-a-food-rating-system/description/)                                                                       | [`foodratings.go`](./internal/solutions/foodratings.go)                                                                     |
| 2357 | [Make Array Zero by Subtracting Equal Amounts](https://leetcode.com/problems/make-array-zero-by-subtracting-equal-amounts/description/)                                     | [`minimumoperations.go`](./internal/solutions/minimumoperations.go)                                                         |
| 2385 | [Amount of Time for Binary Tree to Be Infected](https://leetcode.com/problems/amount-of-time-for-binary-tree-to-be-infected/description/)                                   | [`amountoftime.go`](./internal/solutions/amountoftime.go)                                                                   |
| 2391 | [Minimum Amount of Time to Collect Garbage](https://leetcode.com/problems/minimum-amount-of-time-to-collect-garbage/description/)                                           | [`garbagecollection.go`](./internal/solutions/garbagecollection.go)                                                         |
| 2433 | [Find The Original Array of Prefix Xor](https://leetcode.com/problems/find-the-original-array-of-prefix-xor/description/)                                                   | [`findarray.go`](./internal/solutions/findarray.go)                                                                         |
| 2482 | [Difference Between Ones and Zeros in Row and Column](https://leetcode.com/problems/difference-between-ones-and-zeros-in-row-and-column/description/)                       | [`onesminuszeros.go`](./internal/solutions/onesminuszeros.go)                                                               |
| 2610 | [Convert an Array Into a 2D Array With Conditions](https://leetcode.com/problems/convert-an-array-into-a-2d-array-with-conditions/description/)                             | [`findmatrix.go`](./internal/solutions/findmatrix.go)                                                                       |
| 2706 | [Buy Two Chocolates](https://leetcode.com/problems/buy-two-chocolates/description/)                                                                                         | [`buychoco.go`](./internal/solutions/buychoco.go)                                                                           |
| 2785 | [Sort Vowels in a String](https://leetcode.com/problems/sort-vowels-in-a-string/description/)                                                                               | [`sortvowels.go`](./internal/solutions/sortvowels.go)                                                                       |
| 2849 | [Determine if a Cell Is Reachable at a Given Time](https://leetcode.com/problems/determine-if-a-cell-is-reachable-at-a-given-time/description/)                             | [`isreachableattime.go`](./internal/solutions/isreachableattime.go)                                                         |
| 2870 | [Minimum Number of Operations to Make Array Empty](https://leetcode.com/problems/minimum-number-of-operations-to-make-array-empty/description/)                             | [`minoperations_2.go`](./internal/solutions/minoperations_2.go)                                                             |

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
62. Missing Number
63. Closest Binary Search Tree Value
64. First Bad Version
65. Move Zeroes
66. Find the Duplicate Number
67. Word Pattern
68. Range Sum Query - Immutable
69. Power of Three
70. Odd Even Linked List
71. Increasing Triplet Subsequence
72. Counting Bits
73. Flatten Nested List Iterator
74. Integer Break
75. Reverse String
76. Reverse Vowels of a String
77. Intersection of Two Arrays II
78. Guess Number Higher or Lower
79. Ransom Note
80. Shuffle an Array
81. First Unique Character in a String
82. Find the Difference
83. Is Subsequence
84. Decode String
85. Fizz Buzz
86. Path Sum III
87. Find All Anagrams in a String
88. String Compression
89. Delete Node in a BST
90. Sort Characters By Frequency
91. Hamming Distance
92. Next Greater Element I
93. Diagonal Traverse
94. Find Largest Value in Each Tree Row
95. Minimum Absolute Difference in BST
96. Diameter of Binary Tree
97. Reverse Words in a String III
98. Subarray Sum Equals K
99. Array Partition
100. Permutation in String
101. Can Place Flowers
102. Maximum Average Subarray I
103. Dota2 Senate
104. Search in a Binary Search Tree
105. Insert into a Binary Search Tree
106. Design HashMap
107. Design Linked List
108. Subarray Product Less Than K
109. Find Pivot Index
110. Asteroid Collision
111. Daily Temperatures
112. Find Smallest Letter Greater Than Target
113. Largest Number At Least Twice of Others
114. Jewels and Stones
115. Minimum Distance Between BST Nodes
116. Custom Sort String
117. Bus Routes
118. Binary Trees With Factors
119. Peak Index in a Mountain Array
120. Middle of the Linked List
121. Online Stock Span
122. Reverse Only Letters
123. Binary Subarrays With Sum
124. Number of Recent Calls
125. Validate Stack Sequences
126. Squares of a Sorted Array
127. Max Consecutive Ones III
128. Longest Arithmetic Subsequence
129. Remove All Adjacent Duplicates In String
130. Greatest Common Divisor of Strings
131. Find in Mountain Array
132. Print in Order
133. Largest Unique Number
134. Snapshot Array
135. Maximum Level Sum of a Binary Tree
136. Maximum Number of Balloons
137. Get Equal Substrings Within Budget
138. Count Vowels Permutation
139. Check If It Is a Straight Line
140. Count Number of Nice Subarrays
141. Convert Binary Number in a Linked List to Integer
142. Deepest Leaves Sum
143. Number of Steps to Reduce a Number to Zero
144. Count Negative Numbers in a Sorted Matrix
145. Longest ZigZag Path in a Binary Tree
146. Find Lucky Integer in an Array
147. Minimum Value to Get Positive Step by Step Sum
148. Build Array Where You Can Find The Maximum Exactly K Comparisons
149. Constrained Subsequence Sum
150. Counting Elements
151. Kids With the Greatest Number of Candies
152. Longest Continuous Subarray With Absolute Diff Less Than or Equal to Limit
153. Count Good Nodes in Binary Tree
154. Maximum Number of Vowels in a Substring of Given Length
155. Max Dot Product of Two Subsequences
156. Running Sum of 1d Array
157. Longest Subarray of 1's After Deleting One Element
158. Can Make Arithmetic Progression From Sequence
159. Number of Good Pairs
160. Make The String Great
161. Richest Customer Wealth
162. Max Number of K-Sum Pairs
163. Maximum Erasure Value
164. Swapping Nodes in a Linked List
165. Find the Highest Altitude
166. Sum of Unique Elements
167. Merge Strings Alternately
168. Maximum Score of a Good Subarray
169. Check if the Sentence Is Pangram
170. Check if All Characters Have Equal Number of Occurrences
171. Reverse Prefix of Word
172. Minimum Number of Operations to Make Array Continuous
173. Remove Colored Pieces if Both Neighbors are the Same Color
174. Parallel Courses III
175. Reverse Nodes in Even Length Groups
176. K Radius Subarray Averages
177. Delete the Middle Node of a Linked List
178. Maximum Twin Sum of a Linked List
179. Find the Difference of Two Arrays
180. Add Two Integers
181. Intersection of Multiple Arrays
182. Number of Flowers in Full Bloom
183. Minimum Consecutive Cards to Pick Up
184. Number of Ways to Split Array
185. Successful Pairs of Spells and Potions
186. Max Sum of a Pair With Equal Sum of Digits
187. First Letter to Appear Twice
188. Equal Row and Column Pairs
189. Removing Stars From a String
190. Using a Robot to Print the Lexicographically Smallest String
191. Total Cost to Hire K Workers
192. Design Graph With Shortest Path Calculator
193. Painting the Walls
194. Largest Submatrix With Rearrangements
195. Number of Ways to Divide a Long Corridor
196. Minimum One Bit Operations to Make Integers Zero
197. Decode Ways
198. String Compression II
199. Minimum Difficulty of a Job Schedule
200. Maximum Profit in Job Scheduling
201. Arithmetic Slices II - Subsequence
202. K Inverse Pairs Array
203. Number of Submatrices That Sum to Target
