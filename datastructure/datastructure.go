package datastructure

func Run() {

}

func connectDB(c int) (interface{}, error) {

	if c == 1 {
		return "ss", nil
	}

	return 1, nil
}
