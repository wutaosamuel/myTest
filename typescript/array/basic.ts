let array1 = ["string", "should", "be", "equal"];
let array2 = ["string", "should", "be", "equal"];

console.log("not work")
console.log(array1 == array2);

// only can work for each element
console.log();
console.log("One by one: ")
for (let i = 0; i < array1.length; i++) {
  if (array1[i] != array2[i]) {
    console.log("not equal, false");
    break;
  }
}
console.log("work");