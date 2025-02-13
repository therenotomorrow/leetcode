export const compose = function (functions) {
  return function (x) {
    for (const func of functions.reverse()) {
      x = func(x)
    }

    return x
  }
}
