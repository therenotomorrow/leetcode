import { groupBy } from './groupBy'

describe('groupBy', () => {
  test.each([
    [
      'smoke 1',
      {
        array: [{ id: '1' }, { id: '1' }, { id: '2' }],
        fn: (item) => item.id,
      },
      {
        1: [{ id: '1' }, { id: '1' }],
        2: [{ id: '2' }],
      },
    ],
    [
      'smoke 2',
      {
        array: [
          [1, 2, 3],
          [1, 3, 5],
          [1, 5, 9],
        ],
        fn: (list) => String(list[0]),
      },
      {
        1: [
          [1, 2, 3],
          [1, 3, 5],
          [1, 5, 9],
        ],
      },
    ],
    [
      'smoke 3',
      {
        array: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10],
        fn: (n) => String(n > 5),
      },
      {
        true: [6, 7, 8, 9, 10],
        false: [1, 2, 3, 4, 5],
      },
    ],
  ])('%p', (name, { array, fn }, want) => {
    array.groupBy = groupBy()

    expect(array.groupBy(fn)).toEqual(want)
  })
})
