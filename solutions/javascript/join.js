export const join = function (arr1, arr2) {
  arr1.sort((a, b) => a.id - b.id)
  arr2.sort((a, b) => a.id - b.id)

  let i = 0
  let j = 0

  const ans = []

  while (i < arr1.length && j < arr2.length) {
    if (arr1[i].id < arr2[j].id) {
      ans.push({ ...arr1[i] })
      i++
      continue
    }
    if (arr1[i].id > arr2[j].id) {
      ans.push({ ...arr2[j] })
      j++
      continue
    }

    ans.push({ ...arr1[i], ...arr2[j] })
    i++
    j++
  }

  while (i < arr1.length) {
    ans.push({ ...arr1[i] })
    i++
  }

  while (j < arr2.length) {
    ans.push({ ...arr2[j] })
    j++
  }

  return ans
}
