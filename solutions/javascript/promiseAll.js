export const promiseAll = function (functions) {
  return new Promise((resolve, reject) => {
    if (!functions.length) {
      resolve([])
      return
    }

    let inProgress = functions.length

    const results = new Array(functions.length)

    functions.forEach(async (func, index) => {
      try {
        results[index] = await func()

        inProgress--

        if (!inProgress) {
          resolve(results)
        }
      } catch (err) {
        reject(err)
      }
    })
  })
}
