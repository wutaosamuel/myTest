import * as path from 'path';

const parent: string = "/home/pi";
const folder: string = "git";
const file: string = "test.ts";

const folderPath = path.join(parent, folder);
console.log("folder path: ", folderPath);
const filePath = path.join(parent, folder, file);
console.log("file path: ", filePath);
