var sortPeople = function (names, heights) {
  let tmp = []
  let newNames = []

  for (let i = 0; i < heights.length; i++) {
    let obj = {
      name: names[i],
      height: heights[i]
    }

    tmp.push(obj)
  }

  tmp.sort((a, b) => b.height - a.height)
  for (let i = 0; i < tmp.length; i++) {
    newNames.push(tmp[i].name)
  }

  return newNames
};

console.log(sortPeople(["Mary", "John", "Emma"], [180, 165, 170]));