package find_capital_by_province

var (
	capital = map[string]string{
		"آذربایجان شرقی":      "تبریز",
		"آذربایجان غربی":      "ارومیه",
		"اردبیل":              "اردبیل",
		"اصفهان":              "اصفهان",
		"البرز":               "کرج",
		"ایلام":               "ایلام",
		"بوشهر":               "بوشهر",
		"تهران":               "تهران",
		"چهارمحال و بختیاری":  "شهرکرد",
		"خراسان جنوبی":        "بیرجند",
		"خراسان رضوی":         "مشهد",
		"خراسان شمالی":        "بجنورد",
		"خوزستان":             "اهواز",
		"زنجان":               "زنجان",
		"سمنان":               "سمنان",
		"سیستان و بلوچستان":   "زاهدان",
		"فارس":                "شیراز",
		"قزوین":               "قزوین",
		"قم":                  "قم",
		"کردستان":             "سنندج",
		"کرمان":               "کرمان",
		"کرمانشاه":            "کرمانشاه",
		"کهگیلویه و بویراحمد": "یاسوج",
		"گلستان":              "گرگان",
		"گیلان":               "رشت",
		"لرستان":              "خرم آباد",
		"مازندران":            "ساری",
		"مرکزی":               "اراک",
		"هرمزگان":             "بندرعباس",
		"همدان":               "همدان",
		"یزد":                 "یزد",
	}
)

func FindCapitalByProvince(province string) string {
	if capital[province] == "" {
		return "not found"
	}

	return capital[province]
}
