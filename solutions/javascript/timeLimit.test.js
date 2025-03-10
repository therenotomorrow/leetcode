import { timeLimit, timeLimitExceeded } from './timeLimit'

test('smoke 1', async () => {
  const fn = async (n) => {
    await new Promise((res) => setTimeout(res, 100))

    return n * n
  }

  const limited = await timeLimit(fn, 50)

  try {
    await limited(5)
  } catch (err) {
    expect(err).toEqual(timeLimitExceeded)
  }
})

test('smoke 2', async () => {
  const fn = async (n) => {
    await new Promise((res) => setTimeout(res, 100))

    return n * n
  }

  const got = await timeLimit(fn, 150)(5)

  expect(got).toEqual(25)
})

test('smoke 3', async () => {
  const fn = async (a, b) => {
    await new Promise((res) => setTimeout(res, 120))
    return a + b
  }

  const got = await timeLimit(fn, 150)(5, 10)

  expect(got).toEqual(15)
})

test('smoke 4', async () => {
  const fn = async () => {
    throw 'Error'
  }

  const limited = await timeLimit(fn, 1000)

  try {
    await limited()
  } catch (err) {
    expect(err).toEqual('Error')
  }
})
