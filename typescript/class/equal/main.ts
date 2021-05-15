import { Obj } from '../../object';

let objects: Obj[] = [];
let obj1: Obj = new Obj("obj1", 121, ["this", "test"]);

objects.push(new Obj());
objects.push(new Obj("obj0", 0));
objects.push(new Obj("obj1", 121, ["this", "test"]));
objects.push(new Obj("obj2", 2));
console.log("Original objects: ")
console.log(objects);

console.log("After equal: ")
console.log("2 value same class cannot equal")
for (let i=0; i < objects.length; i++) {
  if (objects[i] == obj1)
    console.log("class compare value only")
}
console.log(objects);

console.log("After equal: ")
console.log("equal work by function")
for (let i=0; i < objects.length; i++) {
  if (objects[i].IsEqual(obj1))
    console.log("class compare value only")
}
console.log(objects);

console.log("Test Equal function: ")
objects[3].Equal(obj1)
console.log(objects)