import { throttle } from './throttle'

test('smoke 1', async () => {
  const calls = []
  const throttled = throttle((x) => calls.push(x), 100)

  throttled(1)

  expect(calls).toEqual([1])
})

test('smoke 2', async () => {
  const calls = []
  const throttled = throttle((x) => calls.push(x), 100)

  throttled(1)
  throttled(2)

  expect(calls).toEqual([1])

  await new Promise((resolve) => setTimeout(resolve, 200))

  expect(calls).toEqual([1, 2])
})

test('smoke 3', async () => {
  const calls = []
  const throttled = throttle((...args) => calls.push(...args), 70)

  throttled(1)
  expect(calls).toEqual([1])

  await new Promise((resolve) => setTimeout(resolve, 25))

  throttled(2)
  expect(calls).toEqual([1])

  await new Promise((resolve) => setTimeout(resolve, 15))

  throttled(8)
  expect(calls).toEqual([1])

  await new Promise((resolve) => setTimeout(resolve, 50))

  throttled(5, 7)
  expect(calls).toEqual([1, 8])

  await new Promise((resolve) => setTimeout(resolve, 160))

  throttled(9, 4)
  expect(calls).toEqual([1, 8, 5, 7, 9, 4])
})
