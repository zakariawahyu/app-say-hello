package main

import (
	"fmt"
	go_say_hello "github.com/zakariawahyu/go-say-hello/v2"
)

/**
   ----> Sebelum ada go modules <--------
Saat kita membuat aplikasi, biasanya kita akan menggunakan librabriy atau depedency dari project lain.
Sebelum ada go modules, management untuk depedency sangat sulit dilakukan di golang
Dulu kita harus mengcopy semua code library atau depedency lain ke project kita, hal ini membuat project kita menjadi bengkak karena penuh dengan librabry orang lain.
Atau biasanya orang lain akan kita save di GOPATH directory, hamun hal ini akan sulit jika ternyata bebrapa aplikasi butuh library yang sama  dengan versi yang berbeda
 */

/**
Menambahkan depedency atau library
go get nama module/library
 */

/**
Upgrade depedency atau library
Untuk upgrade depedency ke versi baru ktia bisa ubah isi version tags module di go.mod dan tambahkan perintah go get
Atay kita bisa langsung go get nama module kembali
 */


func main()  {
	// memanggil module say hello yang telah dibuat
	fmt.Println(go_say_hello.SayHello("Zakaria"))
}
