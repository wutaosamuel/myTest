import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-code-mirror',
  templateUrl: './code-mirror.component.html',
  styleUrls: ['./code-mirror.component.css']
})
export class CodeMirrorComponent implements OnInit {
  codeString: string = "123";
  codeMirrorOptions: any = {
    theme: 'default',
    mode: 'markdown',
    lineNumbers: true,
    autoCloseBrackets: true,
    matchBrackets: true,
    lint: true
  };

  constructor() { }

  ngOnInit(): void {
  }

}
