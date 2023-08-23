package main

import (
	"encoding/base64"
	"log"
	"net/url"

	"github.com/arymaulanamalik/go_sample/blowfish"
	"github.com/arymaulanamalik/go_sample/uuid"
)

func main() {

	// pt := "1204163"
	// pt := "4645034"
	pt := "18332009" //rekanan testing 2
	// pt := "1583198"
	// pt := "9547042"

	key := "mysecretkey"

	a, _ := blowfish.EncryptBlowfish([]byte(pt), []byte(key))
	rknId := base64.StdEncoding.EncodeToString(a)
	log.Println(rknId)

	bEscape := url.QueryEscape(rknId)
	bEscape = "https://sikap.lkpp.go.id/services/getRekananPajak?q=" + bEscape
	log.Println(bEscape)

	b := uuid.NewUUIDString()
	log.Println(b)

	test := url.QueryEscape("https://databunker.eproc.dev/v1/agreement/term-and-condition/userId/39fbb734-85a3-4165-945e-f157d6500299")
	log.Println(test)
}
