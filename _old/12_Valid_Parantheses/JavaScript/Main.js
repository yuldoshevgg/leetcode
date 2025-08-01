let isValid = function (s) {
  let result = []

  for (let i = 0; i < s.length; i++) {
    if (s[i] === "(") {
      result.push(s[i])
    } else if (s[i] === "{") {
      result.push(s[i])
    } else if (s[i] === "[") {
      result.push(s[i])
    } else if (s[i] === ")") {
      if (result[result.length - 1] !== "(") {
        return false
      } else {
        result.pop()
      }
    } else if (s[i] === "}") {
      if (result[result.length - 1] !== "{") {
        return false
      } else {
        result.pop()
      }
    } else if (s[i] === "]") {
      if (result[result.length - 1] !== "[") {
        return false
      } else {
        result.pop()
      }
    }
  }

  return result.length === 0
}

console.log(isValid("([)]"));
