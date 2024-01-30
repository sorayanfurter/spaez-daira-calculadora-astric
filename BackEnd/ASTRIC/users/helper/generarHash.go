package helper

import (
	"ASTRIC/BackEnd/ASTRIC/helper/env"
	md "crypto/sha256"
	"encoding/hex"
	"time"
)

// GetHash Crea un hash sha256 con la clave secreta para validar la secion del toquen
func GetHash() string {
	str := env.GetEnvGeneral().ClaveSecreta + time.Now().String()
	hasher := md.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}
