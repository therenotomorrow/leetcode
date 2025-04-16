export const areDeeplyEqual = function (o1, o2) {
  if (o1 === o2) {
    return true
  }

  if (o1 === null || o2 === null || String(o1) !== String(o2)) {
    return false
  }

  if (o1 instanceof Array && o2 instanceof Array) {
    if (o1.length !== o2.length) {
      return false
    }

    for (let idx in o1) {
      if (!areDeeplyEqual(o1[idx], o2[idx])) {
        return false
      }
    }

    return true
  }

  if (Object.keys(o1).length !== Object.keys(o2).length) {
    return false
  }

  for (const key in o1) {
    if (!areDeeplyEqual(o1[key], o2[key])) {
      return false
    }
  }

  return true
}
