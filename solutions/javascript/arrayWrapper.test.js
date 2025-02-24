import { ArrayWrapper } from './arrayWrapper'

test('smoke 1', async () => {
  const obj1 = new ArrayWrapper([1, 2])
  const obj2 = new ArrayWrapper([3, 4])

  const got = obj1 + obj2

  expect(got).toEqual(10)
})

test('smoke 2', async () => {
  const obj = new ArrayWrapper([23, 98, 42, 70])

  const got = String(obj)

  expect(got).toEqual('[23,98,42,70]')
})

test('smoke 3', async () => {
  const obj1 = new ArrayWrapper([])
  const obj2 = new ArrayWrapper([])

  const got = obj1 + obj2

  expect(got).toEqual(0)
})
