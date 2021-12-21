package services

import (
	"log"
	"reflect"
	"regexp"
)

func isValidInput(data string) bool {
	valid, _ := regexp.MatchString(`^[a-zA-Z0-9À-ÿ/./,/(/)/!/¡/¿/?/$/#/^/@/%/*-://\s]*$`, data)
	return valid
}

func isValidEmail(data string) bool {
	email, _ := regexp.MatchString(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`, data)
	return email
}

func isValidUser(data string) bool {
	return len(data) >= 3 && len(data) <= 20 && isValidInput(data)
}

func isValidPassword(data string) bool {
	return len(data) >= 8 && isValidInput(data)
}

func ValidateData(data interface{}) bool {
	dataValue := reflect.ValueOf(data)
	dataType := dataValue.Type()

	for i := 0; i < dataValue.NumField(); i++ {
		field := dataType.Field(i)
		value := dataValue.Field(i).String()

		if field.Name == "Email" {
			if !isValidEmail(value) {
				log.Println("Error email", value)
				return false
			}
		} else if field.Name == "Username" {
			if !isValidUser(value) {
				log.Println("Error user", value)
				return false
			}
		} else if field.Name == "Password" {
			if !isValidPassword(value) {
				log.Println("Error pass", value)
				return false
			}
		} else if field.Type.String() == "string" {
			if !isValidInput(value) {
				log.Println("Error input", value)
				return false
			}
		}
	}

	return true
}
