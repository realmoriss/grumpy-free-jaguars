module realmoriss/grumpy-free-jaguars/server

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/gin-gonic/nosurf v0.0.0-20150415101651-45adcfcaf706
	github.com/gwatts/gin-adapter v0.0.0-20170508204228-c44433c485ad
	server v0.0.0-00010101000000-000000000000
)

replace server => ./
