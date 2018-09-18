#include "example4.h"

int dump(std::pair<int, int> p) {
  std::cout << "first: " << p.first << std::endl;
  std::cout << "second: " << p.second << std::endl;
  std::cout << "dump: ";
  return 0;
}

int dump2(std::pair<std::string, std::string> p) {
  std::cout << "first: " << p.first << std::endl;
  std::cout << "second: " << p.second << std::endl;
  std::cout << "dump: ";
  return 0;
}

int dump3(std::pair<std::string, std::string> p) {
  std::tuple<std::string, std::string> t5 = p;
//  std::get<0>(t5);


  std::cout << "first-t: " << std::get<0>(t5) << std::endl;
  std::cout << "second-t: " << std::get<1>(t5) << std::endl;


  std::cout << "first: " << p.first << std::endl;
  std::cout << "second: " << p.second << std::endl;
  std::cout << "dump: ";
  return 0;
}

std::pair<std::string, std::string> dump4() {
  return std::make_pair("a", "b");
}
