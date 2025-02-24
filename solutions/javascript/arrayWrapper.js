export const ArrayWrapper = function (nums) {
  this.data = nums
}

ArrayWrapper.prototype.valueOf = function () {
  return this.data.reduce((acc, curr) => acc + curr, 0)
}

ArrayWrapper.prototype.toString = function () {
  return '[' + this.data.toString() + ']'
}
