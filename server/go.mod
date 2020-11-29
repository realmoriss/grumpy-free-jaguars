module realmoriss/grumpy-free-jaguars/server

go 1.15

require (
	github.com/foolin/goview v0.3.0
	github.com/gin-contrib/sessions v0.0.3
	github.com/gin-gonic/gin v1.6.3
	github.com/gin-gonic/nosurf v0.0.0-20150415101651-45adcfcaf706
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/gosimple/slug v1.9.0
	github.com/gwatts/gin-adapter v0.0.0-20170508204228-c44433c485ad
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/ugorji/go v1.2.0 // indirect
	golang.org/x/crypto v0.0.0-20201124201722-c8d3bf9c5392
	golang.org/x/sys v0.0.0-20201126233918-771906719818 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/driver/sqlite v1.1.3
	gorm.io/gorm v1.20.7
	server v0.0.0-00010101000000-000000000000
)

replace server => ./
