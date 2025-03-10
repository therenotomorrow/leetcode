export const timeLimit = function (fn, t) {
  return async function (...args) {
    return new Promise((resolve, reject) => {
      const timeId = setTimeout(() => {
        reject(timeLimitExceeded)
      }, t)

      return fn(...args)
        .then(resolve)
        .catch(reject)
        .finally(() => clearTimeout(timeId))
    })
  }
}

export const timeLimitExceeded = 'Time Limit Exceeded'
