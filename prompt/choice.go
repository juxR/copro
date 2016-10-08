package prompt

type Choice struct {
	ID       int
	Label    string
	Type     string
	Selected bool
	pointer  int
}
