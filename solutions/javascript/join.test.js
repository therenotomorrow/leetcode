import { join } from './join'

describe('join', () => {
  test.each([
    [
      'smoke 1',
      {
        arr1: [
          {
            id: 1,
            x: 1,
          },
          {
            id: 2,
            x: 9,
          },
        ],
        arr2: [
          {
            id: 3,
            x: 5,
          },
        ],
      },
      [
        {
          id: 1,
          x: 1,
        },
        {
          id: 2,
          x: 9,
        },
        {
          id: 3,
          x: 5,
        },
      ],
    ],
    [
      'smoke 2',
      {
        arr1: [
          {
            id: 1,
            x: 2,
            y: 3,
          },
          {
            id: 2,
            x: 3,
            y: 6,
          },
        ],
        arr2: [
          {
            id: 2,
            x: 10,
            y: 20,
          },
          {
            id: 3,
            x: 0,
            y: 0,
          },
        ],
      },
      [
        {
          id: 1,
          x: 2,
          y: 3,
        },
        {
          id: 2,
          x: 10,
          y: 20,
        },
        {
          id: 3,
          x: 0,
          y: 0,
        },
      ],
    ],
    [
      'smoke 3',
      {
        arr1: [
          {
            id: 1,
            b: { b: 94 },
            v: [4, 3],
            y: 48,
          },
        ],
        arr2: [
          {
            id: 1,
            b: { c: 84 },
            v: [1, 3],
          },
        ],
      },
      [
        {
          id: 1,
          b: { c: 84 },
          v: [1, 3],
          y: 48,
        },
      ],
    ],
  ])('%p', (name, { arr1, arr2 }, want) => {
    expect(join(arr1, arr2)).toEqual(want)
  })
})
