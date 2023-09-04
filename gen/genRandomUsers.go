package gen

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var (
	chars       = "abcdefghijklmnopqrstuvwxyz"
	nums 		= "0123456789"
	passChars	= chars + nums + strings.ToUpper(chars)
	idLength    = 8
	userLength  = 8
	emailDomain = "example.com"
	passwordLength = 16
	firstNames  = []string{"Alice", "Bob", "Charlie", "David", "Eve", "Frank", "Grace", "Hannah", "Ivy", "Jack", "Pedro", "Mario", "Leo", "Anna", "Maria"}
	lastNames   = []string{"Smith", "Johnson", "Brown", "Lee", "Davis", "Wilson", "Evans", "Clark", "Wright", "Walker", "Nunes", "Pereira", "Silva", "Gomes"}
	possibleAdress = []string{"Visconde de Piraja", "Maria Helena", "Marielle Franco", "Sacadura Cabral", "Maua", "Neves Ferreira", "28 de Setembro", "Quirino Lopes"}
	possibleAdressComplement = []string{"Rua", "Praceta", "Largo"}
)

func generateRandomPassword(length int) string {
	result := make([]byte, length)
	for i := range result {
		result[i] = passChars[rand.Intn(len(passChars))]
	}
	return string(result)
}

func generateRandomNumber(length int) string {
	result := make([]byte, length)
	for i := range result {
		result[i] = nums[rand.Intn(len(nums))]
	}
	return string(result)
}

func getRandomElement(slice []string) string {
	return slice[rand.Intn(len(slice))]
}

func GenerateRandomUsers(userAmount int) {
	// Create a new file to write the user data
	file, err := os.Create("users.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Generate and write data for 20 random users
	for i := 1; i <= userAmount; i++ {
		id := generateRandomNumber(16)
		firstName := getRandomElement(firstNames)
		lastName := getRandomElement(lastNames)
		user := strings.ToLower(firstName+lastName)
		email := fmt.Sprintf("%s@%s", user, emailDomain)
		password := generateRandomPassword(32)
		phone := "9" + generateRandomNumber(8)
		zipcode := generateRandomNumber(8)
		address := getRandomElement(possibleAdressComplement) + " " + getRandomElement(possibleAdress)
		street_n := generateRandomNumber(3)
		
		line := fmt.Sprintf("id:%s\nuser:%s\nemail:%s\npassword:%s\nfirst_name:%s\nlast_name:%s\nphone:%s\naddress:%s\nstreet_n:%s\nzipcode:%s\n",
			id, user, email, password, firstName, lastName, phone, address, street_n, zipcode)

		_, err := file.WriteString(line)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
	fmt.Println("Random user data has been written to users.txt")
}
