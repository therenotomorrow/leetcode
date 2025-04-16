import { areDeeplyEqual } from './areDeeplyEqual'

describe('areDeeplyEqual', () => {
  test.each([
    [
      'smoke 1',
      {
        o1: {
          x: 1,
          y: 2,
        },
        o2: {
          x: 1,
          y: 2,
        },
      },
      true,
    ],
    [
      'smoke 2',
      {
        o1: {
          y: 2,
          x: 1,
        },
        o2: {
          x: 1,
          y: 2,
        },
      },
      true,
    ],
    [
      'smoke 3',
      {
        o1: {
          x: null,
          L: [1, 2, 3],
        },
        o2: {
          x: null,
          L: ['1', '2', '3'],
        },
      },
      false,
    ],
    [
      'smoke 4',
      {
        o1: true,
        o2: false,
      },
      false,
    ],
  ])('%p', (name, { o1, o2 }, want) => {
    expect(areDeeplyEqual(o1, o2)).toEqual(want)
  })
})
