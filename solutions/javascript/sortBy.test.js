import { sortBy } from './sortBy'

test('smoke 1', () => {
  const arr = [5, 4, 1, 2, 3]
  const fn = (x) => x

  const got = sortBy(arr, fn)

  expect(got).toEqual([1, 2, 3, 4, 5])
})

test('smoke 2', () => {
  const arr = [{ x: 1 }, { x: 0 }, { x: -1 }]
  const fn = (d) => d.x

  const got = sortBy(arr, fn)

  expect(got).toEqual([{ x: -1 }, { x: 0 }, { x: 1 }])
})

test('smoke 3', () => {
  const arr = [
    [3, 4],
    [5, 2],
    [10, 1],
  ]
  const fn = (x) => x[1]

  const got = sortBy(arr, fn)

  expect(got).toEqual([
    [10, 1],
    [5, 2],
    [3, 4],
  ])
})
