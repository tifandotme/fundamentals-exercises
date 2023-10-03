for (var index = 0; index < 5; index++) {
  setTimeout(() => {
    console.log(index);
  }, 1000);
}

// why different output
for (let index = 0; index < 5; index++) {
  setTimeout(() => {
    console.log(index);
  }, 1000);
}

// When using var, index++ will also increment the callback function inside setTimeout
