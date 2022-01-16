// StringTemplate.h
#ifndef STRINGTEMPLATE_H
#define STRINGTEMPLATE_H

#include <iostream>
#include <string>
#include <vector>

#define SIGNLEFT "{{"
#define SIGNRIGHT "}}"

using namespace std;

class SignPair {
 public:
  size_t Pos;
  size_t Length;

  SignPair(size_t pos, size_t length);
  ~SignPair();
  inline size_t NextPos(size_t shift_pos);
};

class StringTemplate {
 public:
  const string& Template;
  const vector<string>& Variables;

 private:
  string signL = SIGNLEFT;
  string signR = SIGNRIGHT;

 public:
  StringTemplate(const string& templt, const vector<string>& variables);
  ~StringTemplate();
  inline void SetSign(string left, string right);
  vector<SignPair*> FindPairs();
  string ToString();
  string ToBigString();

 private:
  SignPair* findPair(size_t);
  SignPair* findP(string, size_t);
};
#endif