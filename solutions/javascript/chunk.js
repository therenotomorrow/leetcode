export const chunk = function (arr, size) {
  const ans = []

  for (let i = 0; i < arr.length; i += size) {
    ans.push(arr.slice(i, i + size))
  }

  return ans
}
