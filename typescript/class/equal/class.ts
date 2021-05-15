import { Obj } from '../../object';

let objects: Obj[] = [];
let obj1: Obj = new Obj("obj1", 121, ["this", "test"], [new Obj("obj+", 11), new Obj("obj+", 12)]);

objects.push(new Obj());
objects.push(new Obj("obj0", 0));
objects.push(new Obj("obj1", 121, ["this", "test"]));
objects.push(new Obj("obj2", 2));
console.log("Original objects: ");
console.log(objects);

console.log();
console.log("After equal: ");
console.log("Test Equal function: ");
console.log("[] or class[] both can work");
objects[3].Equal(obj1);
console.log(objects);
console.log(objects[3].OBJ);

console.log();
console.log("After change obj1's OBJ[], not change objects: ");
console.log("Do changed");
obj1.OBJ[0].Name = "ChangedObj1";
console.log(objects);
console.log(objects[3].OBJ);

console.log();
console.log("After Equal to a new object: ")
let newObjects: Obj[] = [];
for (let o of objects)
  newObjects.push(o);
obj1.OBJ[0].Name = "ChangedNewObj1";
console.log(newObjects);
console.log(newObjects[3].OBJ);