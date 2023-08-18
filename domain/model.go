package domain

type Product struct {
	ID                   string   `json:"id" bson:"_id,omitempty"`
	Title                string   `json:"title" bson:"title,omitempty"`
	Photos               []string `json:"photos" bson:"photos,omitempty"`
	Video                string   `json:"video" bson:"video,omitempty"`
	Cover                string   `json:"cover" bson:"cover,omitempty"`
	Materials            []string `json:"materials" bson:"materials,omitempty"`
	Weight               float64  `json:"weight" bson:"weight,omitempty"`
	Size                 string   `json:"size" bson:"size,omitempty"`
	Category             string   `json:"category" bson:"category,omitempty"`
	Brand                string   `json:"brand" bson:"brand,omitempty"`
	Description          string   `json:"description" bson:"description,omitempty"`
	ManufacturingCountry string   `json:"manufacturing_country" bson:"manufacturing_country,omitempty"`
	PackageType          string   `json:"package_type" bson:"package_type,omitempty"`
	UseCases             []string `json:"use_cases" bson:"use_cases,omitempty"`
	Stock                int      `json:"stock" bson:"stock,omitempty"`
	Price                float64  `json:"price" bson:"price,omitempty"`
	FactorType           string   `json:"factor_type" bson:"factor_type,omitempty"`
	CreatedAt            int64    `json:"created_at" bson:"created_at,omitempty"`
	Status               int      `json:"status" bson:"status,omitempty"` // -1.removed
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

type Filter struct {
	Field    string `json:"field"`
	Value    any    `json:"value"`
	Operator string `json:"operator"`
	// logic
	// id
}

type Pagination struct {
	Skip  int64 `json:"skip"`
	Limit int64 `json:"limit"`
}

type Response struct {
	Error *Error         `json:"errors,omitempty"`
	Data  any            `json:"data,omitempty"`
	Meta  map[string]any `json:"meta,omitempty"`
}
