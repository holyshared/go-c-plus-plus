#pragma once

#include "../3rdparty/example/graphic/graphic.hpp"

int Math_succ(int x);

class GraphicPoint {
  private:
    example::graphic::Point point;
  public:
    GraphicPoint(int x, int y);
    void display();
};
