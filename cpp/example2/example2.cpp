#include <iostream>
#include "example2.h"

using namespace std;

Point::Point() {
  x = y = 0;
}

Point::Point(int x_val, int y_val) {
  x = x_val;
  y = y_val;
}

void Point::moveRelative(int x_val, int y_val) {
  x += x_val;
  y += y_val;
}

void Point::display() {
  cout << "x: " << x << " y: " << y << "\n";
}
