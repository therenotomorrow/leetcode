import { once } from './once'

describe('once', () => {
  test.each([
    [
      'smoke 1',
      {
        fn: (a, b, c) => {
          return a + b + c
        },
        calls: [
          [1, 2, 3],
          [2, 3, 6],
        ],
      },
      [6, undefined],
    ],
    [
      'smoke 2',
      {
        fn: (a, b, c) => {
          return a * b * c
        },
        calls: [
          [5, 7, 4],
          [2, 3, 6],
          [4, 6, 8],
        ],
      },
      [140, undefined, undefined],
    ],
  ])('%p', (name, { fn, calls }, want) => {
    fn = once(fn)

    for (const idx in calls) {
      const got = fn(...calls[idx])

      expect(got).toEqual(want[idx])
    }
  })
})
