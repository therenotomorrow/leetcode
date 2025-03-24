function dfs(obj) {
  if (!obj) {
    return false
  }

  if (!(obj instanceof Object)) {
    return obj
  }

  if (obj instanceof Array) {
    const arr = []

    for (const curr of obj) {
      const inner = dfs(curr)

      if (!inner) {
        continue
      }

      arr.push(inner)
    }

    return arr
  }

  const reduced = {}

  for (const key in obj) {
    const inner = dfs(obj[key])

    if (!inner) {
      continue
    }

    reduced[key] = inner
  }

  return reduced
}

export const compactObject = function (obj) {
  return dfs(obj)
}
