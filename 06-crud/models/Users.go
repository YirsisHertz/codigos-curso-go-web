package models

import (
	"crud/config"

	"golang.org/x/crypto/bcrypt"
)

type Usuario struct {
	Id       int
	Nombre   string
	Correo   string
	Password string
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

func CheckHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

func CreateUser(nombre string, correo string, password string) {
	conexion := config.Conexion()

	query, err := conexion.Prepare("INSERT INTO usuario (nombre, correo, password, activo) VALUE (?, ?, ?, ?)")

	if err != nil {
		panic("Error al crear usuario")
	}

	securePassword, err := HashPassword(password)

	if err != nil {
		panic("Error al cifrar contraseña")
	}

	query.Exec(nombre, correo, securePassword, 1)

	conexion.Close()
}

func DeleteUser(id string) {
	conexion := config.Conexion()

	query, err := conexion.Prepare("UPDATE usuario SET activo = ? WHERE id = ?")

	if err != nil {
		panic("Error al eliminar usuario")
	}

	query.Exec(0, id)

	conexion.Close()
}

func UpdateUser(id, nombre, correo, password string) {
	conexion := config.Conexion()

	query, err := conexion.Prepare("UPDATE usuario SET nombre = ?, correo = ?, password = ? WHERE id = ?")

	if err != nil {
		panic("Error al actualizar usuario")
	}

	securePassword, err := HashPassword(password)

	if err != nil {
		panic("Error al cifrar contraseña")
	}

	query.Exec(nombre, correo, securePassword, id)

	conexion.Close()
}

func ReadUser(id string) Usuario {
	conexion := config.Conexion()

	row := conexion.QueryRow("SELECT id, nombre, correo, password FROM usuario WHERE activo = 1 AND id = ?", id)

	conexion.Close()

	var usuario Usuario

	row.Scan(&usuario.Id, &usuario.Nombre, &usuario.Correo, &usuario.Password)

	return usuario
}

func ReadUsers() []Usuario {
	conexion := config.Conexion()

	rows, err := conexion.Query("SELECT id, nombre, correo FROM usuario WHERE activo = 1")

	if err != nil {
		panic("Error al leer usuarios")
	}

	conexion.Close()

	var usuarios []Usuario

	for rows.Next() {
		var usuario Usuario

		rows.Scan(&usuario.Id, &usuario.Nombre, &usuario.Correo)
		usuarios = append(usuarios, usuario)
	}

	return usuarios
}
