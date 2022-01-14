#include <iostream>

#include "StringTemplate.h"

int main() {
  const string t =
      "this is the test string\n"
      "a paragraph here, {{ }}\n"
      "{{ custom text }}\n";
  vector<string> variables = {"REPLEASE FIRST", "A NEW paragraph generated"};

  cout << t << endl;
  cout << "Check replace ->" << endl;

  StringTemplate *st = new StringTemplate(t, variables);
  auto a = st->FindPairs();
  for (int i = 0; i < a.size(); i++)
    cout << a[i]->Pos << " : " << a[i]->Length << endl;
  cout << st->ToString() << endl;

  return 0;
}