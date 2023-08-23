package main

import (
	"encoding/json"
	"log"
)

type RekananAkta struct {
	LhkId            int64  `json:"lhk_id"`
	LhkNo            string `json:"lhk_no"`
	LhkTanggal       string `json:"lhk_tanggal"`
	LhkNotaris       string `json:"lhk_notaris"`
	LhkTipe          int    `json:"lhk_tipe"`
	RekanID          int32  `json:"rkn_id"`
	Deleted          bool   `json:"deleted"`
	SourceData       string `json:"source_data"`
	JenisAkta        string `json:"jenis_akta"`
	StatusVerifikasi string `json:"status_verifikasi"`
}

func main() {
	s := `[{"lhk_id":8480002793,"lhk_no":"14","lhk_tanggal":"Sep 24, 2018, 12:00:00 AM","lhk_notaris":"sumarsih, SH, Mkn","rkn_id":451999,"deleted":false,"source_data":"SIKaP LKPP","jenis_akta":"AKTA PENDIRIAN","lhk_tipe":0,"status_verifikasi":"unverified"},{"lhk_id":8480003093,"lhk_no":"15","lhk_tanggal":"Jul 26, 2023, 12:00:00 AM","lhk_notaris":"sumarsih, SH, Mkn","rkn_id":451999,"deleted":false,"source_data":"SIKaP LKPP","jenis_akta":"AKTA PERUBAHAN","lhk_tipe":1,"status_verifikasi":"unverified"}]`

	var m []RekananAkta

	err := json.Unmarshal([]byte(s), &m)
	if err != nil {
		log.Println(err)
	}

	log.Println(m)
}
