package adapters

import (
	"demo/src/core"
	"demo/src/users/domain/entities"
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

func (mysql *MySQL) Register(user *entities.User) error {
	query:= "INSERT INTO users (user_id, name, lastname, password,role) VALUES (? , ? , ? , ?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, user.Id, user.Name, user.LastName, user.Password, user.Role)
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

            user.Id = int32(lastInsertID)
		} else {
			log.Printf("[MySQL] - Ninguna fila fue afectada.")
		}
	} else {
		log.Printf("[MySQL] - Resultado de la consulta es nil.")
	}
	return nil
}
func (mysql *MySQL) Update(id int32, name string, lastname string, password string, role int32) error {
	query := "UPDATE users SET name =?, lastname =?, password =?, role =? WHERE user_id =?"
    result, err := mysql.conn.ExecutePreparedQuery(query, name, lastname, password, role, id)
    if err!= nil {
        fmt.Println(err)
        return err
    }
    if result!= nil {
        rowsAffected, _ := result.RowsAffected()
        if rowsAffected == 1 {
            log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)
        } else {
            log.Printf("[MySQL] - Ninguna fila fue afectada.")
        }
    } else {
        log.Printf("[MySQL] - Resultado de la consulta es nil.")
    }
    return nil
}
func (mysql *MySQL) GetAll() ([]*entities.User, error) {
	query := "SELECT * FROM users WHERE deleted = 0"
    rows, err := mysql.conn.FetchRows(query)
    if err!= nil {
        fmt.Println(err)
        return nil, err
    }
    defer rows.Close()

    var users []*entities.User
    var deleted bool
    for rows.Next() {
        user := entities.User{}
        err := rows.Scan(&user.Id, &user.Name, &user.LastName, &user.Password, &user.Role, &deleted)
        if err!= nil {
            fmt.Println(err)
            return nil, err
        }
        users = append(users, &user)
    }
    return users, nil
}
func (mysql *MySQL) Delete(id int32) error {
	query := "UPDATE users SET deleted = 1 WHERE user_id = ?"
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