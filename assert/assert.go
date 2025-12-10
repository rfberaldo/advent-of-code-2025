package assert

func NoErr(err error) {
	if err != nil {
		panic(err)
	}
}

func True(v bool) {
	if !v {
		panic("assert/True")
	}
}
