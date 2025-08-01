function mergeTwoLists(list1, list2) {

  if (list1 === null) {
    return list2
  } else if (list2 === null) {
    return list1
  }

  if (list1.data < list2.data) {
    list1.next = mergeTwoLists(list1.next, list2);
    return list1;
  }
  else {
    list2.next = mergeTwoLists(list1, list2.next);
    return list2;
  }
}

console.log(mergeTwoLists([1, 2, 4], [1, 3, 4]));