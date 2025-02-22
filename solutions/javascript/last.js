export const last = function () {
  const fn = function () {
    return this.length === 0 ? -1 : this[this.length - 1]
  }

  Array.prototype.last = fn

  return fn
}
