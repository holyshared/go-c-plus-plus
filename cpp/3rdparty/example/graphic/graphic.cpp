#include <iostream>
#include "graphic.hpp"

namespace example {
  namespace graphic {
    using namespace std;

    Point::Point(int x_val, int y_val) {
      x = x_val;
      y = y_val;
    }

    void Point::display() {
      cout << "x: " << x << " y: " << y << "\n";
    }
  }
}
