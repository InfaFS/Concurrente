int accionControladores[14] = ([14] 0)
sem esperaControladores[14] = ([14] 0)

colaControladores cola;
sem mutexCola = 1;

sem modulos = 0;

Process Controladores[id=0..14]{
	int temperatura;
	while(true){
		temperatura = medir()
		P(mutexCola)
		cola.push(id)
		V(mutexCola)

		V(modulos)
		P(esperaControladores[id])
		actuar(accionControladores[id])
		//pasa tiempo
	}

}

Process Modulos[id=0..1]{
	int idControlador;
	int accion;
	while(true){
		P(modulos)

		P(mutexCola)
		idControlador = cola.pop()
		V(mutexCola)
		accion = determinar()
		
		accionControladores[idControlador] = accion
		V(esperaControladores[idControlador])
	}

}