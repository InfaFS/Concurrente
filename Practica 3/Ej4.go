
Monitor Puente{
	peso_actual = 0;
	cond cola;

	Procedure pasar(peso: in int){
		while ((peso_actual + peso) > 50000){
			wait(cola);
		}
		peso_actual += peso;
	}
	
	Procedure termino(peso: in int){
		peso_actual -= peso;
		signal(cola);
	}
}

Process Autos[id:0..N-1]{
	Puente.pasar(peso);
	//pasa puente
	Puente.terminar(peso);
}