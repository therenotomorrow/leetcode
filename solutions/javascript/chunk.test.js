import { chunk } from './chunk'

describe('chunk', () => {
  test.each([
    [
      'smoke 1',
      {
        arr: [1, 2, 3, 4, 5],
        size: 1,
      },
      [[1], [2], [3], [4], [5]],
    ],
    [
      'smoke 2',
      {
        arr: [1, 9, 6, 3, 2],
        size: 3,
      },
      [
        [1, 9, 6],
        [3, 2],
      ],
    ],
    [
      'smoke 3',
      {
        arr: [8, 5, 3, 2, 6],
        size: 6,
      },
      [[8, 5, 3, 2, 6]],
    ],
    [
      'smoke 4',
      {
        arr: [],
        size: 1,
      },
      [],
    ],
  ])('%p', (name, { arr, size }, want) => {
    expect(chunk(arr, size)).toEqual(want)
  })
})
