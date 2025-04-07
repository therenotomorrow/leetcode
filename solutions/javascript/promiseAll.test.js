import { promiseAll } from './promiseAll'

test('smoke 1', async () => {
  const functions = [() => new Promise((resolve) => setTimeout(() => resolve(5), 200))]

  const got = await promiseAll(functions)

  expect(got).toEqual([5])
})

test('smoke 2', async () => {
  const error = 'Error'

  const functions = [
    () => new Promise((resolve) => setTimeout(() => resolve(1), 200)),
    () => new Promise((resolve, reject) => setTimeout(() => reject(error), 100)),
  ]

  try {
    await promiseAll(functions)
  } catch (err) {
    expect(err).toEqual(error)
  }
})

test('smoke 3', async () => {
  const functions = [
    () => new Promise((resolve) => setTimeout(() => resolve(4), 50)),
    () => new Promise((resolve) => setTimeout(() => resolve(10), 150)),
    () => new Promise((resolve) => setTimeout(() => resolve(16), 100)),
  ]

  const got = await promiseAll(functions)

  expect(got).toEqual([4, 10, 16])
})
