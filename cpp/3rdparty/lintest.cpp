#include "./example/math/math.hpp"
#include "./example/graphic/graphic.hpp"
#include <iostream>

using namespace example;

int math_succ(int x) {
  return math::succ(x);
}

int main() {
  std::cout << math_succ(1);

  graphic::Point point =* new graphic::Point(1, 2);
  point.diplay();

  return 0;
}
