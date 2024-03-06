const createCounter = function (n) {
  return () => n++
}

module.exports = createCounter
