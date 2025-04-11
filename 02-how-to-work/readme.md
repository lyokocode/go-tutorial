# 🚀 Go Nasıl Çalışır? Derleme, Binary Dosyalar ve Compiler Kavramı

[Bir önceki yazımda](../01-hello-world)  *Golang*’in ne olduğunu, neden tercih edilebileceğini, bir “Hello World” uygulamasının nasıl yazıldığını ve bu uygulamayı oluşturan yapıların ne işe yaradığını detaylı şekilde ele almıştık. Şimdi biraz daha derine inip “Go genel olarak nasıl çalışır?” sorusuna yanıt arayalım.

## 💡 Kod Yazmak Güzel Ama Arka Planı Bilmek de Faydalı

Kabul ediyorum, teorik bilgiler bazen sıkıcı olabilir. Hemen kod yazmak, denemek, görmek istiyoruz — ki yazılım öğrenmenin en keyifli ve etkili yolu da bu zaten.

Ancak bir dili daha iyi öğrenmek ve ileride karşılaşacağımız hataları daha kolay anlayabilmek için, en azından temel seviyede arka planda neler olup bittiğini bilmek gerçekten işe yarıyor.


Mesela Go için sıkça duyduğumuz bir tanım var:

>  🧠 “Go is a compiled language”

>  📌 Yani: “Go derlenen bir dildir.”
#### Peki nedir bu derleme? Ne işe yarar? Nasıl çalışır?

<hr>

### ✍️ Örnek Kodla Anlayalım

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello World!")
}
```

Bu kod parçasında görebildiğimiz bazı temel yapılar var:

- `main` isimli bir paket tanımlanmış,
- `fmt` paketi içe aktarılmış,
- ve bir `main` fonksiyonu yazılmış,
- sonuç olarak da terminale **“Hello World!”** yazdırılmış.

Bunları anlıyoruz, güzel. Ancak bu kod, bizim bilgisayarımız için hiçbir anlam ifade etmiyor çünkü bilgisayar sadece *binary*; **yani ikili** sayılarla çalışır. Bizim fonksiyonlarımız ve yazdığımız kodlar, bilgisayarın anlayabileceği formda değildir.

Yani bu kodları, bilgisayarın anlayabileceği dile çevirmemiz gerekir. İşte bu işleme `compile` (derleme) diyoruz.

### 🛠️ go build ile Derleme

Şu an projem `02-how-to-work` adlı bir dizinde. Bu dizine terminalden girip aşağıdaki komutu çalıştırdığımda:

```bash
go build main.go
```

Go, kodumu compile edip, bulunduğum dizine çalıştırılabilir bir dosya oluşturuyor: `main`

Bu, artık bilgisayarım tarafından doğrudan çalıştırılabilir bir `binary` dosya.

Terminalde bu dosyayı şöyle çalıştırabilirim:

```bash
./main
```

Ve sonuç:

```ngninx
Hello World!
```

🎉 Artık gördüğümüz gibi bu dosya bizim bilgisayarımız tarafından anlanıp çalışıyor.

Evet, çok büyük bir şey yapmadık belki ama bu küçük adım bile derleme sürecinin temelini anlamamız için yeterli.

### 🕐 Compile Time ve Runtime Nedir?

- 🛠️ Kodun yazılıp derlendiği sürece: Compile Time
- ▶️ Oluşan binary dosyasının çalıştığı sürece: Runtime diyoruz.

Yani:

`go build main.go` dediğimiz an → **Compile time**

`./main` dediğimiz ve çıktıyı aldığımız an → **Runtime**

### 🧩 Go’yu Neden İndiriyoruz?
Go'yu bilgisayarımıza kurarken **“Bir program mı indiriyoruz?”** diye düşünebiliriz. Aslında Go ile birlikte indirilen şey, bizim yazdığımız kodları derleyip binary dosyalara dönüştürebilecek bir `compiler` yani derleyicidir.

Yani şu komut:

```bash
go build main.go
```

Bu compiler sayesinde çalışır, compile edilie ve yazdığımız Go kodunu, bilgisayarın anlayabileceği dile çevirir.


### 🆚 Compiled ve Interpreted Diller

Go’nun `compiled` bir dil olduğunu artık biliyoruz. Kodumuzu derleyip bir *binary* oluşturuyoruz ve bu dosyayı alıp başka bir bilgisayarda da doğrudan çalıştırabiliyoruz — ekstra bir şey yapmamıza gerek kalmıyor.

##### Peki ya interpreted yani yorumlanan diller?

Örneğin JavaScript gibi dillerde böyle bir binary dosya yoktur. Kod çalıştığında satır satır yorumlanır. Başka bir bilgisayarda bu kodu çalıştırmak için hem kod dosyasına hem de çalıştıracak bir ortama (örneğin bir tarayıcıya ya da Node.js gibi bir yorumlayıcıya) ihtiyacımız olur.

##### Kısaca fark:

> 🛠️ `Compiled` → *Kod bir kere derlenir, ortaya çalışan bir dosya çıkar*.

> 📜 `Interpreted` → *Kod her çalıştırıldığında satır satır yorumlanır*.

⏱️ Genel olarak compiled diller, interpreted dillere göre daha hızlıdır. Ama elbette her yaklaşımın avantajları ve dezavantajları var — bunlara daha sonra değineceğiz.

## ✅ Sonuç
Go’yu öğrenmeye başlarken bu temel farkları ve çalışma mantığını bilmek, yazacağımız kodları daha bilinçli bir şekilde oluşturmamıza yardımcı olur. Derleme, runtime, compiler gibi kavramlar; birer teknik terim olmanın ötesinde, yazdığımız her satırın arkasında çalışan mekanizmaların parçalarıdır.
Şimdilik bu kadar. Bir sonraki yazıda görüşmek üzere 👋✨