// StringTemplate.cpp
#include "StringTemplate.h"

StringTemplate::StringTemplate(string templt, vector<string> variables) {
  this->Set(templt, variables);
}

StringTemplate::~StringTemplate() {}

inline void StringTemplate::Set(string templt, vector<string> variables) {
  this->Template = templt;
  this->Variables = variables;
}

inline void StringTemplate::SetSign(string left, string right) {
  this->signL = left != "" ? left : SIGNLEFT;
  this->signR = right != "" ? right : SIGNRIGHT;
}

vector<SignPair*> StringTemplate::FindPairs() {
  vector<SignPair*> result;
  size_t leftLength = this->signL.length();
  size_t pos = 0;

  for (;;) {
    SignPair* pair = this->findPair(pos);
    if (pair == nullptr) return result;
    result.push_back(pair);
    pos = pair->NextPos(leftLength);
  }
  return result;
}

string StringTemplate::ToString() {
  // check the No of template found is the same as the Variables
  vector<SignPair*> signPairs = this->FindPairs();
  if (signPairs.size() != this->Variables.size()) return "";

  // replace template string
  // NOTE: can be directly change this->Template to save memory
  string result = this->Template;
  int shiftPos = 0;
  for (int i = 0; i < signPairs.size(); i++) {
    result.replace(signPairs[i]->Pos + shiftPos, signPairs[i]->Length,
                   this->Variables[i]);
    shiftPos += this->Variables[i].length() - signPairs[i]->Length;
  }
  return result;
}

SignPair* StringTemplate::findPair(size_t pos) {
  size_t found = this->Template.find(this->signL, pos);
  if (found == string::npos) return nullptr;

  size_t nextFound =
      this->Template.find(this->signL, found + this->signL.length());
  size_t pairFound =
      this->Template.find(this->signR, found + this->signR.length());
  if (pairFound == string::npos) return nullptr;
  if (nextFound != string::npos) {
    if (pairFound > nextFound) {
      return this->findPair(found + this->signL.length());
    }
  }

  SignPair* result =
      new SignPair(found, pairFound + this->signR.length() - found);
  return result;
}

SignPair::SignPair(size_t pos, size_t length) {
  this->Pos = pos;
  this->Length = length;
}
SignPair::~SignPair() {}

inline size_t SignPair::NextPos(size_t shift_pos) {
  return this->Pos + shift_pos;
}