import { compose } from './compose'

describe('compose', () => {
  test.each([
    [
      'smoke 1',
      {
        functions: [
          (x) => {
            return x + 1
          },
          (x) => {
            return x * x
          },
          (x) => {
            return 2 * x
          },
        ],
        x: 4,
      },
      65,
    ],
    [
      'smoke 2',
      {
        functions: [
          (x) => {
            return 10 * x
          },
          (x) => {
            return 10 * x
          },
          (x) => {
            return 10 * x
          },
        ],
        x: 1,
      },
      1000,
    ],
    [
      'smoke 3',
      {
        functions: [],
        x: 42,
      },
      42,
    ],
  ])('%p', (name, { functions, x }, want) => {
    expect(compose(functions)(x)).toEqual(want)
  })
})
