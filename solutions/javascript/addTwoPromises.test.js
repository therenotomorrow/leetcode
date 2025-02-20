import { addTwoPromises } from './addTwoPromises'

test('smoke 1', async () => {
  const promise1 = new Promise((resolve) => setTimeout(() => resolve(2), 20))
  const promise2 = new Promise((resolve) => setTimeout(() => resolve(5), 60))

  const got = await addTwoPromises(promise1, promise2)

  expect(got).toEqual(7)
})

test('smoke 2', async () => {
  const promise1 = new Promise((resolve) => setTimeout(() => resolve(10), 50))
  const promise2 = new Promise((resolve) => setTimeout(() => resolve(-12), 30))

  const got = await addTwoPromises(promise1, promise2)

  expect(got).toEqual(-2)
})
