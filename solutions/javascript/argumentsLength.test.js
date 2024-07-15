const argumentsLength = require('./argumentsLength')

describe('argumentsLength', () => {
  test.each([
    ['smoke 1', [5], 1],
    ['smoke 2', [{}, null, '3'], 3]
  ])('%p', (name, args, want) => {
    expect(argumentsLength(...args)).toEqual(want)
  })
})
