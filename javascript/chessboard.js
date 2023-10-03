/**
create this pattern

 # # # #
# # # #
 # # # #
# # # #
 # # # #
# # # #
 # # # #
# # # #

 */

for (let i = 0; i < 8; i++) {
  let str = ""
  
  if (i % 2 === 0) {
    str += " "
  }

  console.log(str + "# # # #")
}
