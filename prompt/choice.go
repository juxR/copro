package prompt

type Choice struct {
	ID          int
	Label       string
	Type        string
	Selected    bool
	IsSeparator bool
	pointer     int
}
