package infraestructure

import (
	"database/sql"
	"demo/src/domain"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" 
	"github.com/joho/godotenv"        
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL() *MySQL {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, name)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return &MySQL{
		db: db,
	}
}

func (m *MySQL) Save(product domain.Product) (domain.Product, error) {
	result, err := m.db.Exec("INSERT INTO products (name, price) VALUES (?, ?)", product.Name, product.Price)
	if err != nil {
		return domain.Product{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return domain.Product{}, err
	}
	product.Id = int32(id)
	product = domain.Product{Id: int32(id), Name: product.Name, Price: product.Price}

	fmt.Println("Producto guardado en MySQL")
	return product, nil
}

func (m *MySQL) GetAll() []domain.Product {
	rows, err := m.db.Query("SELECT * FROM products")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Price)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, product)
	}
	return products
}

func (m *MySQL) GetByID(id int32) (domain.Product, error) {
	row := m.db.QueryRow("SELECT * FROM products WHERE id = ?", id)
	var product domain.Product
	err := row.Scan(&product.Id, &product.Name, &product.Price)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (m *MySQL) Update(product domain.Product) (domain.Product, error) {
	result, err := m.db.Exec("UPDATE products SET name = ?, price = ? WHERE id = ?", product.Name, product.Price, product.Id)
	if err != nil {
		return domain.Product{}, err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (m *MySQL) Delete(id int32) (int64, error) {
	result, err := m.db.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		return 0, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows, nil
}
