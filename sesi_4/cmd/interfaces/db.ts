class Product {
  private id: number;
  private name: string;
}

interface ProductRepository {
  save(name: string, price: number): Product;
  deleteById(id: number): void;
  getById(id: number): Product;
  getByName(name: string): Product;
}

class ProductService {
  constructor(private productRepository: ProductRepository) {}

  async createProduct(name: string, price: number) {
    const findProduct = this.productRepository.getByName(name);

    if (findProduct) {
      throw new Error("Product already exists");
    }

    return this.productRepository.save(name, price);
  }
}

class ProductRepositoryPostgres implements ProductRepository {
  executeQuery(sql: string) {
    console.log(`Executing query: ${sql}`);
  }

  save(name: string, price: number) {
    const sql = `INSERT INTO products (name, price) VALUES ('${name}', ${price})`;

    const result = this.executeQuery(sql);

    console.log(`Creating product ${name} with price ${price}`);

    return new Product();
  }

  deleteById(id: number) {
    console.log(`Deleting product with id ${id}`);
  }

  getById(id: number) {
    console.log(`Getting product with id ${id}`);

    return new Product();
  }

  getByName(name: string): Product {
    console.log(`Getting product with name ${name}`);

    return new Product();
  }
}

class ProductRepositoryMongoDB implements ProductRepository {
  save(name: string, price: number): Product {
    throw new Error("Method not implemented.");
  }
  deleteById(id: number): void {
    throw new Error("Method not implemented.");
  }
  getById(id: number): Product {
    throw new Error("Method not implemented.");
  }
  getByName(name: string): Product {
    throw new Error("Method not implemented.");
  }
}

class ProductController {
  constructor(private productService: ProductService) {}

  async createProduct(name: string, price: number) {
    try {
      await this.productService.createProduct(name, price);
    } catch (error) {
      console.error(error);
    }
  }
}

function main() {
  const productRepository: ProductRepository = new ProductRepositoryMongoDB();

  const productService = new ProductService(productRepository);

  const productController = new ProductController(productService);
}
