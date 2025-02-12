export const filter = function (arr, fn) {
  let ans = []

  for (let i = 0; i < arr.length; i++) {
    let val = arr[i]

    if (fn(val, i)) {
      ans.push(val)
    }
  }

  return ans
}
