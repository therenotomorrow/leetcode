import { compactObject } from './compactObject'

describe('compactObject', () => {
  test.each([
    ['smoke 1', [null, 0, false, 1], [1]],
    [
      'smoke 2',
      {
        a: null,
        b: [false, 1],
      },
      { b: [1] },
    ],
    ['smoke 3', [null, 0, 5, [0], [false, 16]], [5, [], [16]]],
  ])('%p', (name, obj, want) => {
    expect(compactObject(obj)).toEqual(want)
  })
})
