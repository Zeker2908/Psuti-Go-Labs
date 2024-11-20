package tenth

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"hash"
	"io"
	"os"
)

func First() {
	for {
		fmt.Println("\nХеширование данных")
		fmt.Println("===================")
		fmt.Println("Выберите действие:")
		fmt.Println("1. Вычислить хэш строки")
		fmt.Println("2. Проверить целостность данных")
		fmt.Println("3. Выйти из программы")
		fmt.Print("Введите номер действия: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			computeHash()
		case 2:
			verifyIntegrity()
		case 3:
			fmt.Println("Выход из программы.")
			return
		default:
			fmt.Println("Неверный выбор. Попробуйте снова.")
		}
	}
}

func computeHash() {
	fmt.Print("Введите строку: ")
	var input string
	fmt.Scanln(&input)

	fmt.Println("Выберите алгоритм:")
	fmt.Println("1. MD5")
	fmt.Println("2. SHA-256")
	fmt.Println("3. SHA-512")
	fmt.Print("Введите номер алгоритма: ")

	var algoChoice int
	fmt.Scanln(&algoChoice)

	var hasher hash.Hash
	switch algoChoice {
	case 1:
		hasher = md5.New()
	case 2:
		hasher = sha256.New()
	case 3:
		hasher = sha512.New()
	default:
		fmt.Println("Неверный выбор алгоритма")
		return
	}

	hasher.Write([]byte(input))
	hashValue := hasher.Sum(nil)
	fmt.Printf("Хэш (%T): %s\n", hasher, hex.EncodeToString(hashValue))
}

func verifyIntegrity() {
	fmt.Print("Введите строку: ")
	var input string
	fmt.Scanln(&input)

	fmt.Print("Введите хэш: ")
	var hashInput string
	fmt.Scanln(&hashInput)

	fmt.Println("Выберите алгоритм:")
	fmt.Println("1. MD5")
	fmt.Println("2. SHA-256")
	fmt.Println("3. SHA-512")
	fmt.Print("Введите номер алгоритма: ")

	var algoChoice int
	fmt.Scanln(&algoChoice)

	var hasher hash.Hash
	switch algoChoice {
	case 1:
		hasher = md5.New()
	case 2:
		hasher = sha256.New()
	case 3:
		hasher = sha512.New()
	default:
		fmt.Println("Неверный выбор алгоритма")
		return
	}

	hasher.Write([]byte(input))
	computedHash := hex.EncodeToString(hasher.Sum(nil))

	if computedHash == hashInput {
		fmt.Println("Целостность данных подтверждена.")
	} else {
		fmt.Println("Целостность данных нарушена!")
		fmt.Printf("Ожидалось: %s\nПолучено: %s\n", hashInput, computedHash)
	}
}

// 2.	Симметричное шифрование:
// •	Реализуйте программу, шифрующую переданные данные с помощью алгоритма AES.
// •	Пользователь должен указать строку и секретный ключ.
// •	Программа должна зашифровать строку и предоставить возможность расшифровать её при вводе того же ключа.
func Second() {
	for {
		fmt.Println("\nСимметричное шифрование (AES)")
		fmt.Println("============================")
		fmt.Println("1. Зашифровать строку")
		fmt.Println("2. Расшифровать строку")
		fmt.Println("3. Выйти из программы")
		fmt.Print("Введите номер действия: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			encryptString()
		case 2:
			decryptString()
		case 3:
			fmt.Println("Выход из программы.")
			return
		default:
			fmt.Println("Неверный выбор. Попробуйте снова.")
		}
	}
}

func encryptString() {
	fmt.Print("Введите строку для шифрования: ")
	var plaintext string
	fmt.Scanln(&plaintext)

	fmt.Print("Введите секретный ключ (16, 24 или 32 символа): ")
	var key string
	fmt.Scanln(&key)

	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		fmt.Println("Ключ должен содержать 16, 24 или 32 символа.")
		return
	}

	ciphertext, err := encrypt([]byte(plaintext), []byte(key))
	if err != nil {
		fmt.Printf("Ошибка шифрования: %v\n", err)
		return
	}

	fmt.Println("Зашифрованная строка:", ciphertext)
}

func decryptString() {
	fmt.Print("Введите строку для расшифровки: ")
	var ciphertext string
	fmt.Scanln(&ciphertext)

	fmt.Print("Введите секретный ключ (16, 24 или 32 символа): ")
	var key string
	fmt.Scanln(&key)

	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		fmt.Println("Ключ должен содержать 16, 24 или 32 символа.")
		return
	}

	plaintext, err := decrypt(ciphertext, []byte(key))
	if err != nil {
		fmt.Printf("Ошибка расшифровки: %v\n", err)
		return
	}

	fmt.Println("Расшифрованная строка:", plaintext)
}

func encrypt(plaintext, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]

	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func decrypt(ciphertext string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	rawCiphertext, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	if len(rawCiphertext) < aes.BlockSize {
		return "", fmt.Errorf("шифртекст слишком короткий")
	}

	iv := rawCiphertext[:aes.BlockSize]
	rawCiphertext = rawCiphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(rawCiphertext, rawCiphertext)

	return string(rawCiphertext), nil
}

//3.	Асимметричное шифрование и цифровая подпись:
//•	Создайте пару ключей (открытый и закрытый) и сохраните их в файл.
//•	Реализуйте программу, которая подписывает сообщение с помощью закрытого ключа и проверяет подпись с использованием открытого ключа.
//•	Продемонстрируйте пример передачи подписанных сообщений между двумя сторонами.

const (
	privateKeyFile = "private_key.pem"
	publicKeyFile  = "public_key.pem"
)

func Third() {
	for {
		fmt.Println("\nАсимметричное шифрование и цифровая подпись")
		fmt.Println("=========================================")
		fmt.Println("1. Создать пару ключей")
		fmt.Println("2. Подписать сообщение")
		fmt.Println("3. Проверить подпись")
		fmt.Println("4. Выйти из программы")
		fmt.Print("Введите номер действия: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			generateKeys()
		case 2:
			signMessage()
		case 3:
			verifySignature()
		case 4:
			fmt.Println("Выход из программы.")
			return
		default:
			fmt.Println("Неверный выбор. Попробуйте снова.")
		}
	}
}

// Создание пары ключей и сохранение в файлы
func generateKeys() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Printf("Ошибка генерации ключей: %v\n", err)
		return
	}

	// Сохранение закрытого ключа
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	})
	err = os.WriteFile(privateKeyFile, privateKeyPEM, 0600)
	if err != nil {
		fmt.Printf("Ошибка сохранения закрытого ключа: %v\n", err)
		return
	}
	fmt.Printf("Закрытый ключ сохранён в файл: %s\n", privateKeyFile)

	// Сохранение открытого ключа
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		fmt.Printf("Ошибка кодирования открытого ключа: %v\n", err)
		return
	}
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	})
	err = os.WriteFile(publicKeyFile, publicKeyPEM, 0644)
	if err != nil {
		fmt.Printf("Ошибка сохранения открытого ключа: %v\n", err)
		return
	}
	fmt.Printf("Открытый ключ сохранён в файл: %s\n", publicKeyFile)
}

// Подпись сообщения
func signMessage() {
	privateKey, err := loadPrivateKey(privateKeyFile)
	if err != nil {
		fmt.Printf("Ошибка загрузки закрытого ключа: %v\n", err)
		return
	}

	fmt.Print("Введите сообщение для подписи: ")
	var message string
	fmt.Scanln(&message)

	hash := sha256.Sum256([]byte(message))
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
	if err != nil {
		fmt.Printf("Ошибка создания подписи: %v\n", err)
		return
	}

	signatureBase64 := base64.StdEncoding.EncodeToString(signature)
	fmt.Println("Подпись:", signatureBase64)
}

// Проверка подписи
func verifySignature() {
	publicKey, err := loadPublicKey(publicKeyFile)
	if err != nil {
		fmt.Printf("Ошибка загрузки открытого ключа: %v\n", err)
		return
	}

	fmt.Print("Введите сообщение: ")
	var message string
	fmt.Scanln(&message)

	fmt.Print("Введите подпись (Base64): ")
	var signatureBase64 string
	fmt.Scanln(&signatureBase64)

	signature, err := base64.StdEncoding.DecodeString(signatureBase64)
	if err != nil {
		fmt.Printf("Ошибка декодирования подписи: %v\n", err)
		return
	}

	hash := sha256.Sum256([]byte(message))
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash[:], signature)
	if err != nil {
		fmt.Println("Подпись недействительна.")
	} else {
		fmt.Println("Подпись успешно подтверждена.")
	}
}

// Загрузка закрытого ключа из файла
func loadPrivateKey(filename string) (*rsa.PrivateKey, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(data)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("неверный формат закрытого ключа")
	}

	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

// Загрузка открытого ключа из файла
func loadPublicKey(filename string) (*rsa.PublicKey, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(data)
	if block == nil || block.Type != "RSA PUBLIC KEY" {
		return nil, fmt.Errorf("неверный формат открытого ключа")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	publicKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("неверный тип ключа")
	}

	return publicKey, nil
}
