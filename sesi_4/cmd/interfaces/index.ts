interface Shape {
  area(): number;
  perimeter(): number;
}

class Rectangle implements Shape {
  constructor(public width: number, public height: number) {}

  area() {
    return this.width * this.height;
  }

  perimeter() {
    return 2 * (this.width + this.height);
  }
}

class Circle implements Shape {
  constructor(public radius: number) {}

  area() {
    return Math.PI * this.radius ** 2;
  }

  perimeter() {
    return 2 * Math.PI * this.radius;
  }
}

const circle_1: Shape = new Circle(5);
const rectangle_1: Shape = new Rectangle(5, 10);
