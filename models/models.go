package models

type Program struct {
	IDProgram     int             `gorm:"notnull;primaryKey" json:"Id-Program"`
	NameProgram   string          `json:"Name-Program"`
	Schedule      string          `json:"Schedule"`
	Student       []Student       `gorm:"foreignkey:ProgramID; references:IDProgram" json:"Student-Id-Program"`
	Teacher       []Teacher       `gorm:"foreignkey:ProgramID; references:IDProgram" json:"Teacher-Id-Program"`
	ProgramSD     []ProgramSD     `gorm:"foreignkey:ProgramID; references:IDProgram" json:"ProgSD-Id-Program"`
	ProgramSMP    []ProgramSMP    `gorm:"foreignkey:ProgramID; references:IDProgram" json:"ProgSMP-Id-Program"`
	ProgramSMAIpa []ProgramSMAIpa `gorm:"foreignkey:ProgramID; references:IDProgram" json:"ProgSMAIpa-Id-Program"`
	ProgramSMAIps []ProgramSMAIps `gorm:"foreignkey:ProgramID; references:IDProgram" json:"ProgSMAIps-Id-Program"`
}

type ProgramSD struct {
	ProgramID int `json:"Id-Program"`
	StudentID int `json:"Id-Student"`
	Math      int `json:"Score-Math"`
	Sains     int `json:"Score-Sains"`
	Indonesia int `json:"Score-Indonesia"`
	English   int `json:"Score-English"`
}

type ProgramSMP struct {
	ProgramID int `json:"Id-Program"`
	StudentID int `json:"Id-Student"`
	Math      int `json:"Score-Math"`
	Sains     int `json:"Score-Sains"`
	Indonesia int `json:"Score-Indonesia"`
	English   int `json:"Score-English"`
	Sosial    int `json:"Score-Sosial"`
}

type ProgramSMAIpa struct {
	ProgramID int ` json:"Id-Program"`
	StudentID int `json:"Id-Student"`
	Math      int `json:"Score-Math"`
	Fisika    int `json:"Score-Fisika"`
	Kimia     int `json:"Score-Kimia"`
	Biologi   int `json:"Score-Biologi"`
	Indonesia int `json:"Score-Indonesia"`
	English   int `json:"Score-English"`
}

type ProgramSMAIps struct {
	ProgramID int `json:"Id-Program"`
	StudentID int `json:"Id-Student"`
	Math      int `json:"Score-Math"`
	Economy   int `json:"Score-Economy"`
	Geografi  int `json:"Score-Geografi"`
	Sejarah   int `json:"Score-Sejarah"`
	Indonesia int `json:"Score-Indonesia"`
	English   int `json:"Score-English"`
}

type Teacher struct {
	IDTeacher   int    `gorm:"notnull;primaryKey" json:"Id-Teacher"`
	NameTeacher string `json:"Name-Teacher"`
	Age         int    `json:"Age"`
	Number      int    `json:"Number"`
	Salary      int    `json:"Salary"`
	ProgramID   int    `json:"Program"`
}

type Student struct {
	IDStudent     int             `gorm:"notnull;primaryKey" json:"Id-Student"`
	NameStudent   string          `json:"Name-Student"`
	Age           int             `json:"Age"`
	Addres        string          `json:"Addres"`
	Number        uint32          `json:"Number"`
	ProgramID     int             `json:"Program"`
	NameSchool    string          `json:"Name-School"`
	Class         int             `json:"Class"`
	ProgramSD     []ProgramSD     `gorm:"foreignKey:StudentID; references:IDStudent" json:"ProgSD-Id-Student"`
	ProgramSMP    []ProgramSMP    `gorm:"foreignKey:StudentID; references:IDStudent" json:"ProgSMP-Id-Student"`
	ProgramSMAIpa []ProgramSMAIpa `gorm:"foreignKey:StudentID; references:IDStudent" json:"ProgSMAIpa-Id-Student"`
	ProgramSMAIps []ProgramSMAIps `gorm:"foreignKey:StudentID; references:IDStudent" json:"ProgSMAIps-Id-Student"`
}

type Tamp struct {
	IDStudent   int    `json:"Id-Student"`
	NameStudent string `json:"Name-Student"`
	NameSchool  string `json:"Name-School"`
	Class       int    `json:"Class"`
	NameProgram string `json:"Name-Program"`
}


