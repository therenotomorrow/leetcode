import { map } from './map'

describe('map', () => {
  test.each([
    [
      'smoke 1',
      {
        arr: [1, 2, 3],
        fn: (n) => {
          return n + 1
        },
      },
      [2, 3, 4],
    ],
    [
      'smoke 2',
      {
        arr: [1, 2, 3],
        fn: (n, i) => {
          return n + i
        },
      },
      [1, 3, 5],
    ],
    [
      'smoke 3',
      {
        arr: [10, 20, 30],
        fn: () => {
          return 42
        },
      },
      [42, 42, 42],
    ],
  ])('%p', (name, { arr, fn }, want) => {
    expect(map(arr, fn)).toEqual(want)
  })
})
