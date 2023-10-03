function deepEqual(obj1, obj2) {
  return JSON.stringify(obj1) === JSON.stringify(obj2)
}

let obj = { here: { is: "an" }, object: 2 }

console.log(deepEqual(obj, obj))
// → true
console.log(deepEqual(obj, { here: 1, object: 2 }))
// → false
console.log(deepEqual(obj, { here: { is: "an" }, object: 2 }))
// → true
