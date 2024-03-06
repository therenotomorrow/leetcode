const createCounter = require('./createCounter')

describe('createCounter', () => {
  test.each([
    ['smoke 1', { n: 10, calls: 3 }, [10, 11, 12]],
    ['smoke 2', { n: -2, calls: 5 }, [-2, -1, 0, 1, 2]]
  ])('%p', (name, { n, calls }, want) => {
    const counter = createCounter(n)

    for (let i = 0; i < calls; i++) {
      expect(counter()).toEqual(want[i])
    }
  })
})
