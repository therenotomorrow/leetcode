LeetCode
========

[LeetCode problems](https://leetcode.com/problemset/all/) and how
I [solved them](https://leetcode.com/therenotomorrow/).

Structure
---------

- [Explore](./explore) - useful technics shortcuts for learn
- [Solutions](./solutions) - all solved and refactored problems
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
./scripts/test.sh
```

Completed explores
------------------

* [The LeetCode Beginner's Guide](https://leetcode.com/explore/featured/card/the-leetcode-beginners-guide/)
* [Array and String. Introduction to Data Structure](https://leetcode.com/explore/learn/card/array-and-string/)
* [Top Interview Questions. Easy Collection](https://leetcode.com/explore/featured/card/top-interview-questions-easy/)

Completed problems
------------------

|  #   | Problem                                                                                                                                                                     | Solution                                                                                                                |
|:----:|:----------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:------------------------------------------------------------------------------------------------------------------------|
|  1   | [Two Sum](https://leetcode.com/problems/two-sum/description/)                                                                                                               | [`twosum_1.go`](solutions/golang/twosum_1.go)                                                                           |
|  2   | [Add Two Numbers](https://leetcode.com/problems/add-two-numbers/description/)                                                                                               | [`addtwonumbers.go`](solutions/golang/addtwonumbers.go)                                                                 |
|  3   | [Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/description/)                                 | [`lengthoflongestsubstring.go`](solutions/golang/lengthoflongestsubstring.go)                                           |
|  4   | [Median of Two Sorted Arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/description/)                                                                       | [`findmediansortedarrays.go`](solutions/golang/findmediansortedarrays.go)                                               |
|  7   | [Reverse Integer](https://leetcode.com/problems/reverse-integer/description/)                                                                                               | [`reverse.go`](solutions/golang/reverse.go)                                                                             |
|  9   | [Palindrome Number](https://leetcode.com/problems/palindrome-number/description/)                                                                                           | [`ispalindrome.go`](solutions/golang/ispalindrome.go)                                                                   |
|  19  | [Remove Nth Node From End of List](https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/)                                                             | [`removenthfromend.go`](solutions/golang/removenthfromend.go)                                                           |
|  35  | [Search Insert Position](https://leetcode.com/problems/search-insert-position/description/)                                                                                 | [`searchinsert.go`](solutions/golang/searchinsert.go)                                                                   |
|  49  | [Group Anagrams](https://leetcode.com/problems/group-anagrams/description/)                                                                                                 | [`groupanagrams.go`](solutions/golang/groupanagrams.go)                                                                 |
|  53  | [Maximum Subarray](https://leetcode.com/problems/maximum-subarray/description/)                                                                                             | [`maxsubarray.go`](solutions/golang/maxsubarray.go)                                                                     |
|  57  | [Insert Interval](https://leetcode.com/problems/insert-interval/description/)                                                                                               | [`insert.go`](solutions/golang/insert.go)                                                                               |
|  58  | [Length of Last Word](https://leetcode.com/problems/length-of-last-word/description/)                                                                                       | [`lengthoflastword.go`](solutions/golang/lengthoflastword.go)                                                           |
|  62  | [Unique Paths](https://leetcode.com/problems/unique-paths/description/)                                                                                                     | [`uniquepaths.go`](solutions/golang/uniquepaths.go)                                                                     |
|  69  | [Sqrt(x)](https://leetcode.com/problems/sqrtx/description/)                                                                                                                 | [`mysqrt.go`](solutions/golang/mysqrt.go)                                                                               |
|  70  | [Climbing Stairs](https://leetcode.com/problems/climbing-stairs/description/)                                                                                               | [`climbstairs.go`](solutions/golang/climbstairs.go)                                                                     |
|  76  | [Minimum Window Substring](https://leetcode.com/problems/minimum-window-substring/description/)                                                                             | [`minwindow.go`](solutions/golang/minwindow.go)                                                                         |
|  94  | [Binary Tree Inorder Traversal](https://leetcode.com/problems/binary-tree-inorder-traversal/description/)                                                                   | [`inordertraversal.go`](solutions/golang/inordertraversal.go)                                                           |
| 100  | [Same Tree](https://leetcode.com/problems/same-tree/description/)                                                                                                           | [`issametree.go`](solutions/golang/issametree.go)                                                                       |
| 110  | [Balanced Binary Tree](https://leetcode.com/problems/balanced-binary-tree/description/)                                                                                     | [`isbalanced.go`](solutions/golang/isbalanced.go)                                                                       |
| 118  | [Pascal's Triangle](https://leetcode.com/problems/pascals-triangle/description/)                                                                                            | [`generate.go`](solutions/golang/generate.go)                                                                           |
| 119  | [Pascal's Triangle II](https://leetcode.com/problems/pascals-triangle-ii/description/)                                                                                      | [`getrow.go`](solutions/golang/getrow.go)                                                                               |
| 141  | [Linked List Cycle](https://leetcode.com/problems/linked-list-cycle/description/)                                                                                           | [`hascycle.go`](solutions/golang/hascycle.go)                                                                           |
| 144  | [Binary Tree Preorder Traversal](https://leetcode.com/problems/binary-tree-preorder-traversal/description/)                                                                 | [`preordertraversal.go`](solutions/golang/preordertraversal.go)                                                         |
| 145  | [Binary Tree Postorder Traversal](https://leetcode.com/problems/binary-tree-postorder-traversal/description/)                                                               | [`postordertraversal.go`](solutions/golang/postordertraversal.go)                                                       |
| 150  | [Evaluate Reverse Polish Notation](https://leetcode.com/problems/evaluate-reverse-polish-notation/description/)                                                             | [`evalrpn.go`](solutions/golang/evalrpn.go)                                                                             |
| 157  | [Read N Characters Given Read4](https://leetcode.com/problems/read-n-characters-given-read4/description/)                                                                   | [`solution.go`](solutions/golang/solution.go)                                                                           |
| 159  | [Longest Substring with At Most Two Distinct Characters](https://leetcode.com/problems/longest-substring-with-at-most-two-distinct-characters/description/)                 | [`lengthoflongestsubstringtwodistinct.go`](solutions/golang/lengthoflongestsubstringtwodistinct.go)                     |
| 160  | [Intersection of Two Linked Lists](https://leetcode.com/problems/intersection-of-two-linked-lists/description/)                                                             | [`getintersectionnode.go`](solutions/golang/getintersectionnode.go)                                                     |
| 163  | [Missing Ranges](https://leetcode.com/problems/missing-ranges/description/)                                                                                                 | [`findmissingranges.go`](solutions/golang/findmissingranges.go)                                                         |
| 168  | [Excel Sheet Column Title](https://leetcode.com/problems/excel-sheet-column-title/description/)                                                                             | [`converttotitle.go`](solutions/golang/converttotitle.go)                                                               |
| 169  | [Majority Element](https://leetcode.com/problems/majority-element/description/)                                                                                             | [`majorityelement.go`](solutions/golang/majorityelement.go)                                                             |
| 171  | [Excel Sheet Column Number](https://leetcode.com/problems/excel-sheet-column-number/description/)                                                                           | [`titletonumber.go`](solutions/golang/titletonumber.go)                                                                 |
| 175  | [Combine Two Tables](https://leetcode.com/problems/combine-two-tables/description/)                                                                                         | [`175.sql`](solutions/sql/175.sql)                                                                                      |
| 181  | [Employees Earning More Than Their Managers](https://leetcode.com/problems/employees-earning-more-than-their-managers/description/)                                         | [`181.sql`](solutions/sql/181.sql)                                                                                      |
| 182  | [Duplicate Emails](https://leetcode.com/problems/duplicate-emails/description/)                                                                                             | [`182.sql`](solutions/sql/182.sql)                                                                                      |
| 183  | [Customers Who Never Order](https://leetcode.com/problems/customers-who-never-order/description/)                                                                           | [`183.sql`](solutions/sql/183.sql)                                                                                      |
| 191  | [Number of 1 Bits](https://leetcode.com/problems/number-of-1-bits/description/)                                                                                             | [`hammingweight.go`](solutions/golang/hammingweight.go)                                                                 |
| 193  | [Valid Phone Numbers](https://leetcode.com/problems/valid-phone-numbers/description/)                                                                                       | [`193.sh`](solutions/bash/193.sh)                                                                                       |
| 195  | [Tenth Line](https://leetcode.com/problems/tenth-line/description/)                                                                                                         | [`195.sh`](solutions/bash/195.sh)                                                                                       |
| 196  | [Delete Duplicate Emails](https://leetcode.com/problems/delete-duplicate-emails/description/)                                                                               | [`196.sql`](solutions/sql/196.sql)                                                                                      |
| 197  | [Rising Temperature](https://leetcode.com/problems/rising-temperature/description/)                                                                                         | [`197.sql`](solutions/sql/197.sql)                                                                                      |
| 198  | [House Robber](https://leetcode.com/problems/house-robber/description/)                                                                                                     | [`rob.go`](solutions/golang/rob.go)                                                                                     |
| 201  | [Bitwise AND of Numbers Range](https://leetcode.com/problems/bitwise-and-of-numbers-range/description/)                                                                     | [`rangebitwiseand.go`](solutions/golang/rangebitwiseand.go)                                                             |
| 202  | [Happy Number](https://leetcode.com/problems/happy-number/description/)                                                                                                     | [`ishappy.go`](solutions/golang/ishappy.go)                                                                             |
| 209  | [Minimum Size Subarray Sum](https://leetcode.com/problems/minimum-size-subarray-sum/description/)                                                                           | [`minsubarraylen.go`](solutions/golang/minsubarraylen.go)                                                               |
| 219  | [Contains Duplicate II](https://leetcode.com/problems/contains-duplicate-ii/description/)                                                                                   | [`containsnearbyduplicate.go`](solutions/golang/containsnearbyduplicate.go)                                             |
| 222  | [Count Complete Tree Nodes](https://leetcode.com/problems/count-complete-tree-nodes/description/)                                                                           | [`countnodes.go`](solutions/golang/countnodes.go)                                                                       |
| 226  | [Invert Binary Tree](https://leetcode.com/problems/invert-binary-tree/description/)                                                                                         | [`inverttree.go`](solutions/golang/inverttree.go)                                                                       |
| 231  | [Power of Two](https://leetcode.com/problems/power-of-two/description/)                                                                                                     | [`ispoweroftwo.go`](solutions/golang/ispoweroftwo.go)                                                                   |
| 232  | [Implement Queue using Stacks](https://leetcode.com/problems/implement-queue-using-stacks/description/)                                                                     | [`myqueue.go`](solutions/golang/myqueue.go)                                                                             |
| 238  | [Product of Array Except Self](https://leetcode.com/problems/product-of-array-except-self/description/)                                                                     | [`productexceptself.go`](solutions/golang/productexceptself.go)                                                         |
| 242  | [Valid Anagram](https://leetcode.com/problems/valid-anagram/description/)                                                                                                   | [`isanagram.go`](solutions/golang/isanagram.go)                                                                         |
| 243  | [Shortest Word Distance](https://leetcode.com/problems/shortest-word-distance/description/)                                                                                 | [`shortestdistance.go`](solutions/golang/shortestdistance.go)                                                           |
| 246  | [Strobogrammatic Number](https://leetcode.com/problems/strobogrammatic-number/description/)                                                                                 | [`isstrobogrammatic.go`](solutions/golang/isstrobogrammatic.go)                                                         |
| 252  | [Meeting Rooms](https://leetcode.com/problems/meeting-rooms/description/)                                                                                                   | [`canattendmeetings.go`](solutions/golang/canattendmeetings.go)                                                         |
| 253  | [Meeting Rooms II](https://leetcode.com/problems/meeting-rooms-ii/description/)                                                                                             | [`minmeetingrooms.go`](solutions/golang/minmeetingrooms.go)                                                             |
| 256  | [Paint House](https://leetcode.com/problems/paint-house/description/)                                                                                                       | [`mincost_2.go`](solutions/golang/mincost_2.go)                                                                         |
| 257  | [Binary Tree Paths](https://leetcode.com/problems/binary-tree-paths/description/)                                                                                           | [`binarytreepaths.go`](solutions/golang/binarytreepaths.go)                                                             |
| 258  | [Add Digits](https://leetcode.com/problems/add-digits/description/)                                                                                                         | [`adddigits.go`](solutions/golang/adddigits.go)                                                                         |
| 266  | [Palindrome Permutation](https://leetcode.com/problems/palindrome-permutation/description/)                                                                                 | [`canpermutepalindrome.go`](solutions/golang/canpermutepalindrome.go)                                                   |
| 268  | [Missing Number](https://leetcode.com/problems/missing-number/description/)                                                                                                 | [`missingnumber.go`](solutions/golang/missingnumber.go)                                                                 |
| 276  | [Paint Fence](https://leetcode.com/problems/paint-fence/description/)                                                                                                       | [`numways_2.go`](solutions/golang/numways_2.go)                                                                         |
| 279  | [Perfect Squares](https://leetcode.com/problems/perfect-squares/description/)                                                                                               | [`numsquares.go`](solutions/golang/numsquares.go)                                                                       |
| 283  | [Move Zeroes](https://leetcode.com/problems/move-zeroes/description/)                                                                                                       | [`movezeroes.go`](solutions/golang/movezeroes.go)                                                                       |
| 293  | [Flip Game](https://leetcode.com/problems/flip-game/description/)                                                                                                           | [`generatepossiblenextmoves.go`](solutions/golang/generatepossiblenextmoves.go)                                         |
| 296  | [Best Meeting Point](https://leetcode.com/problems/best-meeting-point/description/)                                                                                         | [`mintotaldistance.go`](solutions/golang/mintotaldistance.go)                                                           |
| 300  | [Longest Increasing Subsequence](https://leetcode.com/problems/longest-increasing-subsequence/description/)                                                                 | [`lengthoflis.go`](solutions/golang/lengthoflis.go)                                                                     |
| 303  | [Range Sum Query - Immutable](https://leetcode.com/problems/range-sum-query-immutable/description/)                                                                         | [`numarray.go`](solutions/golang/numarray.go)                                                                           |
| 342  | [Power of Four](https://leetcode.com/problems/power-of-four/description/)                                                                                                   | [`ispoweroffour.go`](solutions/golang/ispoweroffour.go)                                                                 |
| 344  | [Reverse String](https://leetcode.com/problems/reverse-string/description/)                                                                                                 | [`reversestring.go`](solutions/golang/reversestring.go)                                                                 |
| 346  | [Moving Average from Data Stream](https://leetcode.com/problems/moving-average-from-data-stream/description/)                                                               | [`movingaverage.go`](solutions/golang/movingaverage.go)                                                                 |
| 349  | [Intersection of Two Arrays](https://leetcode.com/problems/intersection-of-two-arrays/description/)                                                                         | [`intersection.go`](solutions/golang/intersection.go)                                                                   |
| 359  | [Logger Rate Limiter](https://leetcode.com/problems/logger-rate-limiter/description/)                                                                                       | [`logger.go`](solutions/golang/logger.go)                                                                               |
| 367  | [Valid Perfect Square](https://leetcode.com/problems/valid-perfect-square/description/)                                                                                     | [`isperfectsquare.go`](solutions/golang/isperfectsquare.go)                                                             |
| 368  | [Largest Divisible Subset](https://leetcode.com/problems/largest-divisible-subset/description/)                                                                             | [`largestdivisiblesubset.go`](solutions/golang/largestdivisiblesubset.go)                                               |
| 380  | [Insert Delete GetRandom O(1)](https://leetcode.com/problems/insert-delete-getrandom-o1/description/)                                                                       | [`randomizedset.go`](solutions/golang/randomizedset.go)                                                                 |
| 387  | [First Unique Character in a String](https://leetcode.com/problems/first-unique-character-in-a-string/description/)                                                         | [`firstuniqchar.go`](solutions/golang/firstuniqchar.go)                                                                 |
| 404  | [Sum of Left Leaves](https://leetcode.com/problems/sum-of-left-leaves/description/)                                                                                         | [`sumofleftleaves.go`](solutions/golang/sumofleftleaves.go)                                                             |
| 405  | [Convert a Number to Hexadecimal](https://leetcode.com/problems/convert-a-number-to-hexadecimal/description/)                                                               | [`tohex.go`](solutions/golang/tohex.go)                                                                                 |
| 408  | [Valid Word Abbreviation](https://leetcode.com/problems/valid-word-abbreviation/description/)                                                                               | [`validwordabbreviation.go`](solutions/golang/validwordabbreviation.go)                                                 |
| 451  | [Sort Characters By Frequency](https://leetcode.com/problems/sort-characters-by-frequency/description/)                                                                     | [`frequencysort.go`](solutions/golang/frequencysort.go)                                                                 |
| 452  | [Minimum Number of Arrows to Burst Balloons](https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/description/)                                         | [`findminarrowshots.go`](solutions/golang/findminarrowshots.go)                                                         |
| 455  | [Assign Cookies](https://leetcode.com/problems/assign-cookies/description/)                                                                                                 | [`findcontentchildren.go`](solutions/golang/findcontentchildren.go)                                                     |
| 458  | [Poor Pigs](https://leetcode.com/problems/poor-pigs/description/)                                                                                                           | [`poorpigs.go`](solutions/golang/poorpigs.go)                                                                           |
| 485  | [Max Consecutive Ones](https://leetcode.com/problems/max-consecutive-ones/description/)                                                                                     | [`findmaxconsecutiveones_1.go`](solutions/golang/findmaxconsecutiveones_1.go)                                           |
| 487  | [Max Consecutive Ones II](https://leetcode.com/problems/max-consecutive-ones-ii/description/)                                                                               | [`findmaxconsecutiveones_2.go`](solutions/golang/findmaxconsecutiveones_2.go)                                           |
| 501  | [Find Mode in Binary Search Tree](https://leetcode.com/problems/find-mode-in-binary-search-tree/description/)                                                               | [`findmode.go`](solutions/golang/findmode.go)                                                                           |
| 509  | [Fibonacci Number](https://leetcode.com/problems/fibonacci-number/description/)                                                                                             | [`fib.go`](solutions/golang/fib.go)                                                                                     |
| 513  | [Find Bottom Left Tree Value](https://leetcode.com/problems/find-bottom-left-tree-value/description/)                                                                       | [`findbottomleftvalue.go`](solutions/golang/findbottomleftvalue.go)                                                     |
| 525  | [Contiguous Array](https://leetcode.com/problems/contiguous-array/description/)                                                                                             | [`findmaxlength.go`](solutions/golang/findmaxlength.go)                                                                 |
| 539  | [Minimum Time Difference](https://leetcode.com/problems/minimum-time-difference/description/)                                                                               | [`findmindifference.go`](solutions/golang/findmindifference.go)                                                         |
| 543  | [Diameter of Binary Tree](https://leetcode.com/problems/diameter-of-binary-tree/description/)                                                                               | [`diameterofbinarytree.go`](solutions/golang/diameterofbinarytree.go)                                                   |
| 557  | [Reverse Words in a String III](https://leetcode.com/problems/reverse-words-in-a-string-iii/description/)                                                                   | [`reversewords.go`](solutions/golang/reversewords.go)                                                                   |
| 572  | [Subtree of Another Tree](https://leetcode.com/problems/subtree-of-another-tree/description/)                                                                               | [`issubtree.go`](solutions/golang/issubtree.go)                                                                         |
| 573  | [Squirrel Simulation](https://leetcode.com/problems/squirrel-simulation/description/)                                                                                       | [`mindistance.go`](solutions/golang/mindistance.go)                                                                     |
| 576  | [Out of Boundary Paths](https://leetcode.com/problems/out-of-boundary-paths/description/)                                                                                   | [`findpaths.go`](solutions/golang/findpaths.go)                                                                         |
| 606  | [Construct String from Binary Tree](https://leetcode.com/problems/construct-string-from-binary-tree/description/)                                                           | [`tree2str.go`](solutions/golang/tree2str.go)                                                                           |
| 624  | [Maximum Distance in Arrays](https://leetcode.com/problems/maximum-distance-in-arrays/description/)                                                                         | [`maxdistance.go`](solutions/golang/maxdistance.go)                                                                     |
| 643  | [Maximum Average Subarray I](https://leetcode.com/problems/maximum-average-subarray-i/description/)                                                                         | [`findmaxaverage.go`](solutions/golang/findmaxaverage.go)                                                               |
| 645  | [Set Mismatch](https://leetcode.com/problems/set-mismatch/description/)                                                                                                     | [`finderrornums.go`](solutions/golang/finderrornums.go)                                                                 |
| 647  | [Palindromic Substrings](https://leetcode.com/problems/palindromic-substrings/description/)                                                                                 | [`countsubstrings.go`](solutions/golang/countsubstrings.go)                                                             |
| 661  | [Image Smoother](https://leetcode.com/problems/image-smoother/description/)                                                                                                 | [`imagesmoother.go`](solutions/golang/imagesmoother.go)                                                                 |
| 680  | [Valid Palindrome II](https://leetcode.com/problems/valid-palindrome-ii/description/)                                                                                       | [`validpalindrome.go`](solutions/golang/validpalindrome.go)                                                             |
| 704  | [Binary Search](https://leetcode.com/problems/binary-search/description/)                                                                                                   | [`search.go`](solutions/golang/search.go)                                                                               |
| 724  | [Find Pivot Index](https://leetcode.com/problems/find-pivot-index/description/)                                                                                             | [`pivotindex.go`](solutions/golang/pivotindex.go)                                                                       |
| 739  | [Daily Temperatures](https://leetcode.com/problems/daily-temperatures/description/)                                                                                         | [`dailytemperatures.go`](solutions/golang/dailytemperatures.go)                                                         |
| 746  | [Min Cost Climbing Stairs](https://leetcode.com/problems/min-cost-climbing-stairs/description/)                                                                             | [`mincostclimbingstairs.go`](solutions/golang/mincostclimbingstairs.go)                                                 |
| 779  | [K-th Symbol in Grammar](https://leetcode.com/problems/k-th-symbol-in-grammar/description/)                                                                                 | [`kthgrammar.go`](solutions/golang/kthgrammar.go)                                                                       |
| 791  | [Custom Sort String](https://leetcode.com/problems/custom-sort-string/description/)                                                                                         | [`customsortstring.go`](solutions/golang/customsortstring.go)                                                           |
| 844  | [Backspace String Compare](https://leetcode.com/problems/backspace-string-compare/description/)                                                                             | [`backspacecompare.go`](solutions/golang/backspacecompare.go)                                                           |
| 867  | [Transpose Matrix](https://leetcode.com/problems/transpose-matrix/description/)                                                                                             | [`transpose.go`](solutions/golang/transpose.go)                                                                         |
| 872  | [Leaf-Similar Trees](https://leetcode.com/problems/leaf-similar-trees/description/)                                                                                         | [`leafsimilar.go`](solutions/golang/leafsimilar.go)                                                                     |
| 876  | [Middle of the Linked List](https://leetcode.com/problems/middle-of-the-linked-list/description/)                                                                           | [`middlenode.go`](solutions/golang/middlenode.go)                                                                       |
| 905  | [Sort Array By Parity](https://leetcode.com/problems/sort-array-by-parity/description/)                                                                                     | [`sortarraybyparity.go`](solutions/golang/sortarraybyparity.go)                                                         |
| 907  | [Sum of Subarray Minimums](https://leetcode.com/problems/sum-of-subarray-minimums/description/)                                                                             | [`sumsubarraymins.go`](solutions/golang/sumsubarraymins.go)                                                             |
| 917  | [Reverse Only Letters](https://leetcode.com/problems/reverse-only-letters/description/)                                                                                     | [`reverseonlyletters.go`](solutions/golang/reverseonlyletters.go)                                                       |
| 930  | [Binary Subarrays With Sum](https://leetcode.com/problems/binary-subarrays-with-sum/description/)                                                                           | [`numsubarrayswithsum.go`](solutions/golang/numsubarrayswithsum.go)                                                     |
| 931  | [Minimum Falling Path Sum](https://leetcode.com/problems/minimum-falling-path-sum/description/)                                                                             | [`minfallingpathsum.go`](solutions/golang/minfallingpathsum.go)                                                         |
| 935  | [Knight Dialer](https://leetcode.com/problems/knight-dialer/description/)                                                                                                   | [`knightdialer.go`](solutions/golang/knightdialer.go)                                                                   |
| 938  | [Range Sum of BST](https://leetcode.com/problems/range-sum-of-bst/description/)                                                                                             | [`rangesumbst.go`](solutions/golang/rangesumbst.go)                                                                     |
| 948  | [Bag of Tokens](https://leetcode.com/problems/bag-of-tokens/description/)                                                                                                   | [`bagoftokensscore.go`](solutions/golang/bagoftokensscore.go)                                                           |
| 977  | [Squares of a Sorted Array](https://leetcode.com/problems/squares-of-a-sorted-array/description/)                                                                           | [`sortedsquares.go`](solutions/golang/sortedsquares.go)                                                                 |
| 997  | [Find the Town Judge](https://leetcode.com/problems/find-the-town-judge/description/)                                                                                       | [`findjudge.go`](solutions/golang/findjudge.go)                                                                         |
| 1002 | [Find Common Characters](https://leetcode.com/problems/find-common-characters/description/)                                                                                 | [`commonchars.go`](solutions/golang/commonchars.go)                                                                     |
| 1004 | [Max Consecutive Ones III](https://leetcode.com/problems/max-consecutive-ones-iii/description/)                                                                             | [`longestones.go`](solutions/golang/longestones.go)                                                                     |
| 1018 | [Binary Prefix Divisible By 5](https://leetcode.com/problems/binary-prefix-divisible-by-5/description/)                                                                     | [`prefixesdivby5.go`](solutions/golang/prefixesdivby5.go)                                                               |
| 1026 | [Maximum Difference Between Node and Ancestor](https://leetcode.com/problems/maximum-difference-between-node-and-ancestor/description/)                                     | [`maxancestordiff.go`](solutions/golang/maxancestordiff.go)                                                             |
| 1051 | [Height Checker](https://leetcode.com/problems/height-checker/description/)                                                                                                 | [`heightchecker.go`](solutions/golang/heightchecker.go)                                                                 |
| 1057 | [Campus Bikes](https://leetcode.com/problems/campus-bikes/description/)                                                                                                     | [`assignbikes_1.go`](solutions/golang/assignbikes_1.go)                                                                 |
| 1066 | [Campus Bikes II](https://leetcode.com/problems/campus-bikes-ii/description/)                                                                                               | [`assignbikes_2.go`](solutions/golang/assignbikes_2.go)                                                                 |
| 1085 | [Sum of Digits in the Minimum Number](https://leetcode.com/problems/sum-of-digits-in-the-minimum-number/description/)                                                       | [`sumofdigits.go`](solutions/golang/sumofdigits.go)                                                                     |
| 1099 | [Two Sum Less Than K](https://leetcode.com/problems/two-sum-less-than-k/description/)                                                                                       | [`twosumlessthank.go`](solutions/golang/twosumlessthank.go)                                                             |
| 1119 | [Remove Vowels from a String](https://leetcode.com/problems/remove-vowels-from-a-string/description/)                                                                       | [`removevowels.go`](solutions/golang/removevowels.go)                                                                   |
| 1120 | [Maximum Average Subtree](https://leetcode.com/problems/maximum-average-subtree/description/)                                                                               | [`maximumaveragesubtree.go`](solutions/golang/maximumaveragesubtree.go)                                                 |
| 1137 | [N-th Tribonacci Number](https://leetcode.com/problems/n-th-tribonacci-number/description/)                                                                                 | [`tribonacci.go`](solutions/golang/tribonacci.go)                                                                       |
| 1143 | [Longest Common Subsequence](https://leetcode.com/problems/longest-common-subsequence/description/)                                                                         | [`longestcommonsubsequence.go`](solutions/golang/longestcommonsubsequence.go)                                           |
| 1155 | [Number of Dice Rolls With Target Sum](https://leetcode.com/problems/number-of-dice-rolls-with-target-sum/description/)                                                     | [`numrollstotarget.go`](solutions/golang/numrollstotarget.go)                                                           |
| 1160 | [Find Words That Can Be Formed by Characters](https://leetcode.com/problems/find-words-that-can-be-formed-by-characters/description/)                                       | [`countcharacters.go`](solutions/golang/countcharacters.go)                                                             |
| 1165 | [Single-Row Keyboard](https://leetcode.com/problems/single-row-keyboard/description/)                                                                                       | [`calculatetime.go`](solutions/golang/calculatetime.go)                                                                 |
| 1167 | [Minimum Cost to Connect Sticks](https://leetcode.com/problems/minimum-cost-to-connect-sticks/description/)                                                                 | [`connectsticks.go`](solutions/golang/connectsticks.go)                                                                 |
| 1171 | [Remove Zero Sum Consecutive Nodes from Linked List](https://leetcode.com/problems/remove-zero-sum-consecutive-nodes-from-linked-list/description/)                         | [`removezerosumsublists.go`](solutions/golang/removezerosumsublists.go)                                                 |
| 1207 | [Unique Number of Occurrences](https://leetcode.com/problems/unique-number-of-occurrences/description/)                                                                     | [`uniqueoccurrences.go`](solutions/golang/uniqueoccurrences.go)                                                         |
| 1208 | [Get Equal Substrings Within Budget](https://leetcode.com/problems/get-equal-substrings-within-budget/description/)                                                         | [`equalsubstring.go`](solutions/golang/equalsubstring.go)                                                               |
| 1216 | [Valid Palindrome III](https://leetcode.com/problems/valid-palindrome-iii/description/)                                                                                     | [`isvalidpalindrome.go`](solutions/golang/isvalidpalindrome.go)                                                         |
| 1219 | [Path with Maximum Gold](https://leetcode.com/problems/path-with-maximum-gold/description/)                                                                                 | [`getmaximumgold.go`](solutions/golang/getmaximumgold.go)                                                               |
| 1239 | [Maximum Length of a Concatenated String with Unique Characters](https://leetcode.com/problems/maximum-length-of-a-concatenated-string-with-unique-characters/description/) | [`maxlength.go`](solutions/golang/maxlength.go)                                                                         |
| 1243 | [Array Transformation](https://leetcode.com/problems/array-transformation/description/)                                                                                     | [`transformarray.go`](solutions/golang/transformarray.go)                                                               |
| 1266 | [Minimum Time Visiting All Points](https://leetcode.com/problems/minimum-time-visiting-all-points/description/)                                                             | [`mintimetovisitallpoints.go`](solutions/golang/mintimetovisitallpoints.go)                                             |
| 1269 | [Number of Ways to Stay in the Same Place After Some Steps](https://leetcode.com/problems/number-of-ways-to-stay-in-the-same-place-after-some-steps/description/)           | [`numways_1.go`](solutions/golang/numways_1.go)                                                                         |
| 1272 | [Remove Interval](https://leetcode.com/problems/remove-interval/description/)                                                                                               | [`removeinterval.go`](solutions/golang/removeinterval.go)                                                               |
| 1287 | [Element Appearing More Than 25% In Sorted Array](https://leetcode.com/problems/element-appearing-more-than-25-in-sorted-array/description/)                                | [`findspecialinteger.go`](solutions/golang/findspecialinteger.go)                                                       |
| 1291 | [Sequential Digits](https://leetcode.com/problems/sequential-digits/description/)                                                                                           | [`sequentialdigits.go`](solutions/golang/sequentialdigits.go)                                                           |
| 1295 | [Find Numbers with Even Number of Digits](https://leetcode.com/problems/find-numbers-with-even-number-of-digits/description/)                                               | [`findnumbers.go`](solutions/golang/findnumbers.go)                                                                     |
| 1337 | [The K Weakest Rows in a Matrix](https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix/description/)                                                                 | [`kweakestrows.go`](solutions/golang/kweakestrows.go)                                                                   |
| 1347 | [Minimum Number of Steps to Make Two Strings Anagram](https://leetcode.com/problems/minimum-number-of-steps-to-make-two-strings-anagram/description/)                       | [`minsteps.go`](solutions/golang/minsteps.go)                                                                           |
| 1356 | [Sort Integers by The Number of 1 Bits](https://leetcode.com/problems/sort-integers-by-the-number-of-1-bits/description/)                                                   | [`sortbybits.go`](solutions/golang/sortbybits.go)                                                                       |
| 1360 | [Number of Days Between Two Dates](https://leetcode.com/problems/number-of-days-between-two-dates/description/)                                                             | [`daysbetweendates.go`](solutions/golang/daysbetweendates.go)                                                           |
| 1361 | [Validate Binary Tree Nodes](https://leetcode.com/problems/validate-binary-tree-nodes/description/)                                                                         | [`validatebinarytreenodes.go`](solutions/golang/validatebinarytreenodes.go)                                             |
| 1413 | [Minimum Value to Get Positive Step by Step Sum](https://leetcode.com/problems/minimum-value-to-get-positive-step-by-step-sum/description/)                                 | [`minstartvalue.go`](solutions/golang/minstartvalue.go)                                                                 |
| 1422 | [Maximum Score After Splitting a String](https://leetcode.com/problems/maximum-score-after-splitting-a-string/description/)                                                 | [`maxscore.go`](solutions/golang/maxscore.go)                                                                           |
| 1424 | [Diagonal Traverse II](https://leetcode.com/problems/diagonal-traverse-ii/description/)                                                                                     | [`finddiagonalorder.go`](solutions/golang/finddiagonalorder.go)                                                         |
| 1428 | [Leftmost Column with at Least a One](https://leetcode.com/problems/leftmost-column-with-at-least-a-one/description/)                                                       | [`leftmostcolumnwithone.go`](solutions/golang/leftmostcolumnwithone.go)                                                 |
| 1436 | [Destination City](https://leetcode.com/problems/destination-city/description/)                                                                                             | [`destcity.go`](solutions/golang/destcity.go)                                                                           |
| 1441 | [Build an Array With Stack Operations](https://leetcode.com/problems/build-an-array-with-stack-operations/description/)                                                     | [`buildarray.go`](solutions/golang/buildarray.go)                                                                       |
| 1456 | [Maximum Number of Vowels in a Substring of Given Length](https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length/description/)               | [`maxvowels.go`](solutions/golang/maxvowels.go)                                                                         |
| 1457 | [Pseudo-Palindromic Paths in a Binary Tree](https://leetcode.com/problems/pseudo-palindromic-paths-in-a-binary-tree/description/)                                           | [`pseudopalindromicpaths.go`](solutions/golang/pseudopalindromicpaths.go)                                               |
| 1463 | [Cherry Pickup II](https://leetcode.com/problems/cherry-pickup-ii/description/)                                                                                             | [`cherrypickup.go`](solutions/golang/cherrypickup.go)                                                                   |
| 1464 | [Maximum Product of Two Elements in an Array](https://leetcode.com/problems/maximum-product-of-two-elements-in-an-array/description/)                                       | [`maxproduct.go`](solutions/golang/maxproduct.go)                                                                       |
| 1470 | [Shuffle the Array](https://leetcode.com/problems/shuffle-the-array/description/)                                                                                           | [`shuffle.go`](solutions/golang/shuffle.go)                                                                             |
| 1480 | [Running Sum of 1d Array](https://leetcode.com/problems/running-sum-of-1d-array/description/)                                                                               | [`runningsum.go`](solutions/golang/runningsum.go)                                                                       |
| 1481 | [Least Number of Unique Integers after K Removals](https://leetcode.com/problems/least-number-of-unique-integers-after-k-removals/description/)                             | [`findleastnumofuniqueints.go`](solutions/golang/findleastnumofuniqueints.go)                                           |
| 1496 | [Path Crossing](https://leetcode.com/problems/path-crossing/description/)                                                                                                   | [`ispathcrossing.go`](solutions/golang/ispathcrossing.go)                                                               |
| 1503 | [Last Moment Before All Ants Fall Out of a Plank](https://leetcode.com/problems/last-moment-before-all-ants-fall-out-of-a-plank/description/)                               | [`getlastmoment.go`](solutions/golang/getlastmoment.go)                                                                 |
| 1535 | [Find the Winner of an Array Game](https://leetcode.com/problems/find-the-winner-of-an-array-game/description/)                                                             | [`getwinner.go`](solutions/golang/getwinner.go)                                                                         |
| 1561 | [Maximum Number of Coins You Can Get](https://leetcode.com/problems/maximum-number-of-coins-you-can-get/description/)                                                       | [`maxcoins.go`](solutions/golang/maxcoins.go)                                                                           |
| 1578 | [Minimum Time to Make Rope Colorful](https://leetcode.com/problems/minimum-time-to-make-rope-colorful/description/)                                                         | [`mincost_1.go`](solutions/golang/mincost_1.go)                                                                         |
| 1581 | [Customer Who Visited but Did Not Make Any Transactions](https://leetcode.com/problems/customer-who-visited-but-did-not-make-any-transactions/description/)                 | [`1581.sql`](solutions/sql/1581.sql)                                                                                    |
| 1582 | [Special Positions in a Binary Matrix](https://leetcode.com/problems/special-positions-in-a-binary-matrix/description/)                                                     | [`numspecial.go`](solutions/golang/numspecial.go)                                                                       |
| 1603 | [Design Parking System](https://leetcode.com/problems/design-parking-system/description/)                                                                                   | [`parkingsystem.go`](solutions/golang/parkingsystem.go)                                                                 |
| 1609 | [Even Odd Tree](https://leetcode.com/problems/even-odd-tree/description/)                                                                                                   | [`isevenoddtree.go`](solutions/golang/isevenoddtree.go)                                                                 |
| 1624 | [Largest Substring Between Two Equal Characters](https://leetcode.com/problems/largest-substring-between-two-equal-characters/description/)                                 | [`maxlengthbetweenequalcharacters.go`](solutions/golang/maxlengthbetweenequalcharacters.go)                             |
| 1630 | [Arithmetic Subarrays](https://leetcode.com/problems/arithmetic-subarrays/description/)                                                                                     | [`checkarithmeticsubarrays.go`](solutions/golang/checkarithmeticsubarrays.go)                                           |
| 1633 | [Percentage of Users Attended a Contest](https://leetcode.com/problems/percentage-of-users-attended-a-contest/description/)                                                 | [`1633.sql`](solutions/sql/1633.sql)                                                                                    |
| 1637 | [Widest Vertical Area Between Two Points Containing No Points](https://leetcode.com/problems/widest-vertical-area-between-two-points-containing-no-points/description/)     | [`maxwidthofverticalarea.go`](solutions/golang/maxwidthofverticalarea.go)                                               |
| 1642 | [Furthest Building You Can Reach](https://leetcode.com/problems/furthest-building-you-can-reach/description/)                                                               | [`furthestbuilding.go`](solutions/golang/furthestbuilding.go)                                                           |
| 1652 | [Defuse the Bomb](https://leetcode.com/problems/defuse-the-bomb/description/)                                                                                               | [`decrypt.go`](solutions/golang/decrypt.go)                                                                             |
| 1657 | [Determine if Two Strings Are Close](https://leetcode.com/problems/determine-if-two-strings-are-close/description/)                                                         | [`closestrings.go`](solutions/golang/closestrings.go)                                                                   |
| 1662 | [Check If Two String Arrays are Equivalent](https://leetcode.com/problems/check-if-two-string-arrays-are-equivalent/description/)                                           | [`arraystringsareequal.go`](solutions/golang/arraystringsareequal.go)                                                   |
| 1685 | [Sum of Absolute Differences in a Sorted Array](https://leetcode.com/problems/sum-of-absolute-differences-in-a-sorted-array/description/)                                   | [`getsumabsolutedifferences.go`](solutions/golang/getsumabsolutedifferences.go)                                         |
| 1688 | [Count of Matches in Tournament](https://leetcode.com/problems/count-of-matches-in-tournament/description/)                                                                 | [`numberofmatches.go`](solutions/golang/numberofmatches.go)                                                             |
| 1704 | [Determine if String Halves Are Alike](https://leetcode.com/problems/determine-if-string-halves-are-alike/description/)                                                     | [`halvesarealike.go`](solutions/golang/halvesarealike.go)                                                               |
| 1716 | [Calculate Money in Leetcode Bank](https://leetcode.com/problems/calculate-money-in-leetcode-bank/description/)                                                             | [`totalmoney.go`](solutions/golang/totalmoney.go)                                                                       |
| 1720 | [Decode XORed Array](https://leetcode.com/problems/decode-xored-array/description/)                                                                                         | [`decode.go`](solutions/golang/decode.go)                                                                               |
| 1732 | [Find the Highest Altitude](https://leetcode.com/problems/find-the-highest-altitude/description/)                                                                           | [`largestaltitude.go`](solutions/golang/largestaltitude.go)                                                             |
| 1743 | [Restore the Array From Adjacent Pairs](https://leetcode.com/problems/restore-the-array-from-adjacent-pairs/description/)                                                   | [`restorearray.go`](solutions/golang/restorearray.go)                                                                   |
| 1750 | [Minimum Length of String After Deleting Similar Ends](https://leetcode.com/problems/minimum-length-of-string-after-deleting-similar-ends/description/)                     | [`minimumlength.go`](solutions/golang/minimumlength.go)                                                                 |
| 1757 | [Recyclable and Low Fat Products](https://leetcode.com/problems/recyclable-and-low-fat-products/description/)                                                               | [`1757.sql`](solutions/sql/1757.sql)                                                                                    |
| 1758 | [Minimum Changes To Make Alternating Binary String](https://leetcode.com/problems/minimum-changes-to-make-alternating-binary-string/description/)                           | [`minoperations_1.go`](solutions/golang/minoperations_1.go)                                                             |
| 1759 | [Count Number of Homogenous Substrings](https://leetcode.com/problems/count-number-of-homogenous-substrings/description/)                                                   | [`counthomogenous.go`](solutions/golang/counthomogenous.go)                                                             |
| 1814 | [Count Nice Pairs in an Array](https://leetcode.com/problems/count-nice-pairs-in-an-array/description/)                                                                     | [`countnicepairs.go`](solutions/golang/countnicepairs.go)                                                               |
| 1822 | [Sign of the Product of an Array](https://leetcode.com/problems/sign-of-the-product-of-an-array/description/)                                                               | [`arraysign.go`](solutions/golang/arraysign.go)                                                                         |
| 1838 | [Frequency of the Most Frequent Element](https://leetcode.com/problems/frequency-of-the-most-frequent-element/description/)                                                 | [`maxfrequency.go`](solutions/golang/maxfrequency.go)                                                                   |
| 1845 | [Seat Reservation Manager](https://leetcode.com/problems/seat-reservation-manager/description/)                                                                             | [`seatmanager.go`](solutions/golang/seatmanager.go)                                                                     |
| 1846 | [Maximum Element After Decreasing and Rearranging](https://leetcode.com/problems/maximum-element-after-decreasing-and-rearranging/description/)                             | [`maximumelementafterdecrementingandrearranging.go`](solutions/golang/maximumelementafterdecrementingandrearranging.go) |
| 1877 | [Minimize Maximum Pair Sum in Array](https://leetcode.com/problems/minimize-maximum-pair-sum-in-array/description/)                                                         | [`minpairsum.go`](solutions/golang/minpairsum.go)                                                                       |
| 1887 | [Reduction Operations to Make the Array Elements Equal](https://leetcode.com/problems/reduction-operations-to-make-the-array-elements-equal/description/)                   | [`reductionoperations.go`](solutions/golang/reductionoperations.go)                                                     |
| 1897 | [Redistribute Characters to Make All Strings Equal](https://leetcode.com/problems/redistribute-characters-to-make-all-strings-equal/description/)                           | [`makeequal.go`](solutions/golang/makeequal.go)                                                                         |
| 1903 | [Largest Odd Number in String](https://leetcode.com/problems/largest-odd-number-in-string/description/)                                                                     | [`largestoddnumber.go`](solutions/golang/largestoddnumber.go)                                                           |
| 1913 | [Maximum Product Difference Between Two Pairs](https://leetcode.com/problems/maximum-product-difference-between-two-pairs/description/)                                     | [`maxproductdifference.go`](solutions/golang/maxproductdifference.go)                                                   |
| 1921 | [Eliminate Maximum Number of Monsters](https://leetcode.com/problems/eliminate-maximum-number-of-monsters/description/)                                                     | [`eliminatemaximum.go`](solutions/golang/eliminatemaximum.go)                                                           |
| 1930 | [Unique Length-3 Palindromic Subsequences](https://leetcode.com/problems/unique-length-3-palindromic-subsequences/description/)                                             | [`countpalindromicsubsequence.go`](solutions/golang/countpalindromicsubsequence.go)                                     |
| 1961 | [Check If String Is a Prefix of Array](https://leetcode.com/problems/check-if-string-is-a-prefix-of-array/description/)                                                     | [`isprefixstring.go`](solutions/golang/isprefixstring.go)                                                               |
| 1973 | [Count Nodes Equal to Sum of Descendants](https://leetcode.com/problems/count-nodes-equal-to-sum-of-descendants/description/)                                               | [`equaltodescendants.go`](solutions/golang/equaltodescendants.go)                                                       |
| 1980 | [Find Unique Binary String](https://leetcode.com/problems/find-unique-binary-string/description/)                                                                           | [`finddifferentbinarystring.go`](solutions/golang/finddifferentbinarystring.go)                                         |
| 2000 | [Reverse Prefix of Word](https://leetcode.com/problems/reverse-prefix-of-word/description/)                                                                                 | [`reverseprefix.go`](solutions/golang/reverseprefix.go)                                                                 |
| 2090 | [K Radius Subarray Averages](https://leetcode.com/problems/k-radius-subarray-averages/description/)                                                                         | [`getaverages.go`](solutions/golang/getaverages.go)                                                                     |
| 2108 | [Find First Palindromic String in the Array](https://leetcode.com/problems/find-first-palindromic-string-in-the-array/description/)                                         | [`firstpalindrome.go`](solutions/golang/firstpalindrome.go)                                                             |
| 2125 | [Number of Laser Beams in a Bank](https://leetcode.com/problems/number-of-laser-beams-in-a-bank/description/)                                                               | [`numberofbeams.go`](solutions/golang/numberofbeams.go)                                                                 |
| 2149 | [Rearrange Array Elements by Sign](https://leetcode.com/problems/rearrange-array-elements-by-sign/description/)                                                             | [`rearrangearray.go`](solutions/golang/rearrangearray.go)                                                               |
| 2225 | [Find Players With Zero or One Losses](https://leetcode.com/problems/find-players-with-zero-or-one-losses/description/)                                                     | [`findwinners.go`](solutions/golang/findwinners.go)                                                                     |
| 2244 | [Minimum Rounds to Complete All Tasks](https://leetcode.com/problems/minimum-rounds-to-complete-all-tasks/description/)                                                     | [`minimumrounds.go`](solutions/golang/minimumrounds.go)                                                                 |
| 2259 | [Remove Digit From Number to Maximize Result](https://leetcode.com/problems/remove-digit-from-number-to-maximize-result/description/)                                       | [`removedigit.go`](solutions/golang/removedigit.go)                                                                     |
| 2264 | [Largest 3-Same-Digit Number in String](https://leetcode.com/problems/largest-3-same-digit-number-in-string/description/)                                                   | [`largestgoodinteger.go`](solutions/golang/largestgoodinteger.go)                                                       |
| 2265 | [Count Nodes Equal to Average of Subtree](https://leetcode.com/problems/count-nodes-equal-to-average-of-subtree/description/)                                               | [`averageofsubtree.go`](solutions/golang/averageofsubtree.go)                                                           |
| 2353 | [Design a Food Rating System](https://leetcode.com/problems/design-a-food-rating-system/description/)                                                                       | [`foodratings.go`](solutions/golang/foodratings.go)                                                                     |
| 2357 | [Make Array Zero by Subtracting Equal Amounts](https://leetcode.com/problems/make-array-zero-by-subtracting-equal-amounts/description/)                                     | [`minimumoperations.go`](solutions/golang/minimumoperations.go)                                                         |
| 2385 | [Amount of Time for Binary Tree to Be Infected](https://leetcode.com/problems/amount-of-time-for-binary-tree-to-be-infected/description/)                                   | [`amountoftime.go`](solutions/golang/amountoftime.go)                                                                   |
| 2391 | [Minimum Amount of Time to Collect Garbage](https://leetcode.com/problems/minimum-amount-of-time-to-collect-garbage/description/)                                           | [`garbagecollection.go`](solutions/golang/garbagecollection.go)                                                         |
| 2433 | [Find The Original Array of Prefix Xor](https://leetcode.com/problems/find-the-original-array-of-prefix-xor/description/)                                                   | [`findarray.go`](solutions/golang/findarray.go)                                                                         |
| 2482 | [Difference Between Ones and Zeros in Row and Column](https://leetcode.com/problems/difference-between-ones-and-zeros-in-row-and-column/description/)                       | [`onesminuszeros.go`](solutions/golang/onesminuszeros.go)                                                               |
| 2485 | [Find the Pivot Integer](https://leetcode.com/problems/find-the-pivot-integer/description/)                                                                                 | [`pivotinteger.go`](solutions/golang/pivotinteger.go)                                                                   |
| 2490 | [Circular Sentence](https://leetcode.com/problems/circular-sentence/description/)                                                                                           | [`iscircularsentence.go`](solutions/golang/iscircularsentence.go)                                                       |
| 2540 | [Minimum Common Value](https://leetcode.com/problems/minimum-common-value/description/)                                                                                     | [`getcommon.go`](solutions/golang/getcommon.go)                                                                         |
| 2549 | [Count Distinct Numbers on Board](https://leetcode.com/problems/count-distinct-numbers-on-board/description/)                                                               | [`distinctintegers.go`](solutions/golang/distinctintegers.go)                                                           |
| 2595 | [Number of Even and Odd Bits](https://leetcode.com/problems/number-of-even-and-odd-bits/description/)                                                                       | [`evenoddbit.go`](solutions/golang/evenoddbit.go)                                                                       |
| 2610 | [Convert an Array Into a 2D Array With Conditions](https://leetcode.com/problems/convert-an-array-into-a-2d-array-with-conditions/description/)                             | [`findmatrix.go`](solutions/golang/findmatrix.go)                                                                       |
| 2620 | [Counter](https://leetcode.com/problems/counter/description/)                                                                                                               | [`createCounter.js`](solutions/javascript/createCounter.js)                                                             |
| 2667 | [Create Hello World Function](https://leetcode.com/problems/create-hello-world-function/description/)                                                                       | [`createHelloWorld.js`](solutions/javascript/createHelloWorld.js)                                                       |
| 2704 | [To Be Or Not To Be](https://leetcode.com/problems/to-be-or-not-to-be/description/)                                                                                         | [`expect.js`](solutions/javascript/expect.js)                                                                           |
| 2706 | [Buy Two Chocolates](https://leetcode.com/problems/buy-two-chocolates/description/)                                                                                         | [`buychoco.go`](solutions/golang/buychoco.go)                                                                           |
| 2785 | [Sort Vowels in a String](https://leetcode.com/problems/sort-vowels-in-a-string/description/)                                                                               | [`sortvowels.go`](solutions/golang/sortvowels.go)                                                                       |
| 2849 | [Determine if a Cell Is Reachable at a Given Time](https://leetcode.com/problems/determine-if-a-cell-is-reachable-at-a-given-time/description/)                             | [`isreachableattime.go`](solutions/golang/isreachableattime.go)                                                         |
| 2864 | [Maximum Odd Binary Number](https://leetcode.com/problems/maximum-odd-binary-number/description/)                                                                           | [`maximumoddbinarynumber.go`](solutions/golang/maximumoddbinarynumber.go)                                               |
| 2870 | [Minimum Number of Operations to Make Array Empty](https://leetcode.com/problems/minimum-number-of-operations-to-make-array-empty/description/)                             | [`minoperations_2.go`](solutions/golang/minoperations_2.go)                                                             |
| 2877 | [Create a DataFrame from List](https://leetcode.com/problems/create-a-dataframe-from-list/description/)                                                                     | [`create_dataframe.py`](solutions/python/create_dataframe.py)                                                           |
| 2966 | [Divide Array Into Arrays With Max Difference](https://leetcode.com/problems/divide-array-into-arrays-with-max-difference/description/)                                     | [`dividearray.go`](solutions/golang/dividearray.go)                                                                     |
| 2971 | [Find Polygon With the Largest Perimeter](https://leetcode.com/problems/find-polygon-with-the-largest-perimeter/description/)                                               | [`largestperimeter.go`](solutions/golang/largestperimeter.go)                                                           |
| 3005 | [Count Elements With Maximum Frequency](https://leetcode.com/problems/count-elements-with-maximum-frequency/description/)                                                   | [`maxfrequencyelements.go`](solutions/golang/maxfrequencyelements.go)                                                   |
| 3062 | [Winner of the Linked List Game](https://leetcode.com/problems/winner-of-the-linked-list-game/description/)                                                                 | [`gameresult.go`](solutions/golang/gameresult.go)                                                                       |
| 3063 | [Linked List Frequency](https://leetcode.com/problems/linked-list-frequency/description/)                                                                                   | [`frequenciesofelements.go`](solutions/golang/frequenciesofelements.go)                                                 |

In revision
-----------

1. Longest Palindromic Substring
2. String to Integer (atoi)
3. Container With Most Water
4. Roman to Integer
5. Longest Common Prefix
6. Valid Parentheses
7. Merge Two Sorted Lists
8. Swap Nodes in Pairs
9. Remove Duplicates from Sorted Array
10. Remove Element
11. Find the Index of the First Occurrence in a String
12. Find First and Last Position of Element in Sorted Array
13. Valid Sudoku
14. Rotate Image
15. Spiral Matrix
16. Plus One
17. Add Binary
18. Simplify Path
19. Remove Duplicates from Sorted List II
20. Remove Duplicates from Sorted List
21. Merge Sorted Array
22. Reverse Linked List II
23. Validate Binary Search Tree
24. Symmetric Tree
25. Binary Tree Level Order Traversal
26. Binary Tree Zigzag Level Order Traversal
27. Maximum Depth of Binary Tree
28. Convert Sorted Array to Binary Search Tree
29. Minimum Depth of Binary Tree
30. Path Sum
31. Best Time to Buy and Sell Stock
32. Best Time to Buy and Sell Stock II
33. Valid Palindrome
34. Single Number
35. LRU Cache
36. Reverse Words in a String
37. Min Stack
38. Find Peak Element
39. Two Sum II - Input Array Is Sorted
40. Rotate Array
41. Reverse Bits
42. Binary Tree Right Side View
43. Remove Linked List Elements
44. Count Primes
45. Isomorphic Strings
46. Reverse Linked List
47. Contains Duplicate
48. Implement Stack using Queues
49. Summary Ranges
50. Majority Element II
51. Palindrome Linked List
52. Lowest Common Ancestor of a Binary Tree
53. Delete Node in a Linked List
54. Sliding Window Maximum
55. Closest Binary Search Tree Value
56. First Bad Version
57. Find the Duplicate Number
58. Word Pattern
59. Power of Three
60. Odd Even Linked List
61. Increasing Triplet Subsequence
62. Counting Bits
63. Flatten Nested List Iterator
64. Integer Break
65. Reverse Vowels of a String
66. Intersection of Two Arrays II
67. Guess Number Higher or Lower
68. Ransom Note
69. Shuffle an Array
70. Find the Difference
71. Is Subsequence
72. Decode String
73. Fizz Buzz
74. Path Sum III
75. Find All Anagrams in a String
76. String Compression
77. Delete Node in a BST
78. Hamming Distance
79. Next Greater Element I
80. Diagonal Traverse
81. Find Largest Value in Each Tree Row
82. Minimum Absolute Difference in BST
83. Subarray Sum Equals K
84. Array Partition
85. Permutation in String
86. Can Place Flowers
87. Dota2 Senate
88. Search in a Binary Search Tree
89. Insert into a Binary Search Tree
90. Design HashMap
91. Design Linked List
92. Subarray Product Less Than K
93. Asteroid Collision
94. Find Smallest Letter Greater Than Target
95. Largest Number At Least Twice of Others
96. Jewels and Stones
97. Minimum Distance Between BST Nodes
98. Bus Routes
99. Binary Trees With Factors
100. Peak Index in a Mountain Array
101. Online Stock Span
102. Number of Recent Calls
103. Validate Stack Sequences
104. Longest Arithmetic Subsequence
105. Remove All Adjacent Duplicates In String
106. Greatest Common Divisor of Strings
107. Find in Mountain Array
108. Print in Order
109. Largest Unique Number
110. Snapshot Array
111. Maximum Level Sum of a Binary Tree
112. Maximum Number of Balloons
113. Count Vowels Permutation
114. Check If It Is a Straight Line
115. Count Number of Nice Subarrays
116. Convert Binary Number in a Linked List to Integer
117. Deepest Leaves Sum
118. Number of Steps to Reduce a Number to Zero
119. Count Negative Numbers in a Sorted Matrix
120. Longest ZigZag Path in a Binary Tree
121. Find Lucky Integer in an Array
122. Build Array Where You Can Find The Maximum Exactly K Comparisons
123. Constrained Subsequence Sum
124. Counting Elements
125. Kids With the Greatest Number of Candies
126. Longest Continuous Subarray With Absolute Diff Less Than or Equal to Limit
127. Count Good Nodes in Binary Tree
128. Max Dot Product of Two Subsequences
129. Longest Subarray of 1's After Deleting One Element
130. Can Make Arithmetic Progression From Sequence
131. Number of Good Pairs
132. Make The String Great
133. Richest Customer Wealth
134. Max Number of K-Sum Pairs
135. Maximum Erasure Value
136. Swapping Nodes in a Linked List
137. Sum of Unique Elements
138. Merge Strings Alternately
139. Maximum Score of a Good Subarray
140. Check if the Sentence Is Pangram
141. Check if All Characters Have Equal Number of Occurrences
142. Minimum Number of Operations to Make Array Continuous
143. Remove Colored Pieces if Both Neighbors are the Same Color
144. Parallel Courses III
145. Reverse Nodes in Even Length Groups
146. Delete the Middle Node of a Linked List
147. Maximum Twin Sum of a Linked List
148. Find the Difference of Two Arrays
149. Add Two Integers
150. Intersection of Multiple Arrays
151. Number of Flowers in Full Bloom
152. Minimum Consecutive Cards to Pick Up
153. Number of Ways to Split Array
154. Successful Pairs of Spells and Potions
155. Max Sum of a Pair With Equal Sum of Digits
156. First Letter to Appear Twice
157. Equal Row and Column Pairs
158. Removing Stars From a String
159. Using a Robot to Print the Lexicographically Smallest String
160. Total Cost to Hire K Workers
161. Design Graph With Shortest Path Calculator
162. Painting the Walls
163. Largest Submatrix With Rearrangements
164. Number of Ways to Divide a Long Corridor
165. Minimum One Bit Operations to Make Integers Zero
166. Decode Ways
167. String Compression II
168. Minimum Difficulty of a Job Schedule
169. Maximum Profit in Job Scheduling
170. Arithmetic Slices II - Subsequence
171. K Inverse Pairs Array
172. Number of Submatrices That Sum to Target
173. Partition Array for Maximum Sum
174. Meeting Rooms III
175. Tree Diameter
176. Cheapest Flights Within K Stops
177. Greatest Common Divisor Traversal
178. Find All People With Secret
