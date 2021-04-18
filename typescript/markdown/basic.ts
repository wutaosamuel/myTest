import * as marked from 'marked';
import * as fs from 'fs';

export const markdownString = `
# Marked in Typescript

Rendered by **marked**

## Marked TS
`;

export function basicToHtml(): string {
  const mkContent = marked(markdownString);

  return mkContent;
}

export function basicToHtmlFile(): void {
  const html = `
<!doctype html>
<html>
<head>
  <meta charset="utf-8"/>
  <title>Marked in the browser</title>
</head>
<body>
  <div id="content"></div>
  ${basicToHtml()}
</body>
</html>
`;

  fs.open("./basic.html", "w+", function (err, fd) {
    if (err)
      return console.error(err);
  });

  fs.writeFile("./basic.html", html, function (err) {
    if (err) {
      return console.error(err);
    }
  })
}