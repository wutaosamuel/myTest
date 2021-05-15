import { Obj } from '../../object';

let objects: Obj[] = [];
let changeObj: Obj = new Obj("obj1", 121, ["this", "test"]);

objects.push(new Obj());
objects.push(new Obj("obj0", 0));
objects.push(new Obj("obj1", 1));
objects.push(new Obj("obj2", 2));
console.log("before change:");
console.log(objects);
console.log();
console.log("After change(change obj1, 1 to obj1, 111):")
for (let o of objects) {
  if (o.Name == "obj1")
    o.ID = 111;
}
console.log(objects);
for (let i = 0; i < objects.length; i++) {
  if (objects[i].Name == "obj1")
    objects[i] = changeObj;
}
console.log(objects);
console.log("After change(change obj1, 1 to obj1, 121):")
console.log("'o = changeObj' not work. Using Equal function")
for (let o of objects) {
  if (o.Name == "obj1")
    o.Equal(changeObj);
}
console.log(objects);