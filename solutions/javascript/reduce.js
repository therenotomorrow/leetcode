export const reduce = function (nums, fn, init) {
  nums.forEach(function (val) {
    init = fn(init, val)
  })

  return init
}
