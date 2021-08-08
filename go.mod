module golang-example
go 1.16

require (
	"golang-example/addition" v0.0.0
)
replace "golang-example/addition" v0.0.0 => "./addition"
