import { Obj } from '../object';

let objects: Obj[] = [];

objects.push(new Obj());
objects.push(new Obj("obj0", 0));
objects.push(new Obj("obj1", 1));
objects.push(new Obj("obj2", 2));
console.log("before change:")
console.log(objects);

console.log("after for loop:")
for(let i=0; i < objects.length; i++) {
  objects[i].ID = i+3;
}
console.log(objects);