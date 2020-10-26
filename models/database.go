package models

//DBConnection struct
type DBConnection struct{
	Port int
	DBUser string
	DBPassword string
}

//DBConfig struct
type DBConfig struct{
	DBName string
	Collection string
}