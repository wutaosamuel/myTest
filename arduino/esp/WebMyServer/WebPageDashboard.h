// WebPageDashboard.h
// create dashboard html5 page

#ifndef WEBPAGEDASHBOARD_H 
#define WEBPAGEDASHBOARD_H

#include <string>
#include <cstring>

using namespace std;

class WebPageDashboard
{
public:
  WebPageDashboard();
  ~WebPageDashboard();

  string printHtml();

  // setter 
  void setTitle(string s) { title_ = s; }
  // getter
  string getTitle() { return title_; }

  // append
  void nemuItemAdd(string href, string name);
  void mainItemAdd(string name);
  // remove
  void menuItemDel(string name);
  void mainItemDel(string name);


private:
  string title_;
  string menuName_;
  string authorName_;

  string htmlTitle();
  string htmlMenu();
  string htmlOverlay();
  string htmlPageContent();
  string htmlHeader();
  string htmlMain();
  string htmlPagination();
  string htmlFooter();
  string htmlScript();
};

#endif 