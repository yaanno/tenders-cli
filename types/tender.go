package types

import "time"

// https://tenders.guru/api/hu/tenders
type PaginatedTenderList struct {
	PageCount  int      `json:"page_count"`
	PageNumber int      `json:"page_number"`
	PageSize   int      `json:"page_size"`
	Total      int      `json:"total"`
	Data       []Tender `json:"data"`
}

// https://tenders.guru/api/hu/tenders/35337
type Tender struct {
	ID                 string `json:"id"`
	Date               string `json:"date"`
	DeadlineDate       string `json:"deadline_date"`
	DeadlineLengthDays string `json:"deadline_length_days"`
	Title              string `json:"title"`
	Category           string `json:"category"`
	Description        string `json:"description"`
	Phase              string `json:"phase"`
	PhaseEn            string `json:"phase_en"`
	Place              string `json:"place"`
	Sid                string `json:"sid"`
	Eid                string `json:"eid"`
	AwardedValue       string `json:"awarded_value"`
	AwardedCurrency    string `json:"awarded_currency"`
	AwardedValueEur    string `json:"awarded_value_eur"`
	Data               struct {
		Final    string `json:"final,omitempty"`
		Currency string `json:"currency,omitempty"`
	} `json:"data,omitempty"`
	Purchaser struct {
		ID   int    `json:"id"`
		Sid  int64  `json:"sid"`
		Name string `json:"name"`
	} `json:"purchaser"`
	Type struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"type"`
	Notices struct {
		OSSZEGZESAJANLATELBIRALASBLOKK []struct {
			ID        string `json:"id"`
			Sid       string `json:"sid"`
			Name      string `json:"name"`
			Type      string `json:"type"`
			Date      string `json:"date"`
			Comments  string `json:"comments"`
			Block     string `json:"block"`
			DocID     string `json:"doc_id"`
			DocURL    string `json:"doc_url"`
			DocSize   any    `json:"doc_size"`
			DocPages  any    `json:"doc_pages"`
			DocAccept string `json:"doc_accept"`
		} `json:"OSSZEGZES_AJANLAT_ELBIRALAS_BLOKK"`
		AJANLATTETELREFELHIVOTTAKBLOKK []struct {
			ID        string `json:"id"`
			Sid       string `json:"sid"`
			Name      string `json:"name"`
			Type      string `json:"type"`
			Date      string `json:"date"`
			Comments  string `json:"comments"`
			Block     string `json:"block"`
			DocID     string `json:"doc_id"`
			DocURL    string `json:"doc_url"`
			DocSize   any    `json:"doc_size"`
			DocPages  any    `json:"doc_pages"`
			DocAccept string `json:"doc_accept"`
		} `json:"AJANLATTETELRE_FELHIVOTTAK_BLOKK"`
		BONTASIJEGYZOKONYVBLOKK []struct {
			ID        string `json:"id"`
			Sid       string `json:"sid"`
			Name      string `json:"name"`
			Type      string `json:"type"`
			Date      string `json:"date"`
			Comments  string `json:"comments"`
			Block     string `json:"block"`
			DocID     string `json:"doc_id"`
			DocURL    string `json:"doc_url"`
			DocSize   any    `json:"doc_size"`
			DocPages  any    `json:"doc_pages"`
			DocAccept string `json:"doc_accept"`
		} `json:"BONTASI_JEGYZOKONYV_BLOKK"`
		EGYEBDOKUMENTUMBLOKK []struct {
			ID        string `json:"id"`
			Sid       string `json:"sid"`
			Name      string `json:"name"`
			Type      string `json:"type"`
			Date      string `json:"date"`
			Comments  any    `json:"comments"`
			Block     string `json:"block"`
			DocID     string `json:"doc_id"`
			DocURL    string `json:"doc_url"`
			DocSize   any    `json:"doc_size"`
			DocPages  any    `json:"doc_pages"`
			DocAccept string `json:"doc_accept"`
		} `json:"EGYEB_DOKUMENTUM_BLOKK"`
		KOZBESZERZESIDOKUMENTACIOBLOKK []struct {
			ID        string `json:"id"`
			Sid       string `json:"sid"`
			Name      string `json:"name"`
			Type      string `json:"type"`
			Date      string `json:"date"`
			Comments  string `json:"comments"`
			Block     string `json:"block"`
			DocID     string `json:"doc_id"`
			DocURL    string `json:"doc_url"`
			DocSize   any    `json:"doc_size"`
			DocPages  any    `json:"doc_pages"`
			DocAccept string `json:"doc_accept"`
		} `json:"KOZBESZERZESI_DOKUMENTACIO_BLOKK"`
	} `json:"notices"`
	Awarded []struct {
		Date          string `json:"date"`
		SuppliersID   any    `json:"suppliers_id"`
		Value         int    `json:"value"`
		Count         int    `json:"count"`
		SuppliersName string `json:"suppliers_name"`
		Suppliers     []struct {
			ID   int    `json:"id"`
			Slug string `json:"slug"`
			Name string `json:"name"`
		} `json:"suppliers"`
		ValueEur         float32 `json:"value_eur"`
		ValueForOne      float32 `json:"value_for_one"`
		ValueForOneEur   float32 `json:"value_for_one_eur"`
		ValueForTwo      float32 `json:"value_for_two"`
		ValueForTwoEur   float64 `json:"value_for_two_eur"`
		ValueForThree    float32 `json:"value_for_three"`
		ValueForThreeEur float64 `json:"value_for_three_eur"`
		OffersCountData  struct {
			Num2 int `json:"2"`
		} `json:"offers_count_data"`
		OffersCount []int `json:"offers_count"`
	} `json:"awarded"`
}

// https://tenders.guru/api/hu/tenders/35337/source_data
type TenderDataSource struct {
	VezetoAjanlatkeroSzervezet struct {
		ID               int64  `json:"id"`
		RovidMegnevezes  string `json:"rovidMegnevezes"`
		TeljesMegnevezes string `json:"teljesMegnevezes"`
		Szekhely         struct {
			Orszag        string `json:"orszag"`
			Iranyitoszam  string `json:"iranyitoszam"`
			Telepules     string `json:"telepules"`
			KozterNeve    string `json:"kozterNeve"`
			KozterJellege string `json:"kozterJellege"`
			Hazszam       string `json:"hazszam"`
			Helyrajziszam string `json:"helyrajziszam"`
			NutsKod       string `json:"nutsKod"`
		} `json:"szekhely"`
		AjanlatkeroJogalapok20200722Utan string `json:"ajanlatkeroJogalapok20200722Utan"`
		KozszolgaltatoJogalap            any    `json:"kozszolgaltatoJogalap"`
		FotevekenysegiKod                any    `json:"fotevekenysegiKod"`
	} `json:"vezetoAjanlatkeroSzervezet"`
	KozbeszerzesiEljaras struct {
		ID                         int64  `json:"id"`
		TechnikaiAzonosito         string `json:"technikaiAzonosito"`
		VezetoAjanlatkeroSzervezet struct {
			TeljesMegnevezes                    string `json:"teljesMegnevezes"`
			Adoszam                             string `json:"adoszam"`
			HeaEuropaiUniosAdoszam              any    `json:"heaEuropaiUniosAdoszam"`
			KulsoNyilvantartasMegnevezes        any    `json:"kulsoNyilvantartasMegnevezes"`
			KulsoNyilvantartasSzerintiAzonosito any    `json:"kulsoNyilvantartasSzerintiAzonosito"`
			AjanlatkeroJogalapok20200722Utan    any    `json:"ajanlatkeroJogalapok20200722Utan"`
			SzervezetTipusok                    any    `json:"szervezetTipusok"`
			FotevekenysegiEgyeb                 any    `json:"fotevekenysegiEgyeb"`
			Statusz                             any    `json:"statusz"`
			PostaiCimEgyezikASzekhellyel        any    `json:"postaiCimEgyezikASzekhellyel"`
			PostaiCimPostafiok                  any    `json:"postaiCimPostafiok"`
			Kozszolgaltato                      any    `json:"kozszolgaltato"`
			KlasszikusAjanlatkero               any    `json:"klasszikusAjanlatkero"`
			KlasszikusAjanlatkeroTipus          any    `json:"klasszikusAjanlatkeroTipus"`
			AjanlatkeroTipus                    any    `json:"ajanlatkeroTipus"`
			EgyebAjanlatkeroTipus               any    `json:"egyebAjanlatkeroTipus"`
			SzervezetTevekenysegek              []any  `json:"szervezetTevekenysegek"`
			EgyebSzervezetTevekenyseg           any    `json:"egyebSzervezetTevekenyseg"`
			KozszolgaltatoJogalap               any    `json:"kozszolgaltatoJogalap"`
			KozszolgaltatoTevekenysegiKorok     []any  `json:"kozszolgaltatoTevekenysegiKorok"`
			EljarasEllenorzesFeladatok          []any  `json:"eljarasEllenorzesFeladatok"`
			NyilvantartasVezetoFeladatok        []any  `json:"nyilvantartasVezetoFeladatok"`
			EgyebErdekeltSzervezetFeladatok     []any  `json:"egyebErdekeltSzervezetFeladatok"`
			EgyebErdekeltSzervezetEgyebFeladata any    `json:"egyebErdekeltSzervezetEgyebFeladata"`
			VallalkozasTipus                    any    `json:"vallalkozasTipus"`
			CegSzekhelyTipus                    any    `json:"cegSzekhelyTipus"`
			Szekhely                            struct {
				Orszag        string `json:"orszag"`
				Iranyitoszam  string `json:"iranyitoszam"`
				Telepules     string `json:"telepules"`
				KozterNeve    string `json:"kozterNeve"`
				KozterJellege string `json:"kozterJellege"`
				Hazszam       string `json:"hazszam"`
				Helyrajziszam string `json:"helyrajziszam"`
				NutsKod       string `json:"nutsKod"`
			} `json:"szekhely"`
			LevelezesiCim        any    `json:"levelezesiCim"`
			Kapcsolat            any    `json:"kapcsolat"`
			TechnikaiAzonosito   string `json:"technikaiAzonosito"`
			ID                   int64  `json:"id"`
			Szerkesztheto        any    `json:"szerkesztheto"`
			JogalanyTipus        string `json:"jogalanyTipus"`
			ErvenyessegVege      any    `json:"ervenyessegVege"`
			RovidMegnevezes      string `json:"rovidMegnevezes"`
			TulajdonosIlletosege any    `json:"tulajdonosIlletosege"`
			AdoazonositoJel      any    `json:"adoazonositoJel"`
		} `json:"vezetoAjanlatkeroSzervezet"`
		SzerzodesTipusa                         string `json:"szerzodesTipusa"`
		KozbeszerzesiEljarasTargya              string `json:"kozbeszerzesiEljarasTargya"`
		IranyadoEljarasiRend                    string `json:"iranyadoEljarasiRend"`
		KozbeszerzesiEljarasTipusa              string `json:"kozbeszerzesiEljarasTipusa"`
		EljarasAktualisSzakasza                 string `json:"eljarasAktualisSzakasza"`
		JogosultsagKarbantartoJoggalRendelkezik any    `json:"jogosultsagKarbantartoJoggalRendelkezik"`
		EljarasMegnyitoJoggalRendelkezik        any    `json:"eljarasMegnyitoJoggalRendelkezik"`
		TeljesitesHelye                         string `json:"teljesitesHelye"`
		LegkozelebbiHatarido                    string `json:"legkozelebbiHatarido"`
		InditoKm1EljarasTargya                  any    `json:"inditoKm1EljarasTargya"`
		InditoKm1EljarastechnikaiAzonosito      any    `json:"inditoKm1EljarastechnikaiAzonosito"`
		EljarasMeginditasanakDatuma             string `json:"eljarasMeginditasanakDatuma"`
		AjanlatteteliHatarido                   string `json:"ajanlatteteliHatarido"`
		EljarasHataridok                        []any  `json:"eljarasHataridok"`
		Hirdetmenyek                            []struct {
			HirdetmenyIktatoszam string `json:"hirdetmenyIktatoszam"`
			TedAzonosito         any    `json:"tedAzonosito"`
			HirdetmenyTipusa     string `json:"hirdetmenyTipusa"`
		} `json:"hirdetmenyek"`
	} `json:"kozbeszerzesiEljaras"`
	KapcsolodoDokumentumok []struct {
		ID                     int64  `json:"id"`
		DokumentumNeve         string `json:"dokumentumNeve"`
		DokumentumTipusa       string `json:"dokumentumTipusa"`
		PublikalasDatuma       string `json:"publikalasDatuma"`
		Megjegyzes             string `json:"megjegyzes"`
		EljarasDokumentumBlokk string `json:"eljarasDokumentumBlokk"`
	} `json:"kapcsolodoDokumentumok"`
	ErdeklodesJelzeseEljarasra                       bool `json:"erdeklodesJelzeseEljarasra"`
	ErdeklodesJelzeseOsszefoglaloTajekoztatoAlapajan bool `json:"erdeklodesJelzeseOsszefoglaloTajekoztatoAlapajan"`
	ElozetesenErdeklodesJelzeseEljarasra             bool `json:"elozetesenErdeklodesJelzeseEljarasra"`
	Hirdetmenyek                                     []struct {
		ID                                   int64  `json:"id"`
		HirdetmenyEKRazonosito               string `json:"hirdetmenyEKRazonosito"`
		AjanlatkeroNeve                      string `json:"ajanlatkeroNeve"`
		EljarasTargya                        string `json:"eljarasTargya"`
		HirdetmenyIktatoszam                 string `json:"hirdetmenyIktatoszam"`
		TedAzonosito                         any    `json:"tedAzonosito"`
		HirdetmenyTipusa                     string `json:"hirdetmenyTipusa"`
		EljarasAjanlatteteliHatarido         string `json:"eljarasAjanlatteteliHatarido"`
		HirdetmenyKozzetetelDatuma           string `json:"hirdetmenyKozzetetelDatuma"`
		EljarasTechnikaiAzonosito            string `json:"eljarasTechnikaiAzonosito"`
		EljarasBeszerzesTargya               string `json:"eljarasBeszerzesTargya"`
		EljarasCPVkod                        []any  `json:"eljarasCPVkod"`
		EljarasNyertesAjanlattevoNeve        string `json:"eljarasNyertesAjanlattevoNeve"`
		EljarasTeljesitesHelye               string `json:"eljarasTeljesitesHelye"`
		EljarasEljarasRend                   string `json:"eljarasEljarasRend"`
		EljarasEljarasTipusa                 string `json:"eljarasEljarasTipusa"`
		AjanlatkeroTipus                     string `json:"ajanlatkeroTipus"`
		AjanlatkeroKozszolgTevekenysegiKor   any    `json:"ajanlatkeroKozszolgTevekenysegiKor"`
		AjanlatkeroKlasszikusTevekenysegiKor string `json:"ajanlatkeroKlasszikusTevekenysegiKor"`
		HirdetmenyErtesitoLapszam            string `json:"hirdetmenyErtesitoLapszam"`
		DokumentumAzonosito                  string `json:"dokumentumAzonosito"`
		HirdetmenyTedURL                     any    `json:"hirdetmenyTedUrl"`
		TedLapszam                           any    `json:"tedLapszam"`
		TedKozzetetelDatuma                  any    `json:"tedKozzetetelDatuma"`
	} `json:"hirdetmenyek"`
}

// https://tenders.guru/api/hu/tenders/35337/notices
type TenderNotice struct {
	Data []struct {
		ID                         int       `json:"id"`
		Sid                        string    `json:"sid"`
		Date                       time.Time `json:"date"`
		Title                      string    `json:"title"`
		TypeID                     int       `json:"type_id"`
		NoticeTypeID               int       `json:"notice_type_id"`
		Regime                     string    `json:"regime"`
		Category                   string    `json:"category"`
		ContractingAuthorityID     int       `json:"contracting_authority_id"`
		ContractingAuthorityTypeID int       `json:"contracting_authority_type_id"`
		CpvCount                   int       `json:"cpv_count"`
		NotEmptySectionsCount      int       `json:"not_empty_sections_count"`
		ProcurementID              int       `json:"procurement_id"`
		ProcurementSid             string    `json:"procurement_sid"`
	} `json:"data"`
}

// https://tenders.guru/api/hu/tenders/35337/docs
type TenderDocument struct {
	Data []struct {
		ID            int       `json:"id"`
		Sid           string    `json:"sid"`
		ProcurementID int       `json:"procurement_id"`
		Name          string    `json:"name"`
		Type          string    `json:"type"`
		Date          time.Time `json:"date"`
		Comments      string    `json:"comments"`
		Block         string    `json:"block"`
		DocID         int       `json:"doc_id"`
	} `json:"data"`
}

// https://tenders.guru/api/hu/notices
type PaginatedNoticeList struct {
	CurrentPage int `json:"current_page"`
	Data        []struct {
		ID                         int    `json:"id"`
		Sid                        string `json:"sid"`
		Date                       any    `json:"date"`
		Title                      any    `json:"title"`
		TypeID                     any    `json:"type_id"`
		NoticeTypeID               any    `json:"notice_type_id"`
		Regime                     any    `json:"regime"`
		Category                   any    `json:"category"`
		ContractingAuthorityID     any    `json:"contracting_authority_id"`
		ContractingAuthorityTypeID any    `json:"contracting_authority_type_id"`
		CpvCount                   any    `json:"cpv_count"`
		NotEmptySectionsCount      any    `json:"not_empty_sections_count"`
		Status                     string `json:"status"`
		StatusTs                   string `json:"status_ts"`
		Pdf                        string `json:"pdf"`
		PdfTs                      string `json:"pdf_ts"`
		PdfSize                    any    `json:"pdf_size"`
		PdfHash                    any    `json:"pdf_hash"`
		XML                        string `json:"xml"`
		XMLTs                      string `json:"xml_ts"`
		Analysis                   string `json:"analysis"`
		AnalysisTs                 string `json:"analysis_ts"`
		ProcurementID              any    `json:"procurement_id"`
		ProcurementSid             string `json:"procurement_sid"`
		Accept                     int    `json:"accept"`
		HasPhrase                  int    `json:"has_phrase"`
		ReceiverID                 any    `json:"receiver_id"`
	} `json:"data"`
	FirstPageURL string `json:"first_page_url"`
	From         int    `json:"from"`
	LastPage     int    `json:"last_page"`
	LastPageURL  string `json:"last_page_url"`
	Links        []struct {
		URL    any    `json:"url"`
		Label  string `json:"label"`
		Active bool   `json:"active"`
	} `json:"links"`
	NextPageURL string `json:"next_page_url"`
	Path        string `json:"path"`
	PerPage     int    `json:"per_page"`
	PrevPageURL any    `json:"prev_page_url"`
	To          int    `json:"to"`
	Total       int    `json:"total"`
}

// https://tenders.guru/api/hu/notices/87012568
type Notice struct {
	Data struct {
		ID                         int    `json:"id"`
		Sid                        string `json:"sid"`
		Date                       any    `json:"date"`
		Title                      any    `json:"title"`
		TypeID                     any    `json:"type_id"`
		NoticeTypeID               any    `json:"notice_type_id"`
		Regime                     any    `json:"regime"`
		Category                   any    `json:"category"`
		ContractingAuthorityID     any    `json:"contracting_authority_id"`
		ContractingAuthorityTypeID any    `json:"contracting_authority_type_id"`
		CpvCount                   any    `json:"cpv_count"`
		NotEmptySectionsCount      any    `json:"not_empty_sections_count"`
		ProcurementID              any    `json:"procurement_id"`
		ProcurementSid             string `json:"procurement_sid"`
	} `json:"data"`
}

// https://tenders.guru/api/hu/notices/87012568/source_data
type NoticeDataSource struct {
	ID                                   int64  `json:"id"`
	HirdetmenyEKRazonosito               string `json:"hirdetmenyEKRazonosito"`
	AjanlatkeroNeve                      string `json:"ajanlatkeroNeve"`
	EljarasTargya                        string `json:"eljarasTargya"`
	HirdetmenyIktatoszam                 string `json:"hirdetmenyIktatoszam"`
	TedAzonosito                         any    `json:"tedAzonosito"`
	HirdetmenyTipusa                     string `json:"hirdetmenyTipusa"`
	EljarasAjanlatteteliHatarido         string `json:"eljarasAjanlatteteliHatarido"`
	HirdetmenyKozzetetelDatuma           string `json:"hirdetmenyKozzetetelDatuma"`
	EljarasTechnikaiAzonosito            string `json:"eljarasTechnikaiAzonosito"`
	EljarasBeszerzesTargya               string `json:"eljarasBeszerzesTargya"`
	EljarasCPVkod                        []any  `json:"eljarasCPVkod"`
	EljarasNyertesAjanlattevoNeve        any    `json:"eljarasNyertesAjanlattevoNeve"`
	EljarasTeljesitesHelye               string `json:"eljarasTeljesitesHelye"`
	EljarasEljarasRend                   string `json:"eljarasEljarasRend"`
	EljarasEljarasTipusa                 string `json:"eljarasEljarasTipusa"`
	AjanlatkeroTipus                     string `json:"ajanlatkeroTipus"`
	AjanlatkeroKozszolgTevekenysegiKor   any    `json:"ajanlatkeroKozszolgTevekenysegiKor"`
	AjanlatkeroKlasszikusTevekenysegiKor string `json:"ajanlatkeroKlasszikusTevekenysegiKor"`
	HirdetmenyErtesitoLapszam            string `json:"hirdetmenyErtesitoLapszam"`
	DokumentumAzonosito                  string `json:"dokumentumAzonosito"`
	HirdetmenyTedURL                     any    `json:"hirdetmenyTedUrl"`
	TedLapszam                           any    `json:"tedLapszam"`
	TedKozzetetelDatuma                  any    `json:"tedKozzetetelDatuma"`
}

// https://tenders.guru/api/hu/notices/87012568/sections ?

// https://tenders.guru/api/hu/notices/{notice_id}/sections ?

// https://tenders.guru/api/hu/sections/{section_id} ?
