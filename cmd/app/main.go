package main

import (
  "context"
  "fmt"
  "net/url"
  "os"

  //Импортируем библиотеку для работы с postgres
  "github.com/jackc/pgx/v4"
)

func main()  {
  //Создание строки для подключения
  connStr := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable&connect_timeout=%d",
     "postgres",
     url.QueryEscape("db_user"),
     url.QueryEscape("pwd123"),
     "localhost",
     "54320",
     "db_test",
     5)

  //Контекст с отменой
  ctx, _ := context.WithCancel(context.Background())

  //Подключение к базе данных. В случае неуспешного подключения выводим ошибку
  conn, err := pgx.Connect(ctx, connStr)
  if err != nil {
     fmt.Fprintf(os.Stderr, "Connect to database failed: %v\n", err)
     os.Exit(1)
  }
  fmt.Println("Connection OK!")

  //Ping делает пустой запрос к базе данных (";") для проверки наличия фактического соединения
  err = conn.Ping(ctx)
  if err != nil {
     fmt.Fprintf(os.Stderr, "Ping failed: %v\n", err)
     os.Exit(1)
  }
  fmt.Println("Query OK!")

//Закрываем соединение с базой
conn.Close(ctx)

}
