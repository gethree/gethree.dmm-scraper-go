package dmmCoJp

// DmmCoJpItem : DmmCoJpItem Info Struct
type DmmCoJpItem struct {
	ItemCode             string
	Title                string
	PackageImageThumbURL string
	PackageImageURL      string
	ActressList          []*Actress
	DirectorList         []*Director
	SeriesList           []*Series
	KeywordList          []*Keyword
	SampleImageList      []*SampleImage
}

// Actress : Actress Info Struct
type Actress struct {
	ID   string
	Name string
}

// Director : Director Info Struct
type Director struct {
	ID   string
	Name string
}

// Series : Series Info Struct
type Series struct {
	ID   string
	Name string
}

// Keyword : Keyword Info Struct
type Keyword struct {
	ID   string
	Name string
}

// SampleImage : SampleImage Info Struct
type SampleImage struct {
	ThumbnailURL string
	FullImageURL string
}
