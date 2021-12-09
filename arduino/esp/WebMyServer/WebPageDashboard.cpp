// WebPageDashboard.cpp

#include "WebPageDashboard.h"

WebPageDashboard::WebPageDashboard()
{
  title_ = "My Dashboard";
  menuName_ = "PROTFOLIO";
  authorName_ = "Template by W3.CSS";
}

WebPageDashboard::~WebPageDashboard() {}

string WebPageDashboard::htmlTitle()
{
  string buf = "";
  buf += "<!DOCTYPE html>\n";
  buf += "<html>\n";
  buf += "<title>"+title_+"</title>\n";
  buf += "<meta charset=\"UTF-8\">\n";
  buf += "<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n";
  buf += "<link rel=\"stylesheet\" href=\"https://www.w3schools.com/w3css/4/w3.css\">\n";
  buf += "<link rel=\"stylesheet\" href=\"https://fonts.googleapis.com/css?family=Raleway\">\n";
  buf += "<link rel=\"stylesheet\" href=\"https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css\">\n";
  buf += "<style>\n";
  buf += "body,h1,h2,h3,h4,h5,h6 {font-family: \"Raleway\", sans-serif}\n";
  buf += "</style>\n";
  buf += "<body class=\"w3-light-grey w3-content\" style=\"max-width:1600px\">\n\n";
  return buf;
}

string WebPageDashboard::htmlMenu()
{
  string buf = "";
  buf += "<!-- Sidebar/menu -->\n";
  buf += "<nav class=\"w3-sidebar w3-collapse w3-white w3-animate-left\" style=\"z-index:3;width:300px;\" id=\"mySidebar\"><br>\n";
  buf += "  <div class=\"w3-container\">\n";
  buf += "    <a href=\"#\" onclick=\"w3_close()\" class=\"w3-hide-large w3-right w3-jumbo w3-padding w3-hover-grey\" title=\"close menu\">\n";
  buf += "      <i class=\"fa fa-remove\"></i>\n";
  buf += "    </a>\n";
  buf += "    <img src=\"images/avatar_g2.jpg\" style=\"width:45%;\" class=\"w3-round\"><br><br>\n";
  buf += "    <h4><b>"+menuName_+"</b></h4>\n";
  buf += "    <p class=\"w3-text-grey\">"+authorName_+"</p>\n";
  buf += "  </div>\n";
  buf += "  <div class=\"w3-bar-block\">\n";
  buf += "    <a href=\"#Home\" onclick=\"w3_close()\" class=\"w3-bar-item w3-button w3-padding w3-text-teal\"><i class=\"fa fa-th-large fa-fw w3-margin-right\"></i>Home</a> \n";
  buf += "    <a href=\"#Group 0\" onclick=\"w3_close()\" class=\"w3-bar-item w3-button w3-padding\"><i class=\"fa fa-user fa-fw w3-margin-right\"></i>Group 0</a> \n";
  buf += "  </div>\n";
  buf += "  <div class=\"w3-panel w3-large\">\n";
  buf += "    <i class=\"fa fa-facebook-official w3-hover-opacity\"></i>\n";
  buf += "    <i class=\"fa fa-instagram w3-hover-opacity\"></i>\n";
  buf += "    <i class=\"fa fa-snapchat w3-hover-opacity\"></i>\n";
  buf += "    <i class=\"fa fa-pinterest-p w3-hover-opacity\"></i>\n";
  buf += "    <i class=\"fa fa-twitter w3-hover-opacity\"></i>\n";
  buf += "    <i class=\"fa fa-linkedin w3-hover-opacity\"></i>\n";
  buf += "  </div>\n";
  buf += "</nav>\n\n";
  return buf;
}

string WebPageDashboard::htmlOverlay()
{
  string buf = "";
  buf += "<!-- Overlay effect when opening sidebar on small screens -->\n";
  buf += "<div class=\"w3-overlay w3-hide-large w3-animate-opacity\" onclick=\"w3_close()\" style=\"cursor:pointer\" title=\"close side menu\" id=\"myOverlay\"></div>\n\n";
  return buf;
}

string WebPageDashboard::htmlPageContent()
{
  string buf = "";
  buf += "<!-- !PAGE CONTENT! -->\n";
  buf += "<div class=\"w3-main\" style=\"margin-left:300px\">\n\n";
  return buf;
}

string WebPageDashboard::htmlHeader()
{
  string buf = "";
  buf += " <!-- Header -->\n";
  buf += " <header id=\"portfolio\">\n";
  buf += "    <a href=\"#\"><img src=\"images/avatar_g2.jpg\" style=\"width:65px;\" class=\"w3-circle w3-right w3-margin w3-hide-large w3-hover-opacity\"></a>\n";
  buf += "    <span class=\"w3-button w3-hide-large w3-xxlarge w3-hover-text-grey\" onclick=\"w3_open()\"><i class=\"fa fa-bars\"></i></span>\n";
  buf += "    <div class=\"w3-container\">\n";
  buf += "    <h1><b>Dashboard</b></h1>\n";
  buf += "    <div class=\"w3-section w3-bottombar w3-padding-16\">\n";
  buf += "    <span class=\"w3-margin-right\">Filter:</span> \n";
  buf += "    <button class=\"w3-button w3-black\">ALL</button>\n";
  buf += "    <button class=\"w3-button w3-white\"><i class=\"fa fa-diamond w3-margin-right\"></i>Design</button>\n";
  buf += "    <button class=\"w3-button w3-white w3-hide-small\"><i class=\"fa fa-photo w3-margin-right\"></i>Photos</button>\n";
  buf += "    <button class=\"w3-button w3-white w3-hide-small\"><i class=\"fa fa-map-pin w3-margin-right\"></i>Art</button>\n";
  buf += "  </div>\n";
  buf += "  </div>\n";
  buf += "  </header>\n\n";
  return buf;
}

string WebPageDashboard::htmlMain()
{
  // TODO: generate more
  string buf = "";
  buf += "  <!-- First Photo Grid-->\n";
  buf += "  <div class=\"w3-row-padding\">\n";
  buf += "    <div class=\"w3-third w3-container w3-margin-bottom\">\n";
  buf += "      <img src=\"images/mountains.jpg\" alt=\"Norway\" style=\"width:100%\" class=\"w3-hover-opacity\">\n";
  buf += "      <div class=\"w3-container w3-white\">\n";
  buf += "        <p><b>Lorem Ipsum</b></p>\n";
  buf += "        <p>Praesent tincidunt sed tellus ut rutrum. Sed vitae justo condimentum, porta lectus vitae, ultricies congue gravida diam non fringilla.</p>\n";
  buf += "      </div>\n";
  buf += "    </div>\n\n";
  buf += "  </div>\n\n";
  return buf;
}

string WebPageDashboard::htmlPagination()
{
  string buf = "";
  buf += "  <!-- Pagination -->\n";
  buf += "  <div class=\"w3-center w3-padding-32\">\n";
  buf += "    <div class=\"w3-bar\">\n";
  buf += "      <a href=\"#\" class=\"w3-bar-item w3-button w3-hover-black\">«</a>\n";
  buf += "      <a href=\"#\" class=\"w3-bar-item w3-black w3-button\">1</a>\n";
  buf += "      <a href=\"#\" class=\"w3-bar-item w3-button w3-hover-black\">2</a>\n";
  buf += "      <a href=\"#\" class=\"w3-bar-item w3-button w3-hover-black\">3</a>\n";
  buf += "      <a href=\"#\" class=\"w3-bar-item w3-button w3-hover-black\">4</a>\n";
  buf += "      <a href=\"#\" class=\"w3-bar-item w3-button w3-hover-black\">»</a>\n";
  buf += "    </div>\n";
  buf += "  </div>\n\n";
  return buf;
}

string WebPageDashboard::htmlFooter()
{
  string buf = "";
  buf += "  <footer class=\"w3-container w3-padding-32 w3-dark-grey\">\n";
  buf += "  <div class=\"w3-row-padding\">\n";
  buf += "    <div class=\"w3-third\">\n";
  buf += "      <h3>FOOTER</h3>\n";
  buf += "      <p>Praesent tincidunt sed tellus ut rutrum. Sed vitae justo condimentum, porta lectus vitae, ultricies congue gravida diam non fringilla.</p>\n";
  buf += "      <p>Powered by <a href=\"https://www.w3schools.com/w3css/default.asp\" target=\"_blank\">w3.css</a></p>\n";
  buf += "    </div>\n\n";
  buf += "    <div class=\"w3-third\">\n";
  buf += "      <h3>BLOG POSTS</h3>\n";
  buf += "      <ul class=\"w3-ul w3-hoverable\">\n";
  buf += "        <li class=\"w3-padding-16\">\n";
  buf += "          <img src=\"images/workshop.jpg\" class=\"w3-left w3-margin-right\" style=\"width:50px\">\n";
  buf += "          <span class=\"w3-large\">Lorem</span><br>\n";
  buf += "          <span>Sed mattis nunc</span>\n";
  buf += "        </li>\n";
  buf += "        <li class=\"w3-padding-16\">\n";
  buf += "          <img src=\"images/gondol.jpg\" class=\"w3-left w3-margin-right\" style=\"width:50px\">\n";
  buf += "          <span class=\"w3-large\">Ipsum</span><br>\n";
  buf += "          <span>Praes tinci sed</span>\n";
  buf += "        </li> \n";
  buf += "      </ul>\n";
  buf += "    </div>\n\n";
  buf += "    <div class=\"w3-third\">\n";
  buf += "      <h3>POPULAR TAGS</h3>\n";
  buf += "      <p>\n";
  buf += "        <span class=\"w3-tag w3-black w3-margin-bottom\">Travel</span> <span class=\"w3-tag w3-grey w3-small w3-margin-bottom\">New York</span> <span class=\"w3-tag w3-grey w3-small w3-margin-bottom\">London</span>\n";
  buf += "        <span class=\"w3-tag w3-grey w3-small w3-margin-bottom\">IKEA</span> <span class=\"w3-tag w3-grey w3-small w3-margin-bottom\">NORWAY</span> <span class=\"w3-tag w3-grey w3-small w3-margin-bottom\">DIY</span>\n";
  buf += "        <span class=\"w3-tag w3-grey w3-small w3-margin-bottom\">Ideas</span> <span class=\"w3-tag w3-grey w3-small w3-margin-bottom\">Baby</span> <span class=\"w3-tag w3-grey w3-small w3-margin-bottom\">Family</span>\n";
  buf += "        <span class=\"w3-tag w3-grey w3-small w3-margin-bottom\">News</span> <span class=\"w3-tag w3-grey w3-small w3-margin-bottom\">Clothing</span> <span class=\"w3-tag w3-grey w3-small w3-margin-bottom\">Shopping</span>\n";
  buf += "        <span class=\"w3-tag w3-grey w3-small w3-margin-bottom\">Sports</span> <span class=\"w3-tag w3-grey w3-small w3-margin-bottom\">Games</span>\n";
  buf += "      </p>\n";
  buf += "    </div>\n\n";
  buf += "  </div>\n";
  buf += "  </footer>\n";
  buf += "  <div class=\"w3-black w3-center w3-padding-24\">Powered by <a href=\"https://www.w3schools.com/w3css/default.asp\" title=\"W3.CSS\" target=\"_blank\" class=\"w3-hover-opacity\">w3.css</a></div>\n";
  buf += "<!-- End page content -->\n";
  buf += "</div>\n\n";
  return buf;
}

string WebPageDashboard::htmlScript()
{
  string buf = "";
  buf += "<script>\n";
  buf += "// Script to open and close sidebar\n";
  buf += "function w3_open() {\n";
  buf += "    document.getElementById(\"mySidebar\").style.display = \"block\";\n";
  buf += "    document.getElementById(\"myOverlay\").style.display = \"block\";\n";
  buf += "}\n\n";
  buf += "function w3_close() {\n";
  buf += "    document.getElementById(\"mySidebar\").style.display = \"none\";\n";
  buf += "    document.getElementById(\"myOverlay\").style.display = \"none\";\n";
  buf += "}\n";
  buf += "</script>\n\n";
  buf += "</body>\n";
  buf += "</html>\n";
  return buf;
}

string WebPageDashboard::printHtml()
{
  string buf = "";
  buf += htmlTitle();
  buf += htmlMenu();
  buf += htmlOverlay();
  buf += htmlPageContent();
  buf += htmlHeader();
  buf += htmlMain();
  buf += htmlPagination();
  buf += htmlFooter();
  buf += htmlScript();
  return buf;
}