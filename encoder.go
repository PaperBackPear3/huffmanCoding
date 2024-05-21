package main

type char_with_count struct {
	char  string
	count int
}

func main() {
	// Call the encode function
	encode("Hello World")
}

func encode(s string) ([]char_with_count, error) {
	var char_and_count = get_chars_count(s)
	return char_and_count, nil
}

func get_chars_count(s string) []char_with_count {
	data := []char_with_count{}
	for i := 0; i < len(s); i++ {
		char := string(s[i])
		count := 0
		for j := 0; j < len(s); j++ {
			if char == string(s[j]) {
				count = count + 1
			}
		}
		new_data := char_with_count{char: char, count: count}
		data = append(data, new_data)
	}

	data_to_return := []char_with_count{}
	for i := 0; i < len(data); i++ {
		var should_be_added bool = true
		for j := 0; j < len(data_to_return); j++ {
			if data[i].char == data[j].char {
				should_be_added = false
			}
		}
		if should_be_added {
			data_to_return = append(data_to_return, data[i])
		}

	}
	return data_to_return
}
