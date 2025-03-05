import { cancellable2 } from './cancellable2'

test('smoke 1', async () => {
  let calls = 0

  const cancelTimeMs = 190
  const cancelFn = cancellable2(
    (x) => {
      calls++
      return x * 2
    },
    [4],
    35,
  )
  setTimeout(cancelFn, cancelTimeMs)

  await new Promise((resolve) => setTimeout(resolve, 200))

  expect(calls).toBe(6)
})

test('smoke 2', async () => {
  let calls = 0

  const cancelTimeMs = 165
  const cancelFn = cancellable2(
    (x1, x2) => {
      calls++
      return x1 * x2
    },
    [2, 5],
    30,
  )
  setTimeout(cancelFn, cancelTimeMs)

  await new Promise((resolve) => setTimeout(resolve, 200))

  expect(calls).toBe(6)
})

test('smoke 3', async () => {
  let calls = 0

  const cancelTimeMs = 180
  const cancelFn = cancellable2(
    (x1, x2, x3) => {
      calls++
      return x1 + x2 + x3
    },
    [5, 1, 3],
    50,
  )
  setTimeout(cancelFn, cancelTimeMs)

  await new Promise((resolve) => setTimeout(resolve, 200))

  expect(calls).toBe(4)
})
