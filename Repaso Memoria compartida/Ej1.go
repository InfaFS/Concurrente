
colaLlegada c;
sem espera[P] = ([P] 0);
sem mutex = 1;
bool libre = true;

Process persona[i:0..P-1]{
	Sube sube;
	int aux;

	P(mutex);
	if (libre){
		libre = false;
		V(mutex)
	} else {
		c.push(i)
		V(mutex)
		P(espera[i])
	}
	UsarTerminal(sube)
	P(mutex)
	if (c.isEmpty()){
		libre = true;
	} else {
		aux = c.pop()
		V(espera[aux])
	}
	V(mutex)
}

//b

sem espera[P] = ([P] 0);
sem mutex = 1;
colaPersonas c;
sem mutexColas;
colaTerminales colaT;
int terminalesDisp = 4;

Process persona[i:0..P-1]{
	Sube sube;
	int aux;

	P(mutex);
	if (terminalesDisp > 0){
		terminalesDisp--;
		V(mutex)
	} else {
		c.push(i)
		V(mutex)
		P(espera[i])
	}

	P(mutexColas)
	terminal = colaT.pop()
	V(mutexColas)
	UsarTerminal(sube,terminal)
	P(mutexColas)
	colaT.push(terminal)
	V(mutexColas)

	P(mutex)
	
	if (c.isEmpty()){
		terminalesDisp++;
	} else {
		aux = c.pop()
		V(espera[aux])
	}

	V(mutex)
}

