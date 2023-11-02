function exampleVar() {
  if (true) {
    var x = 10

    console.log("Inside if block (using var):", x)
  }

  console.log("Outside if block (ising var): ", x)
}

function exampleVar2() {
  if (true) {
    let x = 10

    console.log("Inside if block (using let):", x)
  }

  console.log("Outside if block (ising let): ", x)
}

