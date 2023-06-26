var sortPeople = function (names, heights) {
  let newNames = []
  
  for (let i = 0; i < heights.length; i++) {
    let max = Math.max(...heights)
    let maxIndex = heights.indexOf(max)
    heights.splice(maxIndex, 1, -1)
    newNames.push(names[maxIndex])
  }
  
  return newNames
};

console.log(sortPeople(["Alice","Bob","Bob"], [155, 185, 150]));