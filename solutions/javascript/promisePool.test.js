import { promisePool } from './promisePool'

const funcs = function () {
  let logs = []

  const functions = [
    () => new Promise((res) => setTimeout(res, 300)).then(() => logs.push(300)),
    () => new Promise((res) => setTimeout(res, 400)).then(() => logs.push(400)),
    () => new Promise((res) => setTimeout(res, 200)).then(() => logs.push(200)),
  ]

  return [logs, functions]
}

test('smoke 1', async () => {
  const [logs, functions] = funcs()

  const start = Date.now()

  await promisePool([...functions], 2)

  const timer = Date.now() - start

  expect(timer).toBeGreaterThanOrEqual(500)
  expect(timer).toBeLessThan(550)

  expect(logs).toEqual([300, 400, 200])
})

test('smoke 2', async () => {
  const [logs, functions] = funcs()

  const start = Date.now()

  await promisePool([...functions], 5)

  const timer = Date.now() - start

  expect(timer).toBeGreaterThanOrEqual(400)
  expect(timer).toBeLessThan(450)

  expect(logs).toEqual([200, 300, 400])
})

test('smoke 3', async () => {
  const [logs, functions] = funcs()

  const start = Date.now()

  await promisePool([...functions], 1)

  const timer = Date.now() - start

  expect(timer).toBeGreaterThanOrEqual(900)
  expect(timer).toBeLessThan(950)

  expect(logs).toEqual([300, 400, 200])
})
