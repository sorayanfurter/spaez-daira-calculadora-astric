package messages

// GetMessagesValid devuelve los mensajes pre establecidos.
func GetMessagesValid() map[string]string {
	messages := make(map[string]string)

	//Declaracion de totos los mensajes que se quierena incorporar a las validaciones de la API
	messages["IMPORTANTE"] = "Los nombre de las propiedades deben tener todos sus cararcteres en minusculas"

	return messages
}
