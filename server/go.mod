module realmoriss/grumpy-free-jaguars/server

go 1.15

require (
	github.com/foolin/goview v0.3.0
	github.com/gin-contrib/sessions v0.0.3
	github.com/gin-gonic/gin v1.6.3
	github.com/gin-gonic/nosurf v0.0.0-20150415101651-45adcfcaf706
	github.com/gosimple/slug v1.9.0
	github.com/gwatts/gin-adapter v0.0.0-20170508204228-c44433c485ad
	golang.org/x/crypto v0.0.0-20201124201722-c8d3bf9c5392
	gorm.io/driver/sqlite v1.1.3
	gorm.io/gorm v1.20.7
	server v0.0.0-00010101000000-000000000000
)

replace server => ./
