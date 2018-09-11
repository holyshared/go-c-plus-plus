#include "../3rdparty/example/math/math.hpp"
#include "../3rdparty/example/graphic/graphic.hpp"
#include "example3.h"

using namespace example;

int Math_succ(int x) {
  return math::succ(x);
}

GraphicPoint::GraphicPoint(int x_v, int y_v) : point(x_v, y_v) {
}

void GraphicPoint::display() {
  point.display();
}
