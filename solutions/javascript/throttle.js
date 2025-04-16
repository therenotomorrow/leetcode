export const throttle = function (fn, t) {
  let intervalId = null
  let inProgress = null

  const intervalFunction = () => {
    if (inProgress === null) {
      clearInterval(intervalId)
      intervalId = null
    } else {
      fn(...inProgress)
      inProgress = null
    }
  }

  return function throttled(...args) {
    if (intervalId) {
      inProgress = args
    } else {
      fn(...args)
      intervalId = setInterval(intervalFunction, t)
    }
  }
}
