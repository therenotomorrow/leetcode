import { TimeLimitedCache } from './timeLimitedCache'

test('smoke 1', async () => {
  let got

  const cache = new TimeLimitedCache()

  got = cache.set(1, 42, 100)
  expect(got).toEqual(false)

  got = cache.get(1)
  expect(got).toEqual(42)

  got = cache.count()
  expect(got).toEqual(1)

  await new Promise((resolve) => setTimeout(resolve, 150))

  got = cache.get(1)
  expect(got).toEqual(-1)
})

test('smoke 2', async () => {
  let got

  const cache = new TimeLimitedCache()

  got = cache.set(1, 42, 50)
  expect(got).toEqual(false)

  got = cache.set(1, 50, 100)
  expect(got).toEqual(true)

  got = cache.get(1)
  expect(got).toEqual(50)

  await new Promise((resolve) => setTimeout(resolve, 80))

  got = cache.get(1)
  expect(got).toEqual(50)

  await new Promise((resolve) => setTimeout(resolve, 80))

  got = cache.get(1)
  expect(got).toEqual(-1)

  got = cache.count()
  expect(got).toEqual(0)
})
