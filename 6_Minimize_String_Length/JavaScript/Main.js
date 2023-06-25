var minimizedStringLength = function (s) {
  let length = 1
  let splitted = s.split("")
  let sorted = splitted.sort()
  let currentChar = sorted[0]

  for (let i = 0; i < splitted.length; i++) {
    if (splitted[i] !== currentChar) {
      currentChar = splitted[i]
      length++
    }
  }

  return length
};

console.log(minimizedStringLength("idi"));