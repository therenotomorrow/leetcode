export const map = function (arr, fn) {
  let ans = []

  for (let i = 0; i < arr.length; i++) {
    let val = fn(arr[i], i)

    ans.push(val)
  }

  return ans
}
