# Kodun İncelenmesi: 🧐

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
}
```

Daha önceki yazılarımızda bu fonksiyonun ne yaptığını görmüştük. Ancak, şimdi daha detaylı bir şekilde inceleyelim. Bu fonksiyon nasıl çıktı verir, paketleri nasıl import ederiz, nasıl compile ederiz gibi bilgileri öğrendik. Fakat `package main ` ne anlama gelir? Neden bunu tanımlıyoruz? Ne işe yarar? 🧐

Haydi, gelin tüm bu kodları adım adım analiz edelim!


## 1. Bölüm: Paket Tanımlaması 📦

```go
// Package Clause 
package main
```

Go programlama dilinde yazdığınız tüm kodlar paketler halinde düzenlenir. Yani, .go uzantılı yazdığınız her kod bir paketin içinde yer alır. Bu, Go'nun modüler yapısını sağlamak için oldukça kullanışlıdır. 🛠️

Projeniz büyüdükçe, kodlarınızı modüllere ayırarak her birine belirli bir paket atayabilirsiniz.

Bunun için en güzel örnek:
Evimize taşındık diyelim ve taşındığımız evde bir sürü paket halinde koliler var. Eğer bu koliler dağınık bir şekilde ortada durursa, içlerinde ne olduğunu bilemeyeceğiz. Evi yerleştirirken (yani kodu yazarken) büyük bir kaos yaşarız. 😵

Fakat, taşınma sürecinde (yani kod yazarken) her koliyi gruplara ayırırsak (örneğin, teknolojik aletler bir kolide, bardaklar bir kolide vb.), hem o an bir şey ararken hem de evi yerleştirirken çok daha rahat ederiz. 🏠📦

Go'da paketleme, kodunuzu düzenlemenin ve yönetmenin tam olarak böyle bir yolu. Modüler bir yapıyla her şey yerli yerine oturur ve projeniz büyüdükçe kodunuzun karmaşıklaşmasını engellersiniz. 🌱

Buradaki kodda, main adında bir paket tanımlıyoruz.

#### Peki, main paketi neden seçiliyor? 🤔

`Go`'da, programınız çalıştığında, main paketine sahip bir fonksiyon çalıştırılır. Eğer paket adını `section` gibi farklı bir şey yaparsak, `go run main.go` komutunu çalıştırdığımızda **package command-line-arguments is not a main package** hatası alırız.

`main` paketi, Go programının *başlangıç noktasıdır*. Yani, **çalıştırılabilir** bir program yazmak istiyorsak, paketin adı main olmak zorundadır. Bu özel paket sayesinde, Go programınızı başlatabiliriz.


## 2. Bölüm: Import 📥

```go
// Import Statement
import (
	"fmt"
)
```

Go'da, yazdığımız kodu sadece kendi paketlerimizle sınırlamak zorunda değiliz. Başka paketleri de kullanabiliriz! 😎

Öncelikle, Go'nun standart paketleri bize birçok hazır işlev sağlar. Bunlar, Go'nun sunduğu fonksiyonlar, veri yapıları ve araçlardır. Bunun dışında, *diğer geliştiriciler tarafından yazılmış paketler*i de kullanabiliriz. Eğer her zaman kendi paketlerimizi yazmaya kalksaydık, bu iş yükünü kaldırmak gerçekten çok zor olurdu.
Örneğin, Go'nun `math` paketini kullanarak matematiksel işlemleri çok kolay yapabilirken, eğer kendimiz bir matematiksel paket yazmaya kalksak, basit bir işlemi yapmak için saatlerce uğraşmamız gerekebilirdi. ⏳🔧
`import` komutuyla, dışarıdan paketleri programa dahil ederiz. Bu sayede, Go'nun sunduğu güçlü ve yaygın kullanılan fonksiyonları, kendi projemizde kolayca kullanabiliriz.
Go dilinde genellikle `fmt` gibi yaygın paketler çok sık kullanılır. `fmt` paketi, ekrana yazdırma, kullanıcıdan veri alma gibi temel işlemleri gerçekleştirmek için kullanılır.

### Neden import Kullanıyoruz?

Düşünün ki, her iş için sıfırdan bir şeyler yazmak zorunda kalıyorsunuz. Her seferinde aynı şeyi yaparak zaman kaybedersiniz. Ama Go'nun sağladığı paketlerle, yazılım geliştirme sürecini çok daha hızlı ve verimli hale getirebilirsiniz. 🏃‍♂️💨

Go'nun paket yapısı ve `import` kullanımı, kodunuzu daha düzenli, modüler ve sürdürülebilir hale getirir. 👨‍💻🔧

## 3. Bölüm: main Fonksiyonu ve Yapısı 🧩

```go
func main() {
    fmt.Println("Hello World!")
}
```

Go dilinde yazdığınız her programda genellikle bir `main` fonksiyonu bulunur. Bu fonksiyon, Go programının başlangıç noktasıdır. Ancak, main fonksiyonunun yapısını ve her bir parçasını detaylı olarak incelemeden, tam olarak ne yaptığını anlayamayız. Hadi, gelin bunu adım adım inceleyelim! 👇

### 1. func Anahtar Kelimesi 🗝️

```go
func
```

Go dilinde bir fonksiyon tanımlamak için `func` anahtar kelimesini kullanırız. Bu, fonksiyonun başladığını gösterir. Fonksiyonlar, belirli bir işlevi yerine getiren ve bir değer döndürebilen (veya döndürmeyen) kod bloklarıdır. Bir fonksiyon tanımlarken, ona ad veririz, sonra parametreler (gerekiyorsa) ekleriz ve fonksiyonun çalıştıracağı kodu süslü parantezler {} içinde belirtiriz.

### 2. main Fonksiyonun Adı 🏷️

```go
main
```

Go programlarının başlangıç noktası olan fonksiyon `main` olarak adlandırılır. Program çalıştırıldığında, main fonksiyonu otomatik olarak çağrılır ve burada yazdığınız kod çalıştırılır. main fonksiyonunun adı özel bir anlam taşır; Go programının çalıştırılabilir bir program olması için bu fonksiyonun adı main olmak zorundadır.

Eğer başka bir isimle fonksiyon yazarsanız, Go çalıştırırken hata verir. Çünkü main fonksiyonu, Go programında programın başlatılmasını sağlayan yerdir.

### 3. Parantezler () ve Argümanlar 🎯

```go
()
```

Fonksiyon adından sonra parantezler `()`, fonksiyonun argümanlarını belirtmek için kullanılır. Fonksiyonlar bazen dışarıdan veri alması gereken bir işlem gerçekleştirebilir. Bu verilere "parametre" veya "argüman" denir. Ancak burada main fonksiyonu hiçbir parametre almaz. Yine de, Go'da bir fonksiyonun parametre alıp almadığını belirlemek için bu parantezler her zaman gereklidir.

Eğer bir fonksiyon parametre alacaksa, parametreler burada tanımlanır. Örneğin:

```go
func greet(name string) {
    fmt.Println("Hello", name)
}
```

Burada, `name` fonksiyona gönderilen bir *argümandır* ve fonksiyon, name parametresini ekrana yazdırır.

### 4. Süslü Parantezler {} ve Fonksiyonun Gövdesi 🏗️

```go
{
    fmt.Println("Hello World!")
}
```

Fonksiyonlar, işlevsel olarak bir işlem yapacakları kodu süslü parantezler `{}` içinde belirtirler. Bu süslü parantezler fonksiyonun başladığını ve bittiğini gösterir. Fonksiyonun içindeki her şey, main fonksiyonunun gövdesi olarak kabul edilir.
Yani, süslü parantezler main fonksiyonunun içerisine yazılacak işlemleri kapsar ve bu fonksiyon çağrıldığında içindeki kodlar çalıştırılır.

### 5. fmt.Println("Hello World!") ve Fonksiyonun İçeriği 🖨️
Bu satır, `fmt` paketini kullanarak ekrana **"Hello World!"** yazdırır. fmt paketi, Go’nun en yaygın kullanılan standart paketlerinden biridir ve ekrana veri yazdırmak gibi işlemleri çok kolay hale getirir.

- **fmt**: Go'daki formatlı giriş ve çıkış işlemleri için kullanılan paket.
**Println**: fmt paketinin bir fonksiyonudur ve argüman olarak aldığı veriyi ekrana yazdırır. Ayrıca, yazdırılan veriden sonra bir yeni satır karakteri ekler.
**"Hello World!"**: Bu, ekrana yazdırılacak metindir.

Bu satırın işlevi, main fonksiyonu çalıştırıldığında ekranda "Hello World!" mesajını göstermek olacaktır.

### 6. Fonksiyonun Çağrılması ve Çalışması 📞
Fonksiyon `main`, Go programı çalıştırıldığında otomatik olarak çağrılır. Yani, Go programının başlatılmasıyla birlikte `main` fonksiyonu da çalışmaya başlar ve içinde yazılan tüm işlemler sırasıyla gerçekleştirilir. Buradaki **fmt.Println("Hello World!")** satırı, fonksiyon çağrıldığında çalışacak ve sonucu ekrana yazdıracaktır.

## Özetle:
- `func` anahtar kelimesi, bir fonksiyonun başladığını belirtir.
- `main` fonksiyonun adı, Go programlarında başlangıç fonksiyonunu temsil eder.
`()` parantezler, fonksiyonun **parametrelerini** belirtir. Eğer fonksiyon parametre almazsa boş kalır.
- Süslü parantezler `{}`, fonksiyonun içeriğini, yani ne yapılacağını belirler.
- `fmt.Println("Hello World!")`, bu fonksiyonun içeriğinde ekrana yazdırma işlemi gerçekleştirir.

Bu yapının tamamı, Go programının başlangıç fonksiyonunun nasıl çalıştığını ve main fonksiyonunun nasıl yapılandığını detaylı bir şekilde gösterir. Gövde içerisinde yazılı her işlem, fonksiyon çağrıldığında çalışacaktır! 🚀

