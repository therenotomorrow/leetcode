export const flat = function (arr, n) {
  const ans = []

  const dive = (nums, level) => {
    for (const num of nums) {
      if (level > 0 && num instanceof Array) {
        dive(num, level - 1)
      } else {
        ans.push(num)
      }
    }
  }

  dive(arr, n)

  return ans
}
