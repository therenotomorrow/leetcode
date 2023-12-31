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

|  #   | Problem                                                                                                                                                     | Solution                                                                                                                    |
|:----:|:------------------------------------------------------------------------------------------------------------------------------------------------------------|:----------------------------------------------------------------------------------------------------------------------------|
|  1   | [Two Sum](https://leetcode.com/problems/two-sum/)                                                                                                           | [`twosum_1.go`](./internal/solutions/twosum_1.go)                                                                           |
|  2   | [Add Two Numbers](https://leetcode.com/problems/add-two-numbers/)                                                                                           | [`addtwonumbers.go`](./internal/solutions/addtwonumbers.go)                                                                 |
|  3   | [Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/)                             | [`lengthoflongestsubstring.go`](./internal/solutions/lengthoflongestsubstring.go)                                           |
|  4   | [Median of Two Sorted Arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/)                                                                   | [`findmediansortedarrays.go`](./internal/solutions/findmediansortedarrays.go)                                               |
|  7   | [Reverse Integer](https://leetcode.com/problems/reverse-integer/)                                                                                           | [`reverse.go`](./internal/solutions/reverse.go)                                                                             |
|  9   | [Palindrome Number](https://leetcode.com/problems/palindrome-number/)                                                                                       | [`ispalindrome.go`](./internal/solutions/ispalindrome.go)                                                                   |
|  35  | [Search Insert Position](https://leetcode.com/problems/search-insert-position/)                                                                             | [`searchinsert.go`](./internal/solutions/searchinsert.go)                                                                   |
|  53  | [Maximum Subarray](https://leetcode.com/problems/maximum-subarray/)                                                                                         | [`maxsubarray.go`](./internal/solutions/maxsubarray.go)                                                                     |
|  58  | [Length of Last Word](https://leetcode.com/problems/length-of-last-word/)                                                                                   | [`lengthoflastword.go`](./internal/solutions/lengthoflastword.go)                                                           |
|  69  | [Sqrt(x)](https://leetcode.com/problems/sqrtx/)                                                                                                             | [`mysqrt.go`](./internal/solutions/mysqrt.go)                                                                               |
|  70  | [Climbing Stairs](https://leetcode.com/problems/climbing-stairs/)                                                                                           | [`climbstairs.go`](./internal/solutions/climbstairs.go)                                                                     |
|  94  | [Binary Tree Inorder Traversal](https://leetcode.com/problems/binary-tree-inorder-traversal/)                                                               | [`inordertraversal.go`](./internal/solutions/inordertraversal.go)                                                           |
| 110  | [Balanced Binary Tree](https://leetcode.com/problems/balanced-binary-tree/)                                                                                 | [`isbalanced.go`](./internal/solutions/isbalanced.go)                                                                       |
| 118  | [Pascal's Triangle](https://leetcode.com/problems/pascals-triangle/)                                                                                        | [`generate.go`](./internal/solutions/generate.go)                                                                           |
| 119  | [Pascal's Triangle II](https://leetcode.com/problems/pascals-triangle-ii/)                                                                                  | [`getrow.go`](./internal/solutions/getrow.go)                                                                               |
| 144  | [Binary Tree Preorder Traversal](https://leetcode.com/problems/binary-tree-preorder-traversal/)                                                             | [`preordertraversal.go`](./internal/solutions/preordertraversal.go)                                                         |
| 145  | [Binary Tree Postorder Traversal](https://leetcode.com/problems/binary-tree-postorder-traversal/)                                                           | [`postordertraversal.go`](./internal/solutions/postordertraversal.go)                                                       |
| 157  | [Read N Characters Given Read4](https://leetcode.com/problems/read-n-characters-given-read4/)                                                               | [`solution.go`](./internal/solutions/solution.go)                                                                           |
| 159  | [Longest Substring with At Most Two Distinct Characters](https://leetcode.com/problems/longest-substring-with-at-most-two-distinct-characters/)             | [`lengthoflongestsubstringtwodistinct.go`](./internal/solutions/lengthoflongestsubstringtwodistinct.go)                     |
| 163  | [Missing Ranges](https://leetcode.com/problems/missing-ranges/)                                                                                             | [`findmissingranges.go`](./internal/solutions/findmissingranges.go)                                                         |
| 168  | [Excel Sheet Column Title](https://leetcode.com/problems/excel-sheet-column-title/)                                                                         | [`converttotitle.go`](./internal/solutions/converttotitle.go)                                                               |
| 169  | [Majority Element](https://leetcode.com/problems/majority-element/)                                                                                         | [`majorityelement.go`](./internal/solutions/majorityelement.go)                                                             |
| 171  | [Excel Sheet Column Number](https://leetcode.com/problems/excel-sheet-column-number/)                                                                       | [`titletonumber.go`](./internal/solutions/titletonumber.go)                                                                 |
| 191  | [Number of 1 Bits](https://leetcode.com/problems/number-of-1-bits/)                                                                                         | [`hammingweight.go`](./internal/solutions/hammingweight.go)                                                                 |
| 198  | [House Robber](https://leetcode.com/problems/house-robber/)                                                                                                 | [`rob.go`](./internal/solutions/rob.go)                                                                                     |
| 202  | [Happy Number](https://leetcode.com/problems/happy-number/)                                                                                                 | [`ishappy.go`](./internal/solutions/ishappy.go)                                                                             |
| 219  | [Contains Duplicate II](https://leetcode.com/problems/contains-duplicate-ii/)                                                                               | [`containsnearbyduplicate.go`](./internal/solutions/containsnearbyduplicate.go)                                             |
| 231  | [Power of Two](https://leetcode.com/problems/power-of-two/)                                                                                                 | [`ispoweroftwo.go`](./internal/solutions/ispoweroftwo.go)                                                                   |
| 242  | [Valid Anagram](https://leetcode.com/problems/valid-anagram/)                                                                                               | [`isanagram.go`](./internal/solutions/isanagram.go)                                                                         |
| 296  | [Best Meeting Point](https://leetcode.com/problems/best-meeting-point/)                                                                                     | [`mintotaldistance.go`](./internal/solutions/mintotaldistance.go)                                                           |
| 300  | [Longest Increasing Subsequence](https://leetcode.com/problems/longest-increasing-subsequence/)                                                             | [`lengthoflis.go`](./internal/solutions/lengthoflis.go)                                                                     |
| 342  | [Power of Four](https://leetcode.com/problems/power-of-four/)                                                                                               | [`ispoweroffour.go`](./internal/solutions/ispoweroffour.go)                                                                 |
| 346  | [Moving Average from Data Stream](https://leetcode.com/problems/moving-average-from-data-stream/)                                                           | [`movingaverage.go`](./internal/solutions/movingaverage.go)                                                                 |
| 359  | [Logger Rate Limiter](https://leetcode.com/problems/logger-rate-limiter/)                                                                                   | [`logger.go`](./internal/solutions/logger.go)                                                                               |
| 455  | [Assign Cookies](https://leetcode.com/problems/assign-cookies/)                                                                                             | [`findcontentchildren.go`](./internal/solutions/findcontentchildren.go)                                                     |
| 458  | [Poor Pigs](https://leetcode.com/problems/poor-pigs/)                                                                                                       | [`poorpigs.go`](./internal/solutions/poorpigs.go)                                                                           |
| 485  | [Max Consecutive Ones](https://leetcode.com/problems/max-consecutive-ones/)                                                                                 | [`findmaxconsecutiveones_1.go`](./internal/solutions/findmaxconsecutiveones_1.go)                                           |
| 487  | [Max Consecutive Ones II](https://leetcode.com/problems/max-consecutive-ones-ii/)                                                                           | [`findmaxconsecutiveones_2.go`](./internal/solutions/findmaxconsecutiveones_2.go)                                           |
| 501  | [Find Mode in Binary Search Tree](https://leetcode.com/problems/find-mode-in-binary-search-tree/)                                                           | [`findmode.go`](./internal/solutions/findmode.go)                                                                           |
| 525  | [Contiguous Array](https://leetcode.com/problems/contiguous-array/)                                                                                         | [`findmaxlength.go`](./internal/solutions/findmaxlength.go)                                                                 |
| 573  | [Squirrel Simulation](https://leetcode.com/problems/squirrel-simulation/)                                                                                   | [`mindistance.go`](./internal/solutions/mindistance.go)                                                                     |
| 606  | [Construct String from Binary Tree](https://leetcode.com/problems/construct-string-from-binary-tree/)                                                       | [`tree2str.go`](./internal/solutions/tree2str.go)                                                                           |
| 624  | [Maximum Distance in Arrays](https://leetcode.com/problems/maximum-distance-in-arrays/)                                                                     | [`maxdistance.go`](./internal/solutions/maxdistance.go)                                                                     |
| 661  | [Image Smoother](https://leetcode.com/problems/image-smoother/)                                                                                             | [`imagesmoother.go`](./internal/solutions/imagesmoother.go)                                                                 |
| 704  | [Binary Search](https://leetcode.com/problems/binary-search/)                                                                                               | [`search.go`](./internal/solutions/search.go)                                                                               |
| 779  | [K-th Symbol in Grammar](https://leetcode.com/problems/k-th-symbol-in-grammar/)                                                                             | [`kthgrammar.go`](./internal/solutions/kthgrammar.go)                                                                       |
| 844  | [Backspace String Compare](https://leetcode.com/problems/backspace-string-compare/)                                                                         | [`backspacecompare.go`](./internal/solutions/backspacecompare.go)                                                           |
| 867  | [Transpose Matrix](https://leetcode.com/problems/transpose-matrix/)                                                                                         | [`transpose.go`](./internal/solutions/transpose.go)                                                                         |
| 935  | [Knight Dialer](https://leetcode.com/problems/knight-dialer/)                                                                                               | [`knightdialer.go`](./internal/solutions/knightdialer.go)                                                                   |
| 1018 | [Binary Prefix Divisible By 5](https://leetcode.com/problems/binary-prefix-divisible-by-5/)                                                                 | [`prefixesdivby5.go`](./internal/solutions/prefixesdivby5.go)                                                               |
| 1057 | [Campus Bikes](https://leetcode.com/problems/campus-bikes/)                                                                                                 | [`assignbikes_1.go`](./internal/solutions/assignbikes_1.go)                                                                 |
| 1066 | [Campus Bikes II](https://leetcode.com/problems/campus-bikes-ii/)                                                                                           | [`assignbikes_2.go`](./internal/solutions/assignbikes_2.go)                                                                 |
| 1085 | [Sum of Digits in the Minimum Number](https://leetcode.com/problems/sum-of-digits-in-the-minimum-number/)                                                   | [`sumofdigits.go`](./internal/solutions/sumofdigits.go)                                                                     |
| 1099 | [Two Sum Less Than K](https://leetcode.com/problems/two-sum-less-than-k/)                                                                                   | [`twosumlessthank.go`](./internal/solutions/twosumlessthank.go)                                                             |
| 1119 | [Remove Vowels from a String](https://leetcode.com/problems/remove-vowels-from-a-string/)                                                                   | [`removevowels.go`](./internal/solutions/removevowels.go)                                                                   |
| 1120 | [Maximum Average Subtree](https://leetcode.com/problems/maximum-average-subtree/)                                                                           | [`maximumaveragesubtree.go`](./internal/solutions/maximumaveragesubtree.go)                                                 |
| 1155 | [Number of Dice Rolls With Target Sum](https://leetcode.com/problems/number-of-dice-rolls-with-target-sum/)                                                 | [`numrollstotarget.go`](./internal/solutions/numrollstotarget.go)                                                           |
| 1160 | [Find Words That Can Be Formed by Characters](https://leetcode.com/problems/find-words-that-can-be-formed-by-characters/)                                   | [`countcharacters.go`](./internal/solutions/countcharacters.go)                                                             |
| 1216 | [Valid Palindrome III](https://leetcode.com/problems/valid-palindrome-iii/)                                                                                 | [`isvalidpalindrome.go`](./internal/solutions/isvalidpalindrome.go)                                                         |
| 1266 | [Minimum Time Visiting All Points](https://leetcode.com/problems/minimum-time-visiting-all-points/)                                                         | [`mintimetovisitallpoints.go`](./internal/solutions/mintimetovisitallpoints.go)                                             |
| 1269 | [Number of Ways to Stay in the Same Place After Some Steps](https://leetcode.com/problems/number-of-ways-to-stay-in-the-same-place-after-some-steps/)       | [`numways.go`](./internal/solutions/numways.go)                                                                             |
| 1287 | [Element Appearing More Than 25% In Sorted Array](https://leetcode.com/problems/element-appearing-more-than-25-in-sorted-array/)                            | [`findspecialinteger.go`](./internal/solutions/findspecialinteger.go)                                                       |
| 1356 | [Sort Integers by The Number of 1 Bits](https://leetcode.com/problems/sort-integers-by-the-number-of-1-bits/)                                               | [`sortbybits.go`](./internal/solutions/sortbybits.go)                                                                       |
| 1361 | [Validate Binary Tree Nodes](https://leetcode.com/problems/validate-binary-tree-nodes/)                                                                     | [`validatebinarytreenodes.go`](./internal/solutions/validatebinarytreenodes.go)                                             |
| 1422 | [Maximum Score After Splitting a String](https://leetcode.com/problems/maximum-score-after-splitting-a-string/)                                             | [`maxscore.go`](./internal/solutions/maxscore.go)                                                                           |
| 1424 | [Diagonal Traverse II](https://leetcode.com/problems/diagonal-traverse-ii/)                                                                                 | [`finddiagonalorder.go`](./internal/solutions/finddiagonalorder.go)                                                         |
| 1436 | [Destination City](https://leetcode.com/problems/destination-city/)                                                                                         | [`destcity.go`](./internal/solutions/destcity.go)                                                                           |
| 1441 | [Build an Array With Stack Operations](https://leetcode.com/problems/build-an-array-with-stack-operations/)                                                 | [`buildarray.go`](./internal/solutions/buildarray.go)                                                                       |
| 1464 | [Maximum Product of Two Elements in an Array](https://leetcode.com/problems/maximum-product-of-two-elements-in-an-array/)                                   | [`maxproduct.go`](./internal/solutions/maxproduct.go)                                                                       |
| 1470 | [Shuffle the Array](https://leetcode.com/problems/shuffle-the-array/)                                                                                       | [`shuffle.go`](./internal/solutions/shuffle.go)                                                                             |
| 1496 | [Path Crossing](https://leetcode.com/problems/path-crossing/)                                                                                               | [`ispathcrossing.go`](./internal/solutions/ispathcrossing.go)                                                               |
| 1503 | [Last Moment Before All Ants Fall Out of a Plank](https://leetcode.com/problems/last-moment-before-all-ants-fall-out-of-a-plank/)                           | [`getlastmoment.go`](./internal/solutions/getlastmoment.go)                                                                 |
| 1535 | [Find the Winner of an Array Game](https://leetcode.com/problems/find-the-winner-of-an-array-game/)                                                         | [`getwinner.go`](./internal/solutions/getwinner.go)                                                                         |
| 1561 | [Maximum Number of Coins You Can Get](https://leetcode.com/problems/maximum-number-of-coins-you-can-get/)                                                   | [`maxcoins.go`](./internal/solutions/maxcoins.go)                                                                           |
| 1578 | [Minimum Time to Make Rope Colorful](https://leetcode.com/problems/minimum-time-to-make-rope-colorful/)                                                     | [`mincost.go`](./internal/solutions/mincost.go)                                                                             |
| 1582 | [Special Positions in a Binary Matrix](https://leetcode.com/problems/special-positions-in-a-binary-matrix/)                                                 | [`numspecial.go`](./internal/solutions/numspecial.go)                                                                       |
| 1624 | [Largest Substring Between Two Equal Characters](https://leetcode.com/problems/largest-substring-between-two-equal-characters/)                             | [`maxlengthbetweenequalcharacters.go`](./internal/solutions/maxlengthbetweenequalcharacters.go)                             |
| 1630 | [Arithmetic Subarrays](https://leetcode.com/problems/arithmetic-subarrays/)                                                                                 | [`checkarithmeticsubarrays.go`](./internal/solutions/checkarithmeticsubarrays.go)                                           |
| 1637 | [Widest Vertical Area Between Two Points Containing No Points](https://leetcode.com/problems/widest-vertical-area-between-two-points-containing-no-points/) | [`maxwidthofverticalarea.go`](./internal/solutions/maxwidthofverticalarea.go)                                               |
| 1662 | [Check If Two String Arrays are Equivalent](https://leetcode.com/problems/check-if-two-string-arrays-are-equivalent/)                                       | [`arraystringsareequal.go`](./internal/solutions/arraystringsareequal.go)                                                   |
| 1685 | [Sum of Absolute Differences in a Sorted Array](https://leetcode.com/problems/sum-of-absolute-differences-in-a-sorted-array/)                               | [`getsumabsolutedifferences.go`](./internal/solutions/getsumabsolutedifferences.go)                                         |
| 1688 | [Count of Matches in Tournament](https://leetcode.com/problems/count-of-matches-in-tournament/)                                                             | [`numberofmatches.go`](./internal/solutions/numberofmatches.go)                                                             |
| 1716 | [Calculate Money in Leetcode Bank](https://leetcode.com/problems/calculate-money-in-leetcode-bank/)                                                         | [`totalmoney.go`](./internal/solutions/totalmoney.go)                                                                       |
| 1720 | [Decode XORed Array](https://leetcode.com/problems/decode-xored-array/)                                                                                     | [`decode.go`](./internal/solutions/decode.go)                                                                               |
| 1743 | [Restore the Array From Adjacent Pairs](https://leetcode.com/problems/restore-the-array-from-adjacent-pairs/)                                               | [`restorearray.go`](./internal/solutions/restorearray.go)                                                                   |
| 1758 | [Minimum Changes To Make Alternating Binary String](https://leetcode.com/problems/minimum-changes-to-make-alternating-binary-string/)                       | [`minoperations_1.go`](./internal/solutions/minoperations_1.go)                                                             |
| 1759 | [Count Number of Homogenous Substrings](https://leetcode.com/problems/count-number-of-homogenous-substrings/)                                               | [`counthomogenous.go`](./internal/solutions/counthomogenous.go)                                                             |
| 1814 | [Count Nice Pairs in an Array](https://leetcode.com/problems/count-nice-pairs-in-an-array/)                                                                 | [`countnicepairs.go`](./internal/solutions/countnicepairs.go)                                                               |
| 1822 | [Sign of the Product of an Array](https://leetcode.com/problems/sign-of-the-product-of-an-array/)                                                           | [`arraysign.go`](./internal/solutions/arraysign.go)                                                                         |
| 1838 | [Frequency of the Most Frequent Element](https://leetcode.com/problems/frequency-of-the-most-frequent-element/)                                             | [`maxfrequency.go`](./internal/solutions/maxfrequency.go)                                                                   |
| 1845 | [Seat Reservation Manager](https://leetcode.com/problems/seat-reservation-manager/)                                                                         | [`seatmanager.go`](./internal/solutions/seatmanager.go)                                                                     |
| 1846 | [Maximum Element After Decreasing and Rearranging](https://leetcode.com/problems/maximum-element-after-decreasing-and-rearranging/)                         | [`maximumelementafterdecrementingandrearranging.go`](./internal/solutions/maximumelementafterdecrementingandrearranging.go) |
| 1877 | [Minimize Maximum Pair Sum in Array](https://leetcode.com/problems/minimize-maximum-pair-sum-in-array/)                                                     | [`minpairsum.go`](./internal/solutions/minpairsum.go)                                                                       |
| 1887 | [Reduction Operations to Make the Array Elements Equal](https://leetcode.com/problems/reduction-operations-to-make-the-array-elements-equal/)               | [`reductionoperations.go`](./internal/solutions/reductionoperations.go)                                                     |
| 1897 | [Redistribute Characters to Make All Strings Equal](https://leetcode.com/problems/redistribute-characters-to-make-all-strings-equal/)                       | [`makeequal.go`](./internal/solutions/makeequal.go)                                                                         |
| 1903 | [Largest Odd Number in String](https://leetcode.com/problems/largest-odd-number-in-string/)                                                                 | [`largestoddnumber.go`](./internal/solutions/largestoddnumber.go)                                                           |
| 1913 | [Maximum Product Difference Between Two Pairs](https://leetcode.com/problems/maximum-product-difference-between-two-pairs/)                                 | [`maxproductdifference.go`](./internal/solutions/maxproductdifference.go)                                                   |
| 1921 | [Eliminate Maximum Number of Monsters](https://leetcode.com/problems/eliminate-maximum-number-of-monsters/)                                                 | [`eliminatemaximum.go`](./internal/solutions/eliminatemaximum.go)                                                           |
| 1930 | [Unique Length-3 Palindromic Subsequences](https://leetcode.com/problems/unique-length-3-palindromic-subsequences/)                                         | [`countpalindromicsubsequence.go`](./internal/solutions/countpalindromicsubsequence.go)                                     |
| 1980 | [Find Unique Binary String](https://leetcode.com/problems/find-unique-binary-string/)                                                                       | [`finddifferentbinarystring.go`](./internal/solutions/finddifferentbinarystring.go)                                         |
| 2125 | [Number of Laser Beams in a Bank](https://leetcode.com/problems/number-of-laser-beams-in-a-bank/)                                                           | [`numberofbeams.go`](./internal/solutions/numberofbeams.go)                                                                 |
| 2244 | [Minimum Rounds to Complete All Tasks](https://leetcode.com/problems/minimum-rounds-to-complete-all-tasks/)                                                 | [`minimumrounds.go`](./internal/solutions/minimumrounds.go)                                                                 |
| 2259 | [Remove Digit From Number to Maximize Result](https://leetcode.com/problems/remove-digit-from-number-to-maximize-result/)                                   | [`removedigit.go`](./internal/solutions/removedigit.go)                                                                     |
| 2264 | [Largest 3-Same-Digit Number in String](https://leetcode.com/problems/largest-3-same-digit-number-in-string/)                                               | [`largestgoodinteger.go`](./internal/solutions/largestgoodinteger.go)                                                       |
| 2265 | [Count Nodes Equal to Average of Subtree](https://leetcode.com/problems/count-nodes-equal-to-average-of-subtree/)                                           | [`averageofsubtree.go`](./internal/solutions/averageofsubtree.go)                                                           |
| 2353 | [Design a Food Rating System](https://leetcode.com/problems/design-a-food-rating-system/)                                                                   | [`foodratings.go`](./internal/solutions/foodratings.go)                                                                     |
| 2391 | [Minimum Amount of Time to Collect Garbage](https://leetcode.com/problems/minimum-amount-of-time-to-collect-garbage/)                                       | [`garbagecollection.go`](./internal/solutions/garbagecollection.go)                                                         |
| 2433 | [Find The Original Array of Prefix Xor](https://leetcode.com/problems/find-the-original-array-of-prefix-xor/)                                               | [`findarray.go`](./internal/solutions/findarray.go)                                                                         |
| 2482 | [Difference Between Ones and Zeros in Row and Column](https://leetcode.com/problems/difference-between-ones-and-zeros-in-row-and-column/)                   | [`onesminuszeros.go`](./internal/solutions/onesminuszeros.go)                                                               |
| 2610 | [Convert an Array Into a 2D Array With Conditions](https://leetcode.com/problems/convert-an-array-into-a-2d-array-with-conditions/)                         | [`findmatrix.go`](./internal/solutions/findmatrix.go)                                                                       |
| 2706 | [Buy Two Chocolates](https://leetcode.com/problems/buy-two-chocolates/)                                                                                     | [`buychoco.go`](./internal/solutions/buychoco.go)                                                                           |
| 2785 | [Sort Vowels in a String](https://leetcode.com/problems/sort-vowels-in-a-string/)                                                                           | [`sortvowels.go`](./internal/solutions/sortvowels.go)                                                                       |
| 2849 | [Determine if a Cell Is Reachable at a Given Time](https://leetcode.com/problems/determine-if-a-cell-is-reachable-at-a-given-time/)                         | [`isreachableattime.go`](./internal/solutions/isreachableattime.go)                                                         |
| 2870 | [Minimum Number of Operations to Make Array Empty](https://leetcode.com/problems/minimum-number-of-operations-to-make-array-empty/)                         | [`minoperations_2.go`](./internal/solutions/minoperations_2.go)                                                             |

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
113. Min Cost Climbing Stairs
114. Largest Number At Least Twice of Others
115. Jewels and Stones
116. Minimum Distance Between BST Nodes
117. Custom Sort String
118. Bus Routes
119. Binary Trees With Factors
120. Peak Index in a Mountain Array
121. Leaf-Similar Trees
122. Middle of the Linked List
123. Online Stock Span
124. Reverse Only Letters
125. Binary Subarrays With Sum
126. Number of Recent Calls
127. Range Sum of BST
128. Validate Stack Sequences
129. Squares of a Sorted Array
130. Max Consecutive Ones III
131. Maximum Difference Between Node and Ancestor
132. Longest Arithmetic Subsequence
133. Remove All Adjacent Duplicates In String
134. Greatest Common Divisor of Strings
135. Find in Mountain Array
136. Print in Order
137. Largest Unique Number
138. N-th Tribonacci Number
139. Snapshot Array
140. Maximum Level Sum of a Binary Tree
141. Maximum Number of Balloons
142. Unique Number of Occurrences
143. Get Equal Substrings Within Budget
144. Count Vowels Permutation
145. Check If It Is a Straight Line
146. Count Number of Nice Subarrays
147. Convert Binary Number in a Linked List to Integer
148. Deepest Leaves Sum
149. Number of Steps to Reduce a Number to Zero
150. Count Negative Numbers in a Sorted Matrix
151. Longest ZigZag Path in a Binary Tree
152. Find Lucky Integer in an Array
153. Minimum Value to Get Positive Step by Step Sum
154. Build Array Where You Can Find The Maximum Exactly K Comparisons
155. Constrained Subsequence Sum
156. Counting Elements
157. Kids With the Greatest Number of Candies
158. Longest Continuous Subarray With Absolute Diff Less Than or Equal to Limit
159. Count Good Nodes in Binary Tree
160. Maximum Number of Vowels in a Substring of Given Length
161. Max Dot Product of Two Subsequences
162. Running Sum of 1d Array
163. Longest Subarray of 1's After Deleting One Element
164. Can Make Arithmetic Progression From Sequence
165. Number of Good Pairs
166. Make The String Great
167. Determine if Two Strings Are Close
168. Richest Customer Wealth
169. Max Number of K-Sum Pairs
170. Maximum Erasure Value
171. Swapping Nodes in a Linked List
172. Find the Highest Altitude
173. Sum of Unique Elements
174. Merge Strings Alternately
175. Maximum Score of a Good Subarray
176. Check if the Sentence Is Pangram
177. Check if All Characters Have Equal Number of Occurrences
178. Reverse Prefix of Word
179. Minimum Number of Operations to Make Array Continuous
180. Remove Colored Pieces if Both Neighbors are the Same Color
181. Parallel Courses III
182. Reverse Nodes in Even Length Groups
183. K Radius Subarray Averages
184. Delete the Middle Node of a Linked List
185. Maximum Twin Sum of a Linked List
186. Find the Difference of Two Arrays
187. Find Players With Zero or One Losses
188. Add Two Integers
189. Intersection of Multiple Arrays
190. Number of Flowers in Full Bloom
191. Minimum Consecutive Cards to Pick Up
192. Number of Ways to Split Array
193. Successful Pairs of Spells and Potions
194. Max Sum of a Pair With Equal Sum of Digits
195. First Letter to Appear Twice
196. Equal Row and Column Pairs
197. Removing Stars From a String
198. Using a Robot to Print the Lexicographically Smallest String
199. Total Cost to Hire K Workers
200. Design Graph With Shortest Path Calculator
201. Painting the Walls
202. Largest Submatrix With Rearrangements
203. Number of Ways to Divide a Long Corridor
204. Minimum One Bit Operations to Make Integers Zero
205. Decode Ways
206. String Compression II
207. Minimum Difficulty of a Job Schedule
208. Maximum Profit in Job Scheduling
209. Arithmetic Slices II - Subsequence
