#include <example/math/math.h>
#include <iostream>

using namespace example;

int math_succ(int x) {
  return math::succ(x);
}

int main() {
  std::cout << math_succ(1);
  return 0;
}
