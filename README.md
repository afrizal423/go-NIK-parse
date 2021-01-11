<p align="right">
بِسْــــــــــــــمِ اللَّهِ الرَّحْمَنِ الرَّحِيم 
</p>


# NIK Parse With Go
<a href="https://afrizalmy.com"><img src="https://img.shields.io/badge/license-MIT-lightgrey" alt="me"></a>

Parse Nomor Induk Kependudukan (NIK) KTP Menggunakan Golang. Dengan mengetahui nomor NIK bisa mendapatkan beberapa informasi seperti pada [hasil output](#output). Cara kerjanya seperti kita mengekstrak perdigit dari NIK.

## Install
```
go get github.com/afrizal423/go-NIK-parse/ParseNIK
```

## Example
```
package main

import (
	"encoding/json"
	"fmt"
	"github.com/afrizal423/go-NIK-parse/ParseNIK"
)

func main() {
	tesnik := "3307114404920004"
	coba := ParseNIK.GetdataNIK(tesnik)
	data, _ := json.Marshal(coba)
	fmt.Println(string(data))
}

```
## Output
```
{
    "nik":"3307114404920004",
    "jenis_kelamin":"PEREMPUAN",
    "tanggal_lahir":"1992-04-04",
    "provinsi":"JAWA TENGAH",
    "kabupaten/kota":"KAB. WONOSOBO",
    "kecamatan":"MOJOTENGAH",
    "kodepos":"56351",
    "uniqcode":"0004",
    "status":"Sukses",
    "pesan":""
}
```
Output bertipe data struct, jadi bisa dilanjutkan dengan mengekstrak tiap data atau langsung simpan ke database.