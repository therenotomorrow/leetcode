const createCounter2 = function (n) {
  let curr = n

  return {
    increment: () => {
      curr++
      return curr
    },
    decrement: () => {
      curr--
      return curr
    },
    reset: () => {
      curr = n
      return curr
    }
  }
}

module.exports = createCounter2
