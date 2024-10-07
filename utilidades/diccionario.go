package utilidades

// Struct que implementa el diccionario sobre un mapa de Go
type Dictionary[K comparable, V any] struct {
	dict map[K]V
}

// Funcion 'NewDictionar' crea una nueva instancia del struct Dictionary
func NewDictionary[K comparable, V any]() *Dictionary[K, V] {
	return &Dictionary[K, V]{dict: make(map[K]V)}
}

// Funcion 'Put' que inserta el par (Key, Value) en el Diccionario. Se verifica que si la clave existe, reemplaza el par existente. Si no existe, el par es agregado al Diccionario.
func (d *Dictionary[K, V]) Put(key K, val V) {
	d.dict[key] = val
}

// Funcion 'Remove' que elimina una clave del diccionario si existe y devuelve 'true', caso contrario 'false'
func (d *Dictionary[K, V]) Remove(key K) bool {
	// Se verifica si el valor asociado a la clave existe o no
	if _, found := d.dict[key]; found {
		delete(d.dict, key)
		return true
	}
	return false
}

// Funcion 'Contains' que verifica si una clave existe dentro del diccionario. Devuelve 'true' o 'false' dependiendo el caso
func (d *Dictionary[K, V]) Contains(key K) bool {
	if _, found := d.dict[key]; found {
		return true
	}
	return false
}

// Funcion 'Get' obtiene el valor asociado a una clave
func (d *Dictionary[K, V]) Get(key K) V {
	val := d.dict[key]
	return val
}

// Funcion 'Size' devuelve el tamaño del diccionario (el número de elementos)
func (d *Dictionary[K, V]) Size() int {
	return len(d.dict)
}

// Fucnion 'Keys' devuelve un slice con todas las claves del diccionario
func (d *Dictionary[K, V]) Keys() []K {
	// Se decalara una variable que almacena un slice con las claves
	keys := make([]K, 0, len(d.dict))
	// Ciclo que itera sobre el diccionario y asigna cada clave al slice
	for key := range d.dict {
		keys = append(keys, key)
	}
	return keys
}

// Funcion 'Values' devuelve un slice con todos los valores del diccionario
func (d *Dictionary[K, V]) Values() []V {
	// Se decalara una variable que almacena un slice con los valores de las claves
	values := make([]V, 0, len(d.dict))
	// Ciclo que itera sobre el diccionario y asigna cada valor al slice
	for _, value := range d.dict {
		values = append(values, value)
	}
	return values
}
