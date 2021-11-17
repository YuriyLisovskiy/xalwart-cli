package core

import (
	"crypto/rand"
	"math/big"
)

func Generate(component Component, overwrite bool) error {
	return component.TemplateBox().Walk(func(template Template) error {
		return template.Render(component)
	}, component, overwrite)
}

func RandomString(length uint) (string, error) {
	n := int(length)
	const letters = "abcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*(-_=+)"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}

		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}
