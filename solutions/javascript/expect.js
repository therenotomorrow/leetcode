export const expect = function (val) {
  return {
    toBe: function (other) {
      if (val !== other) {
        throw Error('Not Equal')
      }

      return true
    },

    notToBe: function (other) {
      if (val === other) {
        throw Error('Equal')
      }

      return true
    },
  }
}
