import { reduce } from './reduce'

describe('map', () => {
  test.each([
    [
      'smoke 1',
      {
        nums: [1, 2, 3, 4],
        fn: (accum, curr) => {
          return accum + curr
        },
        init: 0,
      },
      10,
    ],
    [
      'smoke 2',
      {
        nums: [1, 2, 3, 4],
        fn: (accum, curr) => {
          return accum + curr * curr
        },
        init: 100,
      },
      130,
    ],
    [
      'smoke 3',
      {
        nums: [],
        fn: () => {
          return 0
        },
        init: 25,
      },
      25,
    ],
  ])('%p', (name, { nums, fn, init }, want) => {
    expect(reduce(nums, fn, init)).toEqual(want)
  })
})
