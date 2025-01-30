package adapters

import (
	"demo/src/core"
	"demo/src/products/domain/entities"
	"fmt"
	"log"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

func NewMySQL() (*MySQL, error) {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}, nil
}

func (mysql *MySQL) Save(product *entities.Product) error {
	query := "INSERT INTO products (name, price, description, deleted, stock, user_id_fk) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, product.Name, product.Price, product.Descripcion, product.Deleted, product.Stock, product.User_id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if result != nil {
		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 1 {
			log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)
			lastInsertID, err := result.LastInsertId()
            if err != nil {
                fmt.Println(err)
                return err
            }

            product.ID = int32(lastInsertID)
		} else {
			log.Printf("[MySQL] - Ninguna fila fue afectada.")
		}
	} else {
		log.Printf("[MySQL] - Resultado de la consulta es nil.")
	}
	return nil
}

func (mysql *MySQL) GetAll() ([]*entities.Product, error) {
	query := "SELECT * FROM products WHERE deleted = 0"
	rows, err := mysql.conn.FetchRows(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entities.Product
	for rows.Next() {
		var product_id int32
		var name string
		var price float32
		var description string
		var deleted bool
		var stock int32
		var user_id int32
		if err := rows.Scan(&product_id, &name, &price, &description, &deleted, &stock, &user_id); err != nil {
			fmt.Println("error al escanear la fila: %w", err)
		}
		product := &entities.Product{
			ID:          product_id,
			Name:        name,
			Price:       price,
			Descripcion: description,
			Deleted:     deleted,
			Stock:       stock,
			User_id: user_id,
		}
		products = append(products, product)
		
		fmt.Printf("ID: %d, Name: %s, Price: %f, Descripcion: %s, Stock: %d, Deleted: %v \n", product_id, name, price, description, stock, deleted)
	}
	return products, nil
}

func (mysql *MySQL) Delete(id int32) error {
	query := "UPDATE products SET deleted = 1 WHERE product_id = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		log.Printf("[MySQL] - Error al ejecutar la consulta: %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("[MySQL] - Error al obtener las filas afectadas: %v", err)
		return err
	}
	if rowsAffected == 0 {
		log.Printf("[MySQL] - Ninguna fila fue afectada. Producto con ID %d no encontrado.", id)
		return fmt.Errorf("producto con ID %d no encontrado", id)
	}

	log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)
	return nil
}

func (mysql *MySQL) Update(id int32,name string, price float32, descripcion string, stock int32) error {
	query := "UPDATE products SET name = ?, price = ?, description = ?, stock = ? WHERE product_id = ?"
    result, err := mysql.conn.ExecutePreparedQuery(query, name, price, descripcion, stock, id)
    if err != nil {
        log.Printf("[MySQL] - Error al ejecutar la consulta: %v", err)
        return err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        log.Printf("[MySQL] - Error al obtener las filas afectadas: %v", err)
        return err
    }
	if rowsAffected == 0 {
        log.Printf("[MySQL] - Ninguna fila fue afectada. Producto con ID %d no encontrado.", id)
        return fmt.Errorf("producto con ID %d no encontrado", id)
    }

    log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)

	return nil
}