let findWords = function (words) {
  let existWords = []
  let firstRow = "qwertyuiop"
  let secondRow = "asdfghjkl"
  let thirdRow = "zxcvbnm"

  for (let i = 0; i < words.length; i++) {
    let ac = 0
    let bc = 0
    let cc = 0

    for (let j = 0; j < words[i].length; j ++) {
      let letter = words[i][j].toLowerCase()

      if (firstRow.includes(letter)) {
        ac++
      } else if (secondRow.includes(letter)) {
        bc++
      } else if (thirdRow.includes(letter)) {
        cc++
      }
    }    

    if (ac === words[i].length || bc === words[i].length || cc === words[i].length) {
      existWords.push(words[i])
    }
  }

  return existWords
}

console.log(findWords(["adsdf", "sfd"]));