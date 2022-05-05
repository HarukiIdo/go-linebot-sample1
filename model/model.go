package model

type (
	// Response APIレスポンス
	Response struct {
		Results Results `json:"results"`
	}
	// Results shop一覧
	Results struct {
		Shop []Shop `json:"shop"`
	}

	// Shop shop情報
	Shop struct {
		Name    string `json:"name"`
		Address string `json:"address"`
		Photo   Photo  `json:"photo"`
		URLs    URLs   `json:"urls"`
	}

	// Photo 写真URL一覧
	Photo struct {
		Mobile Mobile `json:"mobile"`
	}

	// Mobile モバイル用の写真url
	Mobile struct {
		URL string `json:"url"`
	}

	// URLs URL一覧
	URLs struct {
		PC string `json:"pc"`
	}
)
