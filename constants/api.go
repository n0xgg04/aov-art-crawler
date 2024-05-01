package constants

import "golang.org/x/exp/maps"

var (
	ArtAPI = map[string]string{
		"cn": "https://image.ngame.proximabeta.com/eoa/CNEXP/CHS/HeroTrainingLoading/##ID##.jpg",
		"tw": "https://dl.ops.kgtw.garenanow.com/CHT/HeroTrainingLoadingNew_B36/##ID##.jpg",
		"vn": "https://dl.ops.kgvn.garenanow.com/hok/VN/SkinShowPic_B36/##ID##.jpg",
		"th": "https://dl.ops.kgth.garenanow.com/TH/HeroTrainingLoading36/##ID##.jpg",
		"jp": "https://image.ngame.proximabeta.com/eoa/Languages/JP_Tencent_JP/image/JA/HeroTrainingLoading/##ID##.jpg",
		"us": "http://image.ngame.proximabeta.com/eoa/Tencent/EU/EN/HeroTrainingLoading/##ID##.jpg",
	}

	API = maps.Keys(ArtAPI)
)
