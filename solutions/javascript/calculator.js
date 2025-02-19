export class Calculator {
  constructor(value) {
    this.state = value
  }

  add(value) {
    this.state += value
    return this
  }

  subtract(value) {
    this.state -= value
    return this
  }

  multiply(value) {
    this.state *= value
    return this
  }

  divide(value) {
    if (value === 0) {
      throw divisionByZero
    }
    this.state /= value
    return this
  }

  power(value) {
    this.state **= value
    return this
  }

  getResult() {
    return this.state
  }
}

export const divisionByZero = 'Division by zero is not allowed'
