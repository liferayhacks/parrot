package model

type LocaleInfo struct {
	Country  string
	Language string
}

var Locales = map[string]LocaleInfo{
	"sq_AL":       LocaleInfo{Country: "Albanian", Language: "Albania"},
	"ar_DZ":       LocaleInfo{Country: "Arabic", Language: "Algeria"},
	"ar_BH":       LocaleInfo{Country: "Arabic", Language: "Bahrain"},
	"ar_EG":       LocaleInfo{Country: "Arabic", Language: "Egypt"},
	"ar_IQ":       LocaleInfo{Country: "Arabic", Language: "Iraq"},
	"ar_JO":       LocaleInfo{Country: "Arabic", Language: "Jordan"},
	"ar_KW":       LocaleInfo{Country: "Arabic", Language: "Kuwait"},
	"ar_LB":       LocaleInfo{Country: "Arabic", Language: "Lebanon"},
	"ar_LY":       LocaleInfo{Country: "Arabic", Language: "Libya"},
	"ar_MA":       LocaleInfo{Country: "Arabic", Language: "Morocco"},
	"ar_OM":       LocaleInfo{Country: "Arabic", Language: "Oman"},
	"ar_QA":       LocaleInfo{Country: "Arabic", Language: "Qatar"},
	"ar_SA":       LocaleInfo{Country: "Arabic", Language: "Saudi Arabia"},
	"ar_SD":       LocaleInfo{Country: "Arabic", Language: "Sudan"},
	"ar_SY":       LocaleInfo{Country: "Arabic", Language: "Syria"},
	"ar_TN":       LocaleInfo{Country: "Arabic", Language: "Tunisia"},
	"ar_AE":       LocaleInfo{Country: "Arabic", Language: "United Arab Emirates"},
	"ar_YE":       LocaleInfo{Country: "Arabic", Language: "Yemen"},
	"be_BY":       LocaleInfo{Country: "Belarusian", Language: "Belarus"},
	"bg_BG":       LocaleInfo{Country: "Bulgarian", Language: "Bulgaria"},
	"ca_ES":       LocaleInfo{Country: "Catalan", Language: "Spain"},
	"zh_CN":       LocaleInfo{Country: "Chinese (Simplified)", Language: "China"},
	"zh_SG":       LocaleInfo{Country: "Chinese (Simplified)", Language: "Singapore"},
	"zh_HK":       LocaleInfo{Country: "Chinese (Traditional)", Language: "Hong Kong"},
	"zh_TW":       LocaleInfo{Country: "Chinese (Traditional)", Language: "Taiwan"},
	"hr_HR":       LocaleInfo{Country: "Croatian", Language: "Croatia"},
	"cs_CZ":       LocaleInfo{Country: "Czech", Language: "Czech Republic"},
	"da_DK":       LocaleInfo{Country: "Danish", Language: "Denmark"},
	"nl_BE":       LocaleInfo{Country: "Dutch", Language: "Belgium"},
	"nl_NL":       LocaleInfo{Country: "Dutch", Language: "Netherlands"},
	"en_AU":       LocaleInfo{Country: "English", Language: "Australia"},
	"en_CA":       LocaleInfo{Country: "English", Language: "Canada"},
	"en_IN":       LocaleInfo{Country: "English", Language: "India"},
	"en_IE":       LocaleInfo{Country: "English", Language: "Ireland"},
	"en_MT":       LocaleInfo{Country: "English", Language: "Malta"},
	"en_NZ":       LocaleInfo{Country: "English", Language: "New Zealand"},
	"en_PH":       LocaleInfo{Country: "English", Language: "Philippines"},
	"en_SG":       LocaleInfo{Country: "English", Language: "Singapore"},
	"en_ZA":       LocaleInfo{Country: "English", Language: "South Africa"},
	"en_GB":       LocaleInfo{Country: "English", Language: "United Kingdom"},
	"en_US":       LocaleInfo{Country: "English", Language: "United States"},
	"et_EE":       LocaleInfo{Country: "Estonian", Language: "Estonia"},
	"fi_FI":       LocaleInfo{Country: "Finnish", Language: "Finland"},
	"fr_BE":       LocaleInfo{Country: "French", Language: "Belgium"},
	"fr_CA":       LocaleInfo{Country: "French", Language: "Canada"},
	"fr_FR":       LocaleInfo{Country: "French", Language: "France"},
	"fr_LU":       LocaleInfo{Country: "French", Language: "Luxembourg"},
	"fr_CH":       LocaleInfo{Country: "French", Language: "Switzerland"},
	"de_AT":       LocaleInfo{Country: "German", Language: "Austria"},
	"de_DE":       LocaleInfo{Country: "German", Language: "Germany"},
	"de_LU":       LocaleInfo{Country: "German", Language: "Luxembourg"},
	"de_CH":       LocaleInfo{Country: "German", Language: "Switzerland"},
	"el_CY":       LocaleInfo{Country: "Greek", Language: "Cyprus"},
	"el_GR":       LocaleInfo{Country: "Greek", Language: "Greece"},
	"iw_IL":       LocaleInfo{Country: "Hebrew", Language: "Israel"},
	"hi_IN":       LocaleInfo{Country: "Hindi", Language: "India"},
	"hu_HU":       LocaleInfo{Country: "Hungarian", Language: "Hungary"},
	"is_IS":       LocaleInfo{Country: "Icelandic", Language: "Iceland"},
	"in_ID":       LocaleInfo{Country: "Indonesian", Language: "Indonesia"},
	"ga_IE":       LocaleInfo{Country: "Irish", Language: "Ireland"},
	"it_IT":       LocaleInfo{Country: "Italian", Language: "Italy"},
	"it_CH":       LocaleInfo{Country: "Italian", Language: "Switzerland"},
	"ja_JP":       LocaleInfo{Country: "Japanese", Language: "Japan"},
	"ko_KR":       LocaleInfo{Country: "Korean", Language: "South Korea"},
	"lv_LV":       LocaleInfo{Country: "Latvian", Language: "Latvia"},
	"lt_LT":       LocaleInfo{Country: "Lithuanian", Language: "Lithuania"},
	"mk_MK":       LocaleInfo{Country: "Macedonian", Language: "Macedonia"},
	"ms_MY":       LocaleInfo{Country: "Malay", Language: "Malaysia"},
	"mt_MT":       LocaleInfo{Country: "Maltese", Language: "Malta"},
	"no_NO":       LocaleInfo{Country: "Norwegian (Bokmål)", Language: "Norway"},
	"no_NO_NY":    LocaleInfo{Country: "Norwegian (Nynorsk)", Language: "Norway"},
	"pl_PL":       LocaleInfo{Country: "Polish", Language: "Poland"},
	"pt_BR":       LocaleInfo{Country: "Portuguese", Language: "Brazil"},
	"pt_PT":       LocaleInfo{Country: "Portuguese", Language: "Portugal"},
	"ro_RO":       LocaleInfo{Country: "Romanian", Language: "Romania"},
	"ru_RU":       LocaleInfo{Country: "Russian", Language: "Russia"},
	"sr_BA":       LocaleInfo{Country: "Serbian (Cyrillic)", Language: "Bosnia and Herzegovina"},
	"sr_ME":       LocaleInfo{Country: "Serbian (Cyrillic)", Language: "Montenegro"},
	"sr_RS":       LocaleInfo{Country: "Serbian (Cyrillic)", Language: "Serbia"},
	"sr_La_tn_BA": LocaleInfo{Country: "Serbian (Latin)", Language: "Bosnia and Herzegovina"},
	"sr_La_tn_ME": LocaleInfo{Country: "Serbian (Latin)", Language: "Montenegro"},
	"sr_La_tn_RS": LocaleInfo{Country: "Serbian (Latin)", Language: "Serbia"},
	"sk_SK":       LocaleInfo{Country: "Slovak", Language: "Slovakia"},
	"sl_SI":       LocaleInfo{Country: "Slovenian", Language: "Slovenia"},
	"es_AR":       LocaleInfo{Country: "Spanish", Language: "Argentina"},
	"es_BO":       LocaleInfo{Country: "Spanish", Language: "Bolivia"},
	"es_CL":       LocaleInfo{Country: "Spanish", Language: "Chile"},
	"es_CO":       LocaleInfo{Country: "Spanish", Language: "Colombia"},
	"es_CR":       LocaleInfo{Country: "Spanish", Language: "Costa Rica"},
	"es_DO":       LocaleInfo{Country: "Spanish", Language: "Dominican Republic"},
	"es_EC":       LocaleInfo{Country: "Spanish", Language: "Ecuador"},
	"es_SV":       LocaleInfo{Country: "Spanish", Language: "El Salvador"},
	"es_GT":       LocaleInfo{Country: "Spanish", Language: "Guatemala"},
	"es_HN":       LocaleInfo{Country: "Spanish", Language: "Honduras"},
	"es_MX":       LocaleInfo{Country: "Spanish", Language: "Mexico"},
	"es_NI":       LocaleInfo{Country: "Spanish", Language: "Nicaragua"},
	"es_PA":       LocaleInfo{Country: "Spanish", Language: "Panama"},
	"es_PY":       LocaleInfo{Country: "Spanish", Language: "Paraguay"},
	"es_PE":       LocaleInfo{Country: "Spanish", Language: "Peru"},
	"es_PR":       LocaleInfo{Country: "Spanish", Language: "Puerto Rico"},
	"es_ES":       LocaleInfo{Country: "Spanish", Language: "Spain"},
	"es_US":       LocaleInfo{Country: "Spanish", Language: "United States"},
	"es_UY":       LocaleInfo{Country: "Spanish", Language: "Uruguay"},
	"es_VE":       LocaleInfo{Country: "Spanish", Language: "Venezuela"},
	"sv_SE":       LocaleInfo{Country: "Swedish", Language: "Sweden"},
	"th_TH":       LocaleInfo{Country: "Thai (Western digits)", Language: "Thailand"},
	"th_TH_TH":    LocaleInfo{Country: "Thai (Thai digits)", Language: "Thailand"},
	"tr_TR":       LocaleInfo{Country: "Turkish", Language: "Turkey"},
	"uk_UA":       LocaleInfo{Country: "Ukrainian", Language: "Ukraine"},
	"vi_VN":       LocaleInfo{Country: "Vietnamese", Language: "Vietnam"},
}
