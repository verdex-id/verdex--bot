package model

type PDDIKTIUniversity struct {
	IDSP             string `json:"idsp,omitempty"`
	KabKotaPT        string `json:"kab_kota_pt,omitempty"`
	ProvinsiPT       string `json:"provinsi_pt,omitempty"`
	Akreditasi       string `json:"akreditasi,omitempty"`
	NamaSingkat      string `json:"nama_singkat,omitempty"`
	StatusPT         string `json:"status_pt,omitempty"`
	NamaPT           string `json:"nama_pt,omitempty"`
	JenisPT          string `json:"jenis_pt,omitempty"`
	JumlahProdi      int    `json:"jumlah_prodi,omitempty"`
	RangeBiayaKuliah string `json:"range_biaya_kuliah,omitempty"`
}

type PDDIKTIUniversityResponse struct {
	Data       []PDDIKTIUniversity `json:"data,omitempty"`
	Page       int                 `json:"page,omitempty"`
	Limit      int                 `json:"limit,omitempty"`
	TotalItems int                 `json:"totalItems,omitempty"`
	TotalPages int                 `json:"totalPages,omitempty"`
}
