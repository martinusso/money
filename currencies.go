package money

type Currency string

const (
	AED Currency = "AED"
	AFN Currency = "AFN"
	ALL Currency = "ALL"
	AMD Currency = "AMD"
	ANG Currency = "ANG"
	AOA Currency = "AOA"
	ARS Currency = "ARS"
	AUD Currency = "AUD"
	AWG Currency = "AWG"
	AZN Currency = "AZN"
	BAM Currency = "BAM"
	BBD Currency = "BBD"
	BDT Currency = "BDT"
	BGN Currency = "BGN"
	BHD Currency = "BHD"
	BIF Currency = "BIF"
	BMD Currency = "BMD"
	BND Currency = "BND"
	BOB Currency = "BOB"
	BOV Currency = "BOV"
	BRL Currency = "BRL"
	BSD Currency = "BSD"
	BTN Currency = "BTN"
	BWP Currency = "BWP"
	BYN Currency = "BYN"
	BYR Currency = "BYR"
	BZD Currency = "BZD"
	CAD Currency = "CAD"
	CDF Currency = "CDF"
	CHF Currency = "CHF"
	CLF Currency = "CLF"
	CLP Currency = "CLP"
	CNY Currency = "CNY"
	COP Currency = "COP"
	COU Currency = "COU"
	CRC Currency = "CRC"
	CUC Currency = "CUC"
	CUP Currency = "CUP"
	CVE Currency = "CVE"
	CZK Currency = "CZK"
	DJF Currency = "DJF"
	DKK Currency = "DKK"
	DOP Currency = "DOP"
	DZD Currency = "DZD"
	EEK Currency = "EEK"
	EGP Currency = "EGP"
	ERN Currency = "ERN"
	ETB Currency = "ETB"
	EUR Currency = "EUR"
	FJD Currency = "FJD"
	FKP Currency = "FKP"
	GBP Currency = "GBP"
	GEL Currency = "GEL"
	GHS Currency = "GHS"
	GIP Currency = "GIP"
	GMD Currency = "GMD"
	GNF Currency = "GNF"
	GTQ Currency = "GTQ"
	GYD Currency = "GYD"
	HKD Currency = "HKD"
	HNL Currency = "HNL"
	HRK Currency = "HRK"
	HTG Currency = "HTG"
	HUF Currency = "HUF"
	IDR Currency = "IDR"
	ILS Currency = "ILS"
	INR Currency = "INR"
	IQD Currency = "IQD"
	IRR Currency = "IRR"
	ISK Currency = "ISK"
	JMD Currency = "JMD"
	JOD Currency = "JOD"
	JPY Currency = "JPY"
	KES Currency = "KES"
	KGS Currency = "KGS"
	KHR Currency = "KHR"
	KMF Currency = "KMF"
	KPW Currency = "KPW"
	KRW Currency = "KRW"
	KWD Currency = "KWD"
	KYD Currency = "KYD"
	KZT Currency = "KZT"
	LAK Currency = "LAK"
	LBP Currency = "LBP"
	LKR Currency = "LKR"
	LRD Currency = "LRD"
	LSL Currency = "LSL"
	LTL Currency = "LTL"
	LVL Currency = "LVL"
	LYD Currency = "LYD"
	MAD Currency = "MAD"
	MDL Currency = "MDL"
	MGA Currency = "MGA"
	MKD Currency = "MKD"
	MMK Currency = "MMK"
	MNT Currency = "MNT"
	MOP Currency = "MOP"
	MRO Currency = "MRO"
	MUR Currency = "MUR"
	MVR Currency = "MVR"
	MWK Currency = "MWK"
	MXN Currency = "MXN"
	MXV Currency = "MXV"
	MYR Currency = "MYR"
	MZN Currency = "MZN"
	NAD Currency = "NAD"
	NGN Currency = "NGN"
	NIO Currency = "NIO"
	NOK Currency = "NOK"
	NPR Currency = "NPR"
	NZD Currency = "NZD"
	OMR Currency = "OMR"
	PAB Currency = "PAB"
	PEN Currency = "PEN"
	PGK Currency = "PGK"
	PHP Currency = "PHP"
	PKR Currency = "PKR"
	PLN Currency = "PLN"
	PYG Currency = "PYG"
	QAR Currency = "QAR"
	RON Currency = "RON"
	RSD Currency = "RSD"
	RUB Currency = "RUB"
	RWF Currency = "RWF"
	SAR Currency = "SAR"
	SBD Currency = "SBD"
	SCR Currency = "SCR"
	SDG Currency = "SDG"
	SEK Currency = "SEK"
	SGD Currency = "SGD"
	SHP Currency = "SHP"
	SLL Currency = "SLL"
	SOS Currency = "SOS"
	SRD Currency = "SRD"
	STD Currency = "STD"
	SVC Currency = "SVC"
	SYP Currency = "SYP"
	SZL Currency = "SZL"
	THB Currency = "THB"
	TJS Currency = "TJS"
	TMT Currency = "TMT"
	TND Currency = "TND"
	TOP Currency = "TOP"
	TRY Currency = "TRY"
	TTD Currency = "TTD"
	TWD Currency = "TWD"
	TZS Currency = "TZS"
	UAH Currency = "UAH"
	UGX Currency = "UGX"
	USD Currency = "USD"
	UYI Currency = "UYI"
	UYU Currency = "UYU"
	UZS Currency = "UZS"
	VEF Currency = "VEF"
	VND Currency = "VND"
	VUV Currency = "VUV"
	WST Currency = "WST"
	XAF Currency = "XAF"
	XAG Currency = "XAG"
	XAU Currency = "XAU"
	XBA Currency = "XBA"
	XBB Currency = "XBB"
	XBC Currency = "XBC"
	XBD Currency = "XBD"
	XCD Currency = "XCD"
	XDR Currency = "XDR"
	XFU Currency = "XFU"
	XOF Currency = "XOF"
	XPD Currency = "XPD"
	XPF Currency = "XPF"
	XPT Currency = "XPT"
	XTS Currency = "XTS"
	YER Currency = "YER"
	ZAR Currency = "ZAR"
	ZMK Currency = "ZMK"
	ZWL Currency = "ZWL"
)

var (
	currencies = map[Currency]pattern{
		AED: pattern{name: "UAE Dirham", symbol: "د.إ", precision: 2},
		AFN: pattern{name: "Afghani", symbol: "؋", precision: 2},
		ALL: pattern{name: "Lek", symbol: "Lek", precision: 2},
		AMD: pattern{name: "Armenian Dram", symbol: "AMD", precision: 2},
		ANG: pattern{name: "Netherlands Antillian Guilder", symbol: "ƒ", precision: 2},
		AOA: pattern{name: "Kwanza", symbol: "Kz", precision: 2},
		ARS: pattern{name: "Argentine Peso", symbol: "$", precision: 2},
		AUD: pattern{name: "Australian Dollar", symbol: "$", precision: 2},
		AWG: pattern{name: "Aruban Guilder", symbol: "ƒ", precision: 2},
		AZN: pattern{name: "Azerbaijanian Manat", symbol: "ман", precision: 2},
		BAM: pattern{name: "Convertible Marks", symbol: "KM", precision: 2},
		BBD: pattern{name: "Barbados Dollar", symbol: "$", precision: 2},
		BDT: pattern{name: "Taka", symbol: "৳", precision: 2},
		BGN: pattern{name: "Bulgarian Lev", symbol: "лв", precision: 2},
		BHD: pattern{name: "Bahraini Dinar", symbol: ".د.ب", precision: 3},
		BIF: pattern{name: "Burundi Franc", symbol: "FBu", precision: 0},
		BMD: pattern{name: "Bermudian Dollar", symbol: "$", precision: 2},
		BND: pattern{name: "Brunei Dollar", symbol: "$", precision: 2},
		BOB: pattern{name: "Boliviano Mvdol", symbol: "$b", precision: 2},
		BOV: pattern{name: "Boliviano Mvdol", symbol: "$b", precision: 2},
		BRL: pattern{name: "Brazilian Real", symbol: "R$", precision: 2, thousand: '.', decimal: ','},
		BSD: pattern{name: "Bahamian Dollar", symbol: "$", precision: 2},
		BTN: pattern{name: "Indian Rupee Ngultrum", symbol: "Nu.", precision: 2},
		BWP: pattern{name: "Pula", symbol: "P", precision: 2},
		BYN: pattern{name: "Belarusian Ruble", symbol: "p.", precision: 2},
		BYR: pattern{name: "Belarusian Ruble", symbol: "p.", precision: 0},
		BZD: pattern{name: "Belize Dollar", symbol: "BZ$", precision: 2},
		CAD: pattern{name: "Canadian Dollar", symbol: "$", precision: 2},
		CDF: pattern{name: "Congolese Franc", symbol: "CF", precision: 2},
		CHF: pattern{name: "Swiss Franc", symbol: "CHF", precision: 2},
		CLF: pattern{name: "Chilean Peso Unidades de fomento", symbol: "$", precision: 4},
		CLP: pattern{name: "Chilean Peso Unidades de fomento", symbol: "$", precision: 0},
		CNY: pattern{name: "Yuan Renminbi", symbol: "¥", precision: 2},
		COP: pattern{name: "Colombian Peso", symbol: "$", precision: 2},
		COU: pattern{name: "Colombian Peso Unidad de Valor Real", symbol: "$", precision: 2},
		CRC: pattern{name: "Costa Rican Colon", symbol: "₡", precision: 2},
		CUC: pattern{name: "Cuban Peso Peso Convertible", symbol: "₱", precision: 2},
		CUP: pattern{name: "Cuban Peso Peso Convertible", symbol: "₱", precision: 2},
		CVE: pattern{name: "Cape Verde Escudo", symbol: "$", precision: 0},
		CZK: pattern{name: "Czech Koruna", symbol: "Kč", precision: 2},
		DJF: pattern{name: "Djibouti Franc", symbol: "Fdj", precision: 0},
		DKK: pattern{name: "Danish Krone", symbol: "kr.", precision: 2},
		DOP: pattern{name: "Dominican Peso", symbol: "RD$", precision: 2},
		DZD: pattern{name: "Algerian Dinar", symbol: "دج", precision: 2},
		EEK: pattern{name: "Kroon", symbol: "KR", precision: 2},
		EGP: pattern{name: "Egyptian Pound", symbol: "£", precision: 2},
		ERN: pattern{name: "Nakfa", symbol: "Nfk", precision: 2},
		ETB: pattern{name: "Ethiopian Birr", symbol: "Br", precision: 2},
		EUR: pattern{name: "Euro", symbol: "€", precision: 2},
		FJD: pattern{name: "Fiji Dollar", symbol: "$", precision: 2},
		FKP: pattern{name: "Falkland Islands Pound", symbol: "£", precision: 2},
		GBP: pattern{name: "Pound Sterling", symbol: "£", precision: 2},
		GEL: pattern{name: "Lari", symbol: "₾", precision: 2},
		GHS: pattern{name: "Cedi", symbol: "GH₵", precision: 2},
		GIP: pattern{name: "Gibraltar Pound", symbol: "£", precision: 2},
		GMD: pattern{name: "Dalasi", symbol: "D", precision: 2},
		GNF: pattern{name: "Guinea Franc", symbol: "FG", precision: 0},
		GTQ: pattern{name: "Quetzal", symbol: "Q", precision: 2},
		GYD: pattern{name: "Guyana Dollar", symbol: "$", precision: 2},
		HKD: pattern{name: "Hong Kong Dollar", symbol: "$", precision: 2},
		HNL: pattern{name: "Lempira", symbol: "L", precision: 2},
		HRK: pattern{name: "Croatian Kuna", symbol: "kn", precision: 2},
		HTG: pattern{name: "Gourde US Dollar", symbol: " ", precision: 2},
		HUF: pattern{name: "Forint", symbol: "Ft", precision: 2},
		IDR: pattern{name: "Rupiah", symbol: "Rp", precision: 2},
		ILS: pattern{name: "New Israeli Sheqel", symbol: "₪", precision: 2},
		INR: pattern{name: "Indian Rupee", symbol: "₹", precision: 2},
		IQD: pattern{name: "Iraqi Dinar", symbol: "ع.د", precision: 3},
		IRR: pattern{name: "Iranian Rial", symbol: "﷼", precision: 2},
		ISK: pattern{name: "Iceland Krona", symbol: "kr", precision: 0},
		JMD: pattern{name: "Jamaican Dollar", symbol: "J$", precision: 2},
		JOD: pattern{name: "Jordanian Dinar", symbol: "JOD", precision: 3},
		JPY: pattern{name: "Yen", symbol: "¥", precision: 0},
		KES: pattern{name: "Kenyan Shilling", symbol: "KSh", precision: 2},
		KGS: pattern{name: "Som", symbol: "лв", precision: 2},
		KHR: pattern{name: "Riel", symbol: "៛", precision: 2},
		KMF: pattern{name: "Comoro Franc", symbol: "CF", precision: 0},
		KPW: pattern{name: "North Korean Won", symbol: "₩", precision: 2},
		KRW: pattern{name: "Won", symbol: "₩", precision: 0},
		KWD: pattern{name: "Kuwaiti Dinar", symbol: "د.ك", precision: 3},
		KYD: pattern{name: "Cayman Islands Dollar", symbol: "$", precision: 2},
		KZT: pattern{name: "Tenge", symbol: "лв", precision: 2},
		LAK: pattern{name: "Kip", symbol: "₭", precision: 2},
		LBP: pattern{name: "Lebanese Pound", symbol: "£", precision: 2},
		LKR: pattern{name: "Sri Lanka Rupee", symbol: "₨", precision: 2},
		LRD: pattern{name: "Liberian Dollar", symbol: "$", precision: 2},
		LSL: pattern{name: "Rand Loti", symbol: " ", precision: 2},
		LTL: pattern{name: "Lithuanian Litas", symbol: "Lt", precision: 2},
		LVL: pattern{name: "Latvian Lats", symbol: "Ls", precision: 2},
		LYD: pattern{name: "Libyan Dinar", symbol: "ل.د", precision: 3},
		MAD: pattern{name: "Moroccan Dirham", symbol: "د.م.", precision: 2},
		MDL: pattern{name: "Moldovan Leu", symbol: "MDL", precision: 2},
		MGA: pattern{name: "Malagasy Ariary", symbol: "Ar", precision: 2},
		MKD: pattern{name: "Denar", symbol: "ден", precision: 2},
		MMK: pattern{name: "Kyat", symbol: "K", precision: 2},
		MNT: pattern{name: "Tugrik", symbol: "₮", precision: 2},
		MOP: pattern{name: "Pataca", symbol: "MOP$", precision: 2},
		MRO: pattern{name: "Ouguiya", symbol: "UM", precision: 2},
		MUR: pattern{name: "Mauritius Rupee", symbol: "₨", precision: 2},
		MVR: pattern{name: "Rufiyaa", symbol: "Rf", precision: 2},
		MWK: pattern{name: "Kwacha", symbol: "MK", precision: 2},
		MXN: pattern{name: "Mexican Peso", symbol: "$", precision: 2},
		MXV: pattern{name: "Mexican Peso Mexican Unidad de Inversion (UDI)", symbol: "UDI", precision: 2},
		MYR: pattern{name: "Malaysian Ringgit", symbol: "RM", precision: 2},
		MZN: pattern{name: "Metical", symbol: "MT", precision: 2},
		NAD: pattern{name: "Rand Namibia Dollar", symbol: "$", precision: 2},
		NGN: pattern{name: "Naira", symbol: "₦", precision: 2},
		NIO: pattern{name: "Cordoba Oro", symbol: "C$", precision: 2},
		NOK: pattern{name: "Norwegian Krone", symbol: "kr", precision: 2},
		NPR: pattern{name: "Nepalese Rupee", symbol: "₨", precision: 2},
		NZD: pattern{name: "New Zealand Dollar", symbol: "$", precision: 2},
		OMR: pattern{name: "Rial Omani", symbol: "﷼", precision: 3},
		PAB: pattern{name: "Balboa US Dollar", symbol: "B/.", precision: 2},
		PEN: pattern{name: "Nuevo Sol", symbol: "S/.", precision: 2},
		PGK: pattern{name: "Kina", symbol: "K", precision: 2},
		PHP: pattern{name: "Philippine Peso", symbol: "₱", precision: 2},
		PKR: pattern{name: "Pakistan Rupee", symbol: "₨", precision: 2},
		PLN: pattern{name: "Zloty", symbol: "zł", precision: 2},
		PYG: pattern{name: "Guarani", symbol: "₲", precision: 0},
		QAR: pattern{name: "Qatari Rial", symbol: "﷼", precision: 2},
		RON: pattern{name: "New Leu", symbol: "lei", precision: 2},
		RSD: pattern{name: "Serbian Dinar", symbol: "Дин.", precision: 2},
		RUB: pattern{name: "Russian Ruble", symbol: "₽", precision: 2},
		RWF: pattern{name: "Rwanda Franc", symbol: " ", precision: 0},
		SAR: pattern{name: "Saudi Riyal", symbol: "﷼", precision: 2},
		SBD: pattern{name: "Solomon Islands Dollar", symbol: "$", precision: 2},
		SCR: pattern{name: "Seychelles Rupee", symbol: "₨", precision: 2},
		SDG: pattern{name: "Sudanese Pound", symbol: "SDG", precision: 2},
		SEK: pattern{name: "Swedish Krona", symbol: "kr", precision: 2},
		SGD: pattern{name: "Singapore Dollar", symbol: "$", precision: 2},
		SHP: pattern{name: "Saint Helena Pound", symbol: "£", precision: 2},
		SLL: pattern{name: "Leone", symbol: "Le", precision: 2},
		SOS: pattern{name: "Somali Shilling", symbol: "S", precision: 2},
		SRD: pattern{name: "Surinam Dollar", symbol: "$", precision: 2},
		STD: pattern{name: "Dobra", symbol: "Db", precision: 2},
		SVC: pattern{name: "El Salvador Colon US Dollar", symbol: "$", precision: 2},
		SYP: pattern{name: "Syrian Pound", symbol: "£", precision: 2},
		SZL: pattern{name: "Lilangeni", symbol: "E", precision: 2},
		THB: pattern{name: "Baht", symbol: "฿", precision: 2},
		TJS: pattern{name: "Somoni", symbol: " ", precision: 2},
		TMT: pattern{name: "Manat", symbol: "₼", precision: 2},
		TND: pattern{name: "Tunisian Dinar", symbol: "د.ت", precision: 2},
		TOP: pattern{name: "Pa'anga", symbol: "T$", precision: 2},
		TRY: pattern{name: "Turkish Lira", symbol: "TL", precision: 2},
		TTD: pattern{name: "Trinidad and Tobago Dollar", symbol: "TT$", precision: 2},
		TWD: pattern{name: "New Taiwan Dollar", symbol: "NT$", precision: 2},
		TZS: pattern{name: "Tanzanian Shilling", symbol: "Tsh", precision: 2},
		UAH: pattern{name: "Hryvnia", symbol: "₴", precision: 2},
		UGX: pattern{name: "Uganda Shilling", symbol: "Ush", precision: 0},
		USD: pattern{name: "US Dollar", symbol: "$", precision: 2},
		UYI: pattern{name: "Peso Uruguayo Uruguay Peso en Unidades Indexadas", symbol: "$U", precision: 0},
		UYU: pattern{name: "Peso Uruguayo Uruguay Peso en Unidades Indexadas", symbol: "$U", precision: 2},
		UZS: pattern{name: "Uzbekistan Sum", symbol: "лв", precision: 2},
		VEF: pattern{name: "Bolivar Fuerte", symbol: "Bs", precision: 2},
		VND: pattern{name: "Dong", symbol: "₫", precision: 0},
		VUV: pattern{name: "Vatu", symbol: "VT", precision: 0},
		WST: pattern{name: "Tala", symbol: "WS$", precision: 2},
		XAF: pattern{name: "CFA Franc BEAC", symbol: "FCFA", precision: 0},
		XAG: pattern{name: "Silver", symbol: " ", precision: 2},
		XAU: pattern{name: "Gold", symbol: " ", precision: 2},
		XBA: pattern{name: "Bond Markets Units European Composite Unit (EURCO)", symbol: " ", precision: 2},
		XBB: pattern{name: "European Monetary Unit (E.M.U.-6)", symbol: " ", precision: 2},
		XBC: pattern{name: "European Unit of Account 9(E.U.A.-9)", symbol: " ", precision: 2},
		XBD: pattern{name: "European Unit of Account 17(E.U.A.-17)", symbol: " ", precision: 2},
		XCD: pattern{name: "East Caribbean Dollar", symbol: "$", precision: 2},
		XDR: pattern{name: "SDR", symbol: " ", precision: 2},
		XFU: pattern{name: "UIC-Franc", symbol: " ", precision: 2},
		XOF: pattern{name: "CFA Franc BCEAO", symbol: " ", precision: 0},
		XPD: pattern{name: "Palladium", symbol: " ", precision: 2},
		XPF: pattern{name: "CFP Franc", symbol: " ", precision: 0},
		XPT: pattern{name: "Platinum", symbol: " ", precision: 2},
		XTS: pattern{name: "Codes specifically reserved for testing purposes", symbol: " ", precision: 2},
		YER: pattern{name: "Yemeni Rial", symbol: "﷼", precision: 2},
		ZAR: pattern{name: "Rand", symbol: "R", precision: 2},
		ZMK: pattern{name: "Zambian Kwacha", symbol: "ZK", precision: 2},
		ZWL: pattern{name: "Zimbabwe Dollar", symbol: "$", precision: 2},
	}
)

type pattern struct {
	name      string
	thousand  byte
	decimal   byte
	precision int
	symbol    string
}
