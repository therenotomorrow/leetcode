import { cancellable } from './cancellable'

test('smoke 1', async () => {
  let got = -1

  const cancelTimeMs = 50
  const cancelFn = cancellable(
    (x) => {
      got = x * 5
      return got
    },
    [2],
    20,
  )
  setTimeout(cancelFn, cancelTimeMs)

  await new Promise((resolve) => setTimeout(resolve, 30))

  expect(got).toBe(10)
})

test('smoke 2', async () => {
  let got = -1

  const cancelTimeMs = 50
  const cancelFn = cancellable(
    (x) => {
      got = x ** 2
      return got
    },
    [2],
    100,
  )

  setTimeout(cancelFn, cancelTimeMs)

  await new Promise((resolve) => setTimeout(resolve, 60))

  expect(got).toEqual(-1)
})

test('smoke 3', async () => {
  let got = -1

  const cancelTimeMs = 100
  const cancelFn = cancellable(
    (x1, x2) => {
      got = x1 * x2
      return got
    },
    [2, 4],
    30,
  )
  setTimeout(cancelFn, cancelTimeMs)

  await new Promise((resolve) => setTimeout(resolve, 50))

  expect(got).toEqual(8)
})
