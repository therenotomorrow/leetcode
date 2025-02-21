import { sleep } from './sleep'

describe('sleep', () => {
  test.each([
    ['smoke 1', { millis: 100 }, 100],
    ['smoke 2', { millis: 100 }, 100],
  ])('%p', async (name, { millis }, want) => {
    const start = Date.now()

    await sleep(millis)

    expect(Date.now() - start).toBeGreaterThan(want)
  })
})
