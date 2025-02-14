export const memoize = function (fn) {
  const cache = new Map()

  return function (...args) {
    const key = `${fn.name}:${JSON.stringify(args)}`

    if (!cache.has(key)) {
      cache.set(key, fn(...args))
    }

    return cache.get(key)
  }
}
