export const cancellable = function (fn, args, t) {
  const timeId = setTimeout(function () {
    return fn(...args)
  }, t)

  return () => {
    clearTimeout(timeId)
  }
}
