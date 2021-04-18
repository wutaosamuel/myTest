import { Obj } from '../object'

let obj0 = new Obj("obj0", 0);
let obj1 = new Obj("obj1", 1);

console.log("Before setting: ");
console.log(obj0);
console.log(obj1);

console.log("After setting: ");
obj1 = obj0;
console.log(obj1);

console.log("let obj = obj0: ");
let obj = obj0;
console.log(obj);