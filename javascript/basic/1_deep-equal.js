function deepEqualAlt(obj1, obj2) {
  return JSON.stringify(obj1) === JSON.stringify(obj2)
}

function deepEqual(obj1, obj2) {
  const keys1 = Object.keys(obj1)
  const keys2 = Object.keys(obj2)

  if (keys1.length !== keys2.length) return false

  for (let key of keys1) {
    const value1 = obj1[key]
    const value2 = obj2[key]

    if (isObject(value1) && isObject(value2)) {
      return deepEqual(value1, value2)
    } else {
      return value1 === value2
    }
  }

  return true
}

function isObject(obj) {
  return typeof obj === "object"
}

let obj = { here: { is: "an" }, object: 2 }

console.log(deepEqual(obj, obj))
// → true
console.log(deepEqual(obj, { here: 1, object: 2 }))
// → false
console.log(deepEqual(obj, { here: { is: "an" }, object: 2 }))
// → true

console.log(
  deepEqual(
    {
      hello: {
        world: {
          n: 1,
        },
      },
    },
    {
      hello: {
        world: {
          n: 1,
        },
      },
    }
  )
)
// → true

console.log(
  JSON.stringify(function foo() {
    var foo = 1
  })
)
