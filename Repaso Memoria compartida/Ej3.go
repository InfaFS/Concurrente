
int capacidad = 100;
colaPersonas cola;
bool libre = true;
sem mutex = 1;
sem usuarios[U] = ([U] 0)
sem repositor = 0;
sem repuesto = 0;

Process Usuario[id:0..U-1]{
	P(mutex)
	if (libre){
		libre = false;
		V(mutex)
	} else {
		colaPersonas.push(id)
		V(mutex)
		P(usuarios[id])
	}

	//tomar botella
	capacidad --
	if (capacidad == 0){
		V(repositor)
		P(repuesto)
	}
	P(mutex)
	if(cola.empty()){
		libre = true
	} else {
		aux = cola.pop()
		V(usuarios[id])
	} 
	V(mutex)

}

Process repositor{
	while(true){
		P(repositor)
		capacidad = 100;
		V(repuesto)
	}
}