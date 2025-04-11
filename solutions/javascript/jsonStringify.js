export const jsonStringify = function (object) {
  let str = ''

  switch (typeof object) {
    case 'object':
      if (object instanceof Array) {
        str = '[' + object.map((elem) => jsonStringify(elem)).join(',') + ']'
      } else if (object) {
        str = '{' + Object.keys(object).map((key) => `"${key}":` + jsonStringify(object[key])) + '}'
      } else {
        str = 'null'
      }
      break
    case 'boolean':
    case 'number':
      str = `${object}`
      break
    case 'string':
      str = `"${object}"`
  }

  return str
}
