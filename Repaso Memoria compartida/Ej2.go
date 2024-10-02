colaTransacciones colaT;
int vector[10] = ([10] 0);
sem mutexVector[10] = ([10] 0)
int ultimoWorker = 7;
sem mutexUltimoWorker = 1;
sem mutex = 1;


process Worker[0..6]{
	P(mutex)
	while(!colaT.empty()){
		transaccion = colaT.pop()
		V(mutex)
		numero = validar(transaccion)
		
		P(mutexVector[numero])
		vector[numero] = vector[numero] + 1;	
		V(mutexVector[numero])

		P(mutex)
	}
	V(mutex)
	P(mutexUltimoWorker)
	ultimoWorker = ultimoWorker - 1;
	if (ultimoWorker != 0){
		V(ultimoWorker)
	} else {
		for j:0..9{
			print(vector[j])
		}
		V(mutexUltimoWorker)
	}
	
}