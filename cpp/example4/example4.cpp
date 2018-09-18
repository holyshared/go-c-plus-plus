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
