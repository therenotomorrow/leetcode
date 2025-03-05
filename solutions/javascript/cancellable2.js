export const cancellable2 = function (fn, args, t) {
  fn(...args) // should be called immediately

  const intervalId = setInterval(function () {
    return fn(...args)
  }, t)

  return () => {
    clearInterval(intervalId)
  }
}
