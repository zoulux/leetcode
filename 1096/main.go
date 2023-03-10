package main

func main() {

}

func braceExpansionII(expression string) []string {

	for i := 0; i < len(expression); i++ {
		v := expression[i]

		switch v {
		case '{':

		case '}':
		case ',':
		default:

		}

	}
	return nil
}
