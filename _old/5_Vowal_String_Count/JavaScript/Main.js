var vowalString = function (words, left, right) {
  let vowals = "aeiou"
  let count = 0

  for (let i = left; i <= right; i++ ) {
    let spString = words[i].split('')

    if (vowals.includes(spString[0]) && vowals.includes(spString[spString.length - 1])) {
      count++
    }
  }
  
  return count
};

console.log(vowalString(["hey", "aeo", "mu", "ooo", "artro"], 1, 4));