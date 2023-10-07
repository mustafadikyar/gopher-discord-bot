# gopher-discord-bot

Bu proje, Go (Golang) programlama diliyle geliştirilmiş temel bir Discord botu şablonudur. 

Proje, yeni başlayanlar ve hızlı başlangıç yapmak isteyen geliştiriciler için tasarlanmıştır.

Temel Discord bot fonksiyonlarına sahiptir.


## Yükleme


```sh
git clone https://github.com/mustafadikyar/gopher-discord-bot.git
cd gopher-discord-bot
```

## Nasıl Kullanılır

#### Discord Botu Oluşturun:

Discord Developer Portal üzerinden yönergeleri takip ederek yeni bir bot oluşturun ve botunuzun token'ını alın.

#### Botu Sunucunuza Ekleyin:

Oluşturduğunuz Bot'un Client ID'sini kullanarak aşağıdaki URL'yi oluşturun ve tarayıcıda açın:
https://discord.com/oauth2/authorize?client_id=CLIENT_ID&scope=bot

Not: CLIENT_ID kısmını kendi botunuzun Client ID'siyle değiştirin.

#### Kanal ve Yetkileri Yönetin:

Botunuzun katılacağı kanalı ve yetkilerini bu link üzerinden düzenleyin.


## Kullanım

- `config.json` dosyasını düzenleyerek Discord botunuzun token'ını ve prefix'inizi ayarlayın.
- Botu başlatmak için terminalde `go run main.go` komutunu kullanabilir veya `go build` komutunu kullanarak bir .exe dosyası oluşturup, bu .exe dosyasını çalıştırarak botu başlatabilirsiniz.
- Projeyi genişletmek ve özelleştirmek için temel yapıyı kullanarak kendi komutlarınızı ekleyebilirsiniz.

  ENJOY!! :)
