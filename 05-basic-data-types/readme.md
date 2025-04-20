# 🧑‍💻 Golang'da Temel Veri Tipleri:

## 🌍 Programlamada Veri Tiplerinin Önemi

Her programlama dilinde veri tipleri, dilin en önemli yapı taşlarından biridir. Veri tipleri, bir değişkenin ne tür verileri tutabileceğini belirler ve bu veriler üzerinde işlem yapabilmek için uygun türde veri tiplerinin kullanılması gerekir. Golang da güçlü ve statik bir tür sistemine sahiptir. Bu yazıda, Go dilinde en çok kullanılan temel veri tiplerini inceleyeğiz.

## 🏗️ Golang'da Temel Veri Tipleri

Golang, birçok temel veri tipi sunar. Bu veri tiplerini doğru bir şekilde anlamak, hatasız ve verimli bir Go kodu yazmak için kritik öneme sahiptir. İşte Golang'da en yaygın kullanılan temel veri tipleri:

### 1. Sayısal Veri Tipleri (int, float, complex)

| Veri Tipi   | Açıklama                                | Boyut (Bit)    | Değer Aralığı                                                               | Örnek Kullanım                   |
|-------------|-----------------------------------------|----------------|-------------------------------------------------------------------------------|----------------------------------|
| int         | İşaretli tam sayılar, platforma bağlı büyüklükte | 32 veya 64 bit | Platforma bağlı (genellikle -2^31 - 2^31-1)                                  | `var x int = 42`                 |
| int8        | 8 bit işaretli tam sayılar              | 8 bit          | -128 ile 127                                                                 | `var a int8 = 100`               |
| int16       | 16 bit işaretli tam sayılar             | 16 bit         | -32,768 ile 32,767                                                            | `var b int16 = -500`             |
| int32       | 32 bit işaretli tam sayılar             | 32 bit         | -2,147,483,648 ile 2,147,483,647                                              | `var c int32 = 5000`             |
| int64       | 64 bit işaretli tam sayılar             | 64 bit         | -9,223,372,036,854,775,808 ile 9,223,372,036,854,775,807                    | `var d int64 = 10000000000`      |
| uint        | İşaretsiz tam sayılar, platforma bağlı büyüklükte | 32 veya 64 bit | 0 ile platforma bağlı maksimum değer (genellikle 2^32-1 veya 2^64-1)       | `var e uint = 42`                |
| uint8       | 8 bit işaretsiz tam sayılar             | 8 bit          | 0 ile 255                                                                    | `var f uint8 = 200`              |
| uint16      | 16 bit işaretsiz tam sayılar            | 16 bit         | 0 ile 65,535                                                                 | `var g uint16 = 40000`           |
| uint32      | 32 bit işaretsiz tam sayılar            | 32 bit         | 0 ile 4,294,967,295                                                          | `var h uint32 = 1000000000`      |
| uint64      | 64 bit işaretsiz tam sayılar            | 64 bit         | 0 ile 18,446,744,073,709,551,615                                            | `var i uint64 = 100000000000000`|
| float32     | 32 bit kayan nokta sayılar (kesirli sayılar) | 32 bit         | ±1.5 × 10^−45 ile ±3.4 × 10^38                                              | `var j float32 = 3.14`           |
| float64     | 64 bit kayan nokta sayılar (kesirli sayılar) | 64 bit         | ±5.0 × 10^−324 ile ±1.7 × 10^308                                            | `var k float64 = 19.99`          |
| complex64   | 64 bit karmaşık sayılar (gerçek + sanal kısmı) | 64 bit (2 × 32 bit) | Gerçek ve sanal kısımlar için float32 aralığı                               | `var l complex64 = 1 + 2i`       |
| complex128  | 128 bit karmaşık sayılar (gerçek + sanal kısmı) | 128 bit (2 × 64 bit) | Gerçek ve sanal kısımlar için float64 aralığı                               | `var m complex128 = 2 + 3i`      |


Go dilinde sayısal veri tipleri, tam sayılar ve kesirli sayılar için kullanılır. Ayrıca karmaşık sayılarla çalışmak da mümkündür.

#### a. int, int8, int16, int32, int64

int veri tipi, tam sayılar için kullanılır. Go, 32-bit veya 64-bit platformlara göre uygun bir int türü seçer. int türü, platforma bağlı olarak farklı boyutlarda olabilir. Ancak, 8, 16, 32 ve 64 bit'lik versiyonları da mevcuttur.

- int: Genel tam sayı türüdür.
- int8, int16, int32, int64: Belirli bit uzunluklarına sahip tam sayılardır.

Örnek;

```go
var x int = 42
var y int64 = 10000000000
fmt.Println(x, y) // Output: 42 10000000000
```

#### b. uint, uint8, uint16, uint32, uint64

Bu türler, işaretsiz tam sayılar için kullanılır (sadece pozitif sayılar).

- uint: Platform bağımlı işaretsiz tam sayıdır.
- uint8, uint16, uint32, uint64: Farklı bit uzunluklarına sahip işaretsiz tam sayılardır.

Örnek: 

```go
var a uint = 50
var b uint8 = 300
fmt.Println(a, b) // Output: 50 300
```

Peki kuralların dışına çıkarsak, yani yukarıdaki tabloya uygun değerler vermezsek ne olur. Örneğin; `uint` için tabloda 0-255 arasında bir değer alır olarak gösteriyor biz buna 256 değeri verirsek ne olur?

```go
var a uint = 256
fmt.Println(a) // cannot use 256 (untyped int constant) as uint8 value in variable declaration (overflows)
```

*cannot use 256 (untyped int constant) as uint8 value in variable declaration (overflows)* şeklinde hata alırız. Yani burada diyor ki; *Sen `uint9` olarak değişkeni tanımladın ve ben sana hafızamda `8 bitlik` bir yer ayırdım ve sen 256 değeri atayarak bu alanı aşıyorsun*

#### c. float32, float64

float32 ve float64 veri tipleri, kesirli sayıları (ondalıklı sayılar) saklamak için kullanılır. float64 daha hassas ve yaygın olarak tercih edilir.

Örnek:

```go
var price float64 = 19.99
fmt.Println(price) // Output: 19.99
```

#### d. complex64, complex128
Bu veri tipleri, karmaşık sayılar için kullanılır. Bir karmaşık sayı, bir gerçek ve bir sanal sayıdan oluşur.

Örnek:

```go
var c complex128 = 1 + 2i
fmt.Println(c) // Output: (1+2i)
```

### 2. Metinsel Veri Tipleri (string)
Go dilinde metin (string) veri tipi, karakter dizilerini temsil eder. String'ler tırnak işaretleriyle ("" veya '') çevrelenir.

Örnek:

```go
var greeting string = "Merhaba, Golang!"
fmt.Println(greeting) // Output: Merhaba, Golang!
```

String'ler immutable'dır, yani bir string değeri oluşturulduktan sonra değiştirilmesi mümkün değildir. Ancak, yeni bir string oluşturmak mümkündür.

### 3. Mantıksal Veri Tipi (bool)
Go dilinde bool veri tipi, doğru (true) veya yanlış (false) değerlerini tutar. Genellikle koşul ifadelerinde kullanılır.

Örnek:

```go
var isGoFun bool = true
fmt.Println(isGoFun) // Output: true
```

#### 4. Boş Değerler ve nil (nil)
Go dilinde bazı veri tipleri boş bir değer alabilir. Örneğin, işaretçi türleri ve arayüzler (interface{}) gibi türler, nil değerini alabilirler.

Örnek:

```go
var ptr *int
fmt.Println(ptr) // Output: <nil>
```
