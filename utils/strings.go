package utils

import (
	"github.com/maolchen/krm-backend/constants"
	"golang.org/x/crypto/bcrypt"
)

func Random(n int, chars string) string {
	return Random(n, chars)
}

func RandomLetters(n int) string {
	return Random(n, constants.LETTERS)
}

func RandomNumberic(n int) string {
	return Random(n, constants.NUMBERS)
}

func RandomAscii(n int) string {
	return Random(n, constants.ASCII)
}

// 明文加密
func MakeHashPassword(RawPassword string) (HashPass string, err error) {
	pwd := []byte(RawPassword)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return
	}
	HashPass = string(hash)
	return
}

// 密码比对
func CompareHashAndPassword(hashedPassword, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

// 处理ClusterKubeconfig日志打印
func PrintClusterKubeconfig(kubeconfig map[string][]byte) map[string]string {
	result := make(map[string]string, len(kubeconfig))
	for k, v := range kubeconfig {
		result[k] = string(v)
	}
	return result
}
