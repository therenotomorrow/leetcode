import { filter } from './filter'

describe('filter', () => {
  test.each([
    [
      'smoke 1',
      {
        arr: [0, 10, 20, 30],
        fn: (n) => {
          return n > 10
        },
      },
      [20, 30],
    ],
    [
      'smoke 2',
      {
        arr: [1, 2, 3],
        fn: (n, i) => {
          return i === 0
        },
      },
      [1],
    ],
    [
      'smoke 3',
      {
        arr: [-2, -1, 0, 1, 2],
        fn: (n) => {
          return n + 1
        },
      },
      [-2, 0, 1, 2],
    ],
  ])('%p', (name, { arr, fn }, want) => {
    expect(filter(arr, fn)).toEqual(want)
  })
})
