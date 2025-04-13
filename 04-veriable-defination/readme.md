# 🐹 Programlamada Değişkenler ve Go (Golang) Diline Özgü Tanımlama Yöntemleri

## 🌍 Genel Programlama Dillerinde Değişkenler

Değişkenler, tüm programlama dillerinde veri saklamak amacıyla kullanılan bellek alanlarıdır. Değişkenlerin türleri, programcıya veriyle ilgili işlemleri yapmak için gereken esnekliği sağlar. Çoğu dilde, bir değişkenin tanımlanması ve kullanılabilmesi için tür belirtmesi veya başlangıç değeri verilmesi gerekir.

### Değişken Tanımlama Genel Kuralları:

- Tür Belirtme: Çoğu dilde değişken tanımlarken, veri tipini belirtmek gerekir (örneğin, int, float, string).
- İsim Verme: Değişkenler anlamlı bir isimle tanımlanmalıdır. İyi bir isimlendirme, kodun okunabilirliğini artırır.
- Sıfır Değerler: Çoğu dil, değişkenin ilk değerinin verilmediği durumlarda bir "sıfır değer" kullanır. Örneğin, int türünde bir değişkenin sıfır değeri 0’dır.
- Değişken Türünü Sonradan Belirleme: Bazı dillerde (Go dışında) değişkenin türü, değer atandıktan sonra otomatik olarak belirlenir (örneğin Python, JavaScript gibi dinamik dillerde).

Örnek:

JavaScript'te bir değişkenin tanımlanması:

```js
const name = 'Yummy';
const age = '25';
```
Bu dillerde tür belirtmek gerekmez, çünkü Python otomatik olarak türü belirler.

<hr>

## 🏗️ Go Diline Özgü Değişken Tanımlama Yöntemleri

Go dilinde değişkenler, diğer dillerden farklı olarak çok katı bir tür sistemine sahiptir. Her değişkenin türü açıkça belirtilmelidir, ya da kısa tanımlama operatörü (:=) ile türü Go tarafından otomatik olarak çıkarılabilir. Go'da değişkenler fonksiyonlar içinde ya da dışında farklı şekillerde tanımlanabilir.

### 1. `var` Anahtar Kelimesi ile Değişken Tanımlama

Go dilinde en yaygın kullanılan yöntemlerden biri, `var` anahtar kelimesi ile değişken tanımlamaktır. Bu yöntemle, türü açıkça belirtiriz.

```go
var name string
name = "Jeremy"
```

Burada öncelikle bir değişken oluşturup tipini tanımlıyoruz, daha sonra da değişkene bir değer atıyoruz.
Yine bu örnekte isim değişkeni string türünde tanımlandı. Eğer isim değişkenine bir değer verilmezse, Go bunu sıfır değer (zero value) ile başlatır. Yani, bir string değişkeni için bu değer boş bir string olur ("").

🎯 Örnek:
```go
var age int
fmt.Println(age) // Output: 0 (sıfır değer)
```

```go
var name string
fmt.Println(name) // Output: "" (sıfır değer)
```


Bu yöntemin avantajı, özellikle uzun fonksiyonlarda veya birden fazla değişken tanımlarken türleri net bir şekilde belirtmekten geçer.

<hr>

### 2. Kısa Tanımlama `:=` ile Değişken Tanımlama

Go dilinde, özellikle fonksiyonlar içinde, kısa tanımlama `:=` kullanarak değişkenleri tanımlayabilirsiniz. Bu operatör, tür belirtmeden otomatik olarak değişkenin türünü çıkarır.

```go
name := "Odd"
```

Bu tanımda isim değişkeni, değeri olan *"Odd"* stringi üzerinden otomatik olarak bir string türü alır.

#### Önemli Not: Kısa Tanımlama Yalnızca Fonksiyon İçinde Geçerlidir!

`:=` yalnızca fonksiyon içinde geçerlidir. Dışında bu şekilde tanımlama yapılamaz. Eğer fonksiyon dışında bir değişken tanımlamak istiyorsanız, `var` anahtar kelimesini kullanmalısınız.

🎯 Örnek:

```go
func main() {
    name := "Ulrich" // Geçerli, çünkü fonksiyon içindeyiz
    fmt.Println(name)
}

name := "William" // Hata! Kısa tanımlama sadece fonksiyon içinde kullanılabilir.
```

Yukarıdaki örnekte, `name := "Ulrich"` ifadesi fonksiyon içinde olduğu için geçerlidir. Ancak `name := "William"` dışarıda kullanıldığında hata alırsınız çünkü kısa tanımlama yalnızca fonksiyonlar içinde kullanılabilir.

### 3. Çoklu Değişken Tanımlama

Go dilinde birden fazla değişkeni aynı anda tanımlayabilirsiniz. Bu, özellikle birden çok değişkenin bir fonksiyonda birlikte kullanılacağı durumlarda işe yarar.

```go
var a, b, c int
a = 10
b = 20
c = 30
```

Veya kısa tanımlama ile:

```go
x, y, z := 1, 2, "üç"
```

Bu özellik, dönüş değerleri olan fonksiyonlarla çalışırken oldukça kullanışlıdır:

```go
func hesapla() (int, int) {
    return 10, 20
}

a, b := hesapla()
fmt.Println(a, b) // Output: 10 20
```

### 4. Sıfır Değerler (Zero Values)
Go dilinde, bir değişkenin ilk değeri verilmediğinde, değişken türüne bağlı olarak sıfır değer atanır.

| Tür        | Sıfır Değer |
|------------|-------------|
| `int`      | `0`         |
| `float64`  | `0.0`       |
| `string`   | `""` (boş)  |
| `bool`     | `false`     |
| `pointer`  | `nil`       |

🎯 Örnek:

```go
var yas int
fmt.Println(yas) // Output: 0 (sıfır değer)
```

## 🧠 Sonuç ve İpuçları
- `var` ile tür belirterek değişkenler tanımlanabilir ve daha okunabilir bir kod elde edilebilir.
- Kısa tanımlama `:=` sadece fonksiyonlarda kullanılabilir, bu da kodu daha kompakt hale getirmeye yardımcı olur.
- Go’da sıfır değerler sayesinde, değişkenler her zaman başlatılacaktır.
- Değişken isimlerini anlamlı seçmek, kodun sürdürülebilirliğini artırır. 🌟
ß