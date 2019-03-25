package oxforddict

type OxfordSuccessResponse struct {
	// Metadata struct {
	// } `json:"metadata"`
	Results []Result `json:"results"`
}

type Result struct {
	// ID             string          `json:"id"`
	Language       string         `json:"language"`
	LexicalEntries []LexicalEntry `json:"lexicalEntries"`
	// Pronunciations []Pronunciation `json:"pronunciations"`
	Type string `json:"type"`
	Word string `json:"word"`
}

type LexicalEntry struct {
	// DerivativeOf        []RelatedEntry       `json:"derivativeOf"`
	// Derivatives         []RelatedEntry       `json:"derivatives"`
	Entries             []Entry              `json:"entries"`
	GrammaticalFeatures []GrammaticalFeature `json:"grammaticalFeatures"`
	// Language            string               `json:"language"`
	LexicalCategory string          `json:"lexicalCategory"`
	Notes           []Note          `json:"notes"`
	Pronunciations  []Pronunciation `json:"pronunciations"`
	// Text                string               `json:"text"`
	// VariantForms        []VariantForm        `json:"variantForms"`
}

type Pronunciation struct {
	AudioFile string `json:"audioFile"`
	// Dialects         []string `json:"dialects"`
	PhoneticNotation string `json:"phoneticNotation"`
	PhoneticSpelling string `json:"phoneticSpelling"`
	// Regions          []string `json:"regions"`
}

type VariantForm struct {
	Regions []string `json:"regions"`
	Text    string   `json:"text"`
}

type Note struct {
	// ID   string `json:"id"`
	Text string `json:"text"`
	Type string `json:"type"`
}

type GrammaticalFeature struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

type Translation struct {
	Domains             []string             `json:"domains"`
	GrammaticalFeatures []GrammaticalFeature `json:"grammaticalFeatures"`
	Language            string               `json:"language"`
	Notes               []Note               `json:"notes"`
	Regions             []string             `json:"regions"`
	Registers           []string             `json:"registers"`
	Text                string               `json:"text"`
}

type ThesaurusLink struct {
	EntryID string `json:"entry_id"`
	SenseID string `json:"sense_id"`
}

type Example struct {
	// Definitions  []string      `json:"definitions"`
	// Domains      []string      `json:"domains"`
	// Notes        []Note        `json:"notes"`
	// Regions      []string      `json:"regions"`
	// Registers    []string      `json:"registers"`
	// SenseIds     []string      `json:"senseIds"`
	Text string `json:"text"`
	// Translations []Translation `json:"translations"`
}

type CrossReference struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Type string `json:"type"`
}

type Sense struct {
	// CrossReferenceMarkers []string         `json:"crossReferenceMarkers"`
	// CrossReferences       []CrossReference `json:"crossReferences"`
	Definitions []string `json:"definitions"`
	// Domains               []string         `json:"domains"`
	Examples []Example `json:"examples"`
	// ID                    string           `json:"id"`
	// Notes                 []Note           `json:"notes"`
	Pronunciations []Pronunciation `json:"pronunciations"`
	// Regions               []string         `json:"regions"`
	// Registers             []string         `json:"registers"`
	ShortDefinitions []string `json:"short_definitions"`
	// Subsenses             []Sense          `json:"subsenses"`
	// ThesaurusLinks        []ThesaurusLink  `json:"thesaurusLinks"`
	// Translations          []Translation    `json:"translations"`
	// VariantForms          []VariantForm    `json:"variantForms"`
}

type RelatedEntry struct {
	Domains   []string `json:"domains"`
	ID        string   `json:"id"`
	Language  string   `json:"language"`
	Regions   []string `json:"regions"`
	Registers []string `json:"registers"`
	Text      string   `json:"text"`
}

type Entry struct {
	Etymologies []string `json:"etymologies"`
	// GrammaticalFeatures []GrammaticalFeature `json:"grammaticalFeatures"`
	// HomographNumber     string               `json:"homographNumber"`
	Notes []Note `json:"notes"`
	// Pronunciations []Pronunciation `json:"pronunciations"`
	Senses []Sense `json:"senses"`
	// VariantForms        []VariantForm        `json:"variantForms"`
}

type BriefWordInfo struct {
	Word           string
	Language       string
	Type           string
	LexicalEntries []struct {
		Definitions      []string
		ShortDefinitions string
		LexicalCategory  string
		PhoneticSpelling string
		Examples         string
	}
}
