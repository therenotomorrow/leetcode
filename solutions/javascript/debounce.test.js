import { debounce } from './debounce'

test('smoke 1', async () => {
  let got = 0

  const log = (...inputs) => {
    got = inputs.reduce((acc, curr) => acc + curr, got)
  }

  const dlog = debounce(log, 50)

  setTimeout(() => dlog(1), 50)
  setTimeout(() => dlog(2), 75)

  await new Promise((resolve) => setTimeout(resolve, 150))

  expect(got).toEqual(2)
})

test('smoke 2', async () => {
  let got = 0

  const log = (...inputs) => {
    got = inputs.reduce((acc, curr) => acc + curr, got)
  }

  const dlog = debounce(log, 20)

  setTimeout(() => dlog(1), 50)
  setTimeout(() => dlog(2), 100)

  await new Promise((resolve) => setTimeout(resolve, 150))

  expect(got).toEqual(3)
})

test('smoke 3', async () => {
  let got = 0

  const log = (...inputs) => {
    got = inputs.reduce((acc, curr) => acc + curr, got)
  }

  const dlog = debounce(log, 150)

  setTimeout(() => dlog(1, 2), 50)
  setTimeout(() => dlog(3, 4), 300)
  setTimeout(() => dlog(5, 6), 300)

  await new Promise((resolve) => setTimeout(resolve, 500))

  expect(got).toEqual(14)
})
