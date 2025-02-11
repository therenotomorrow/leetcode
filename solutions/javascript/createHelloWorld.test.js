import { createHelloWorld } from './createHelloWorld'

describe('createHelloWorld', () => {
  test.each([
    ['smoke 1', [], 'Hello World'],
    ['smoke 2', [{}, null, 42], 'Hello World'],
  ])('%p', (name, args, want) => {
    expect(createHelloWorld()(...args)).toEqual(want)
  })
})
