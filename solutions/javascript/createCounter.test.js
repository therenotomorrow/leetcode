const createCounter = require('./createCounter')

describe('createCounter', () => {
  test.each([
    ['smoke 1', { n: 10, calls: 3 }, [10, 11, 12]],
    ['smoke 2', { n: -2, calls: 5 }, [-2, -1, 0, 1, 2]]
  ])('%p', (name, args, want) => {
    const counter = createCounter(args.n)

    for (let i = 0; i < args.calls; i++) {
      expect(counter()).toEqual(want[i])
    }
  })
})
