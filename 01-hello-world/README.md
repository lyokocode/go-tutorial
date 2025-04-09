## 📌 Golang Nedir?

- Go (veya Golang), Google tarafından geliştirilmiş, açık kaynaklı, statik tipli ve derlenen bir programlama dilidir.
- Basit sözdizimi, yüksek performansı ve eşzamanlı programlama (concurrency) konusundaki güçlü yapılarıyla bilinir.

## 💡 Golang nasıl ortaya çıktı?
Golang, Google mühendislerinin büyük ve karmaşık kod tabanlarını (codebase) daha verimli bir şekilde yönetme ihtiyacından doğmuştur. Google’ın devasa projelerinde yapılan her değişikliğin test edilmesi saatler sürebiliyordu. Bu durum geliştirme sürecini yavaşlatıyor ve verimliliği düşürüyordu.

Bu probleme çözüm olarak; daha sade, anlaşılır, hızlı derlenen ve yüksek performanslı bir programlama dili ihtiyacı ortaya çıktı. İşte bu ihtiyaçtan doğan Go, geliştirici deneyimini iyileştirmekle kalmadı, aynı zamanda sistem kaynaklarını da oldukça verimli kullanan bir yapıya sahip oldu.

## ✅ Golang’ın Öne Çıkan Avantajları:
- Son derece hızlı derleme süresi
- Sade ve okunabilir kod yapısı
- Gereksiz karmaşıklıklardan arındırılmış dil yapısı
- Gelişmiş eşzamanlılık desteği
- Platformlar arası taşınabilirlik

## 🚀 Go’nun Temel Özellikleri:
- Açık kaynaklı
- Statik olarak yazılmıştır
- Yüksek performans
- Hızlı derlenme (milyarlarca satırlık proje saniyeler içinde derlenebilir)
- Garbage collector (çöp toplayıcı) desteği
- Dahili paket yöneticisi
- Eşzamanlılık (dil seviyesinde)
- Her şeyin belgelendiği ayrıntılı dokümantasyonlar

# 💻 Kurulum
## 🖥️ Lokale Kurulum
Golang’ı kullanmaya başlamak için öncelikle sistemimize kurmamız gerekiyor. Kurulum dosyalarına [bu bağlantıdan](https://go.dev/doc/install) ulaşabilirsiniz.

Kurulumu tamamladıktan sonra terminal üzerinden aşağıdaki komutu çalıştırarak Go’nun doğru şekilde kurulup kurulmadığını kontrol edebilirsiniz:

```bash
go version
```
Komut başarılı çalışıyorsa Go versiyon bilgisi ekrana yazdırılır.

## 🧑‍💻 Hangi IDE?
Ben şu an için Visual Studio Code kullanıyorum. Hem alışkanlığım olduğu için hem de yeni başlayan biri olarak kullanım kolaylığı açısından benim için ideal bir tercih. İlerleyen zamanlarda JetBrains’in Go için geliştirdiği IDE’ye geçmeyi düşünebilirim.


## 🧩 Golang için Önerilen Eklentiler
İlk aşamada sadece golang uzantısını kurmak yeterli olacaktır. İlerledikçe farklı eklentiler ve paketlerle deneyerek öğrenmeye devam etmeyi planlıyorum.

# 👋 Go ile İlk Adım: "Hello World" Uygulaması

Yeni bir programlama diline başlarken klasik bir "Hello, World!" uygulaması yazmak neredeyse bir gelenek haline gelmiştir. Ne kadar basit görünse de, terminalde “Hello, World!” çıktısını görmek başlangıç için oldukça tatmin edici olacaktır.

Go dilinde bu klasik uygulama şu şekilde yazılır:

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World! I'm Aelita.")
}
```

### ▶️ Uygulamayı Nasıl Çalıştırırız?

Peki bu kodun satır satır ne anlama geldiğini biliyor musunuz? Gelin beraber inceleyelim ve öğrenelim.

Go dilinde bir .go dosyasını çalıştırmak oldukça kolaydır. Aşağıdaki adımları takip ederek yazdığınız programı terminal (komut satırı) üzerinden çalıştırabilirsiniz.

#### 1. Dosyayı Kaydedin
Örneğin kodu bir dosyaya kaydedelim:

```bash
hello-world.go
```

#### 2. Terminali Açın ve Dosyanın Bulunduğu Klasöre Geçin
```
cd /path/to/your/file
```

örneğin;

```bash
cd go-tutorial/01-hello-world
```

##### 3. Programı Çalıştırın
Go dosyasını çalıştırmak için aşağıdaki komutu kullanabilirsin:

```go
go run hello-world.go
```
Eğer her şey yolundaysa, terminalde şu çıktıyı görmelisin: `Hello World! I'm Aelita`

### 📦 package main Nedir?

Go dilinde her dosya bir package (paket) içinde yer alır. Bu, Go’nun modüler ve okunabilir bir yapıya sahip olmasını sağlar.
package main ise Go programlarının başlangıç noktasıdır. Yani çalıştırılabilir bir uygulama oluşturmak istiyorsanız, programınızı main paketi içinde yazmalısınız.
🔍 Unutmayın:
- package main sadece çalıştırılabilir (executable) programlar için kullanılır.
- Eğer bir kütüphane (library) geliştiriyorsanız, paket adını farklı tanımlayabilirsiniz (örneğin: package utils, package math gibi).

### 📥 import "fmt" Nedir?
Go'da başka paketleri kullanmak için import anahtar kelimesi kullanılır.
Bu örnekte fmt paketini içe aktarıyoruz.

#### 📘 fmt Paketi Ne İşe Yarar?
fmt, Go’nun standart kütüphanelerinden biridir ve formatlı I/O (giriş/çıkış) işlemlerini yapmanıza olanak tanır.
fmt, Go’nun standart kütüphanelerinden biridir ve formatlı I/O (giriş/çıkış) işlemlerini yapmanıza olanak tanır.

##### Neden Parantez içinde yazılıyor?
```go
import (
    "fmt"
)
```

Birden fazla paket import edilecekse bu şekilde parantezli yazım tercih edilir. Tek satırlık import işlemleri şöyle de yazılabilir:

```go 
import "fmt"
```

Ama genelde proje büyüdükçe daha fazla paket import edileceği için parantezli yapı standarttır.

### 🔧 func main() Nedir?

Go programlarında kodun çalışmaya başladığı nokta main fonksiyonudur.
Bu fonksiyon, package main içinde tanımlanmalı ve şu şekilde yazılmalıdır:

```go
func main() {
    // Kodlar buraya
}
```

Peki başka isimde fonksiyon tanımlanabilir mi?
Evet! Go'da birçok fonksiyon tanımlayabilirsiniz. Örneğin:

```go
func hello() {
	fmt.Println("Merhaba!")
}
```

Ancak bir programın başlatılabilmesi için mutlaka main() fonksiyonunun tanımlanmış olması gerekir. Go derleyicisi programı çalıştırmak için bu fonksiyonu arar.

### 🧱 fmt.Println(...) Ne Yapar?
fmt.Println() fonksiyonu, parametre olarak aldığı veriyi ekrana yazdırır ve ardından bir alt satıra geçer.

```go
fmt.Println("Hello World! I'm Aelita.")
```
Yukarıdaki satır terminal ekranına şu çıktıyı verir:

```bash
Hello World! I'm Aelita.
```
📝 Eğer alt satıra geçmeden yazmak isterseniz fmt.Print() fonksiyonunu kullanabilirsiniz.

### 🧪 Ekstra: Kendi Fonksiyonunu Yaz!

```go
package main

import (
	"fmt"
)

func main() {
	hello()
}

func hello() {
	fmt.Println("Welcome to Golang's world!!!")
}
```

🎯 Sonuç
Golang, sadeliği ve performansı ile yazılıma yeni başlayanlardan profesyonellere kadar pek çok geliştiriciye hitap eden güçlü bir dildir.
"Hello World" gibi basit bir örnek bile, dilin yapısını ve mantığını anlamamız için harika bir başlangıçtır.
Bir sonraki adımda, fonksiyonlara parametre geçirme, veri tipleri, değişken tanımlamaları ve kontrol yapıları gibi konulara geçiş yapabiliriz.
