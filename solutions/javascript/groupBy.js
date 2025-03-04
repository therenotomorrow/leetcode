export const groupBy = function () {
  const wrapped = function (fn) {
    return this.reduce((acc, curr) => {
      const key = fn(curr)

      acc[key] ||= []
      acc[key].push(curr)

      return acc
    }, {})
  }

  Array.prototype.groupBy = wrapped

  return wrapped
}
