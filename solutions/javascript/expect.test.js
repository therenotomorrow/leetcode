const expectFunc = require('./expect')

describe('Expect function tests', () => {
  test('smoke 1', () => {
    expect(expectFunc(5).toBe(5)).toEqual(true)
  })

  test('smoke 2', () => {
    expect(() => expectFunc(5).toBe(null)).toThrowError('Not Equal')
  })

  test('smoke 3', () => {
    expect(expectFunc(5).notToBe(null)).toEqual(true)
  })
})
