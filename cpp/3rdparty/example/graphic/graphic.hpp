/**
 * graphic
 */

#ifndef EXAMPLE_GRAPHIC_HPP
#define EXAMPLE_GRAPHIC_HPP

namespace example {
  namespace graphic {
    class Point {
      private:
        int x, y;
      public:
        Point(int x, int y);
        void display();
    };
  }
}

#endif
