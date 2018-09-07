class Point {
  private:
    int x, y;

  public:
    Point();
    Point(int x, int y);
    void moveRelative(int x, int y);
    void display();
};
