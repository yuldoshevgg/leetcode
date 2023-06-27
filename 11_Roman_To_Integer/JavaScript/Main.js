let romanToInt = function (s) {
  let result = 0
  let romans = [
    { roman: "I", number: 1 },
    { roman: "V", number: 5 },
    { roman: "X", number: 10 },
    { roman: "L", number: 50 },
    { roman: "C", number: 100 },
    { roman: "D", number: 500 },
    { roman: "M", number: 1000 },
  ]

  if (s.includes("IV")) {
    result += 4
    s = s.replace("IV", "")
  }

  if (s.includes("IX")) {
    result += 9
    s = s.replace("IX", "")
  }

  if (s.includes("XL")) {
    result += 40
    s = s.replace("XL", "")
  }

  if (s.includes("XC")) {
    result += 90
    s = s.replace("XC", "")
  }

  if (s.includes("CD")) {
    result += 400
    s = s.replace("CD", "")
  }

  if (s.includes("CM")) {
    result += 900
    s = s.replace("CM", "")
  }

  romans.forEach(item => {
    if (s.includes(item.roman)) {
      result += ((s.split(item.roman).length - 1) * item.number)
    }
  })

  return result
}

console.log(romanToInt("MCMXCIV"));
