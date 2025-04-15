const execute = async function (functions) {
  if (functions.length === 0) {
    return
  }

  const fn = functions.shift()
  const res = await fn()

  await execute(functions)

  return res
}

export const promisePool = async function (functions, size) {
  const nPromises = Array(size)
    .fill(null)
    .map(() => execute(functions))

  return await Promise.all(nPromises)
}
