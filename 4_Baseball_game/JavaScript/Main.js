var baseballGame = function (ops) {
  let record = []
  let total = 0

  for (let i = 0; i < ops.length; i++) {
    if (!isNaN(Number(ops[i]))) {
      record.push(Number(ops[i]))
    } else if (ops[i] === "C") {
      record.pop()
    } else if (ops[i] === "D") {
      record.push(record[record.length - 1] * 2)
    } else if (ops[i] === "+") {
      record.push(record[record.length - 1] + record[record.length - 2])
    }
  }

  if (record.length) {
    for (let i = 0; i < record.length; i++) {
      total += record[i]
    }
  }

  return total
};

console.log(baseballGame(["1","C"]));