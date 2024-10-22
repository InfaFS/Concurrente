//a
Process persona[id:0..P-1]{
	empleado!pedirAcceso(id);
	empleado?accesoConcedido();
	//usar simulador
	empleado!simulacionCompletada();
}


Process empleado{
	int idP;
	while(true){
		persona[*]?pedirAcceso(idP);
		persona[idP]!accesoConcedido();

		persona?simulacionCompletada();
	}
}


//b
Process persona[id:0..P-1]{
	admin!pedirAcceso(id);
	empleado?accesoConcedido();
	//usar simulador
	empleado!simulacionCompletada();
}

Process admin{
	int idP;
	cola Fila;
	do persona[*]?pedirAcceso(idP) -> Fila.push(idP);
	[] !empty(fila);empleado?pedirSiguiente() -> empleado!pedirAcceso(Fila.pop())
	od
}

Process empleado{
	int idP;
	while(true){
		admin!pedirSiguiente();
		admin?pedirAcceso(idP);
		persona[idP]!accesoConcedido();

		persona?simulacionCompletada();
	}
}


//b Version agus


Process Persona [id: 0..P-1] {
    Empleado!solicitarPaso(id);
    Empleado?pasar();
    usarSimulador();
    Empleado!salir();
}

Process Empleado {
    cola cola;
    int idAux;
    bool libre = true;

    do Persona[*]?solicitarPaso(idAux) ->
		if (!libre) {
			cola.push(idAux);
		} else {
			libre = false;
			Persona[idAux]!pasar();
		}
    [] Persona[*]?salir() ->
		if (empty(cola)) {
			libre = true;
		} else {
			idAux = cola.pop();
			Persona[idAux]!pasar();
		}
    od
}


