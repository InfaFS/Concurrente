//a

chan llegada(id);
chan turno[N] (res);

Process Cliente[id:0..N-1]{
	Respuesta res;
	send llegada(id);
	receive turno[id] (res);
}

Process Empleado{
	int idAux;
	res;
	while (true){
		receive llegada(idAux);
		send turno[idAux](res)
	}
}

//b


chan llegada(id);
chan turno[N] (res);

Process Cliente[id:0..N-1]{
	Respuesta res;
	send llegada(id);
	receive turno[id] (res);
}

Process Empleado[id:0..1]{
	int idAux;
	res;
	while (true){
		receive llegada(idAux);
		send turno[idAux](res)
	}
}


//c 
chan siguiente[2] (int)
chan Pedido(int);
chan llegada(id);
chan turno[N] (res);

Process Cliente[id:0..N-1]{
	Respuesta res;
	send llegada(id);
	receive turno[id] (res);
}

Process Empleado[id:0..1]{
	int idAux;
	res;
	int resCoordinador;
	while(true){
		send Pedido(id);
		receive Siguiente[id] (resCoordinador);
		if (resCoordinador <> -1) send turno[resCoordinador] (res)
		else delay(15);
	}
}

Process Coordinador{

	while(true){
		receive Pedido(idE);
		if (empty(llegada)) Rep = -1;
		else receive llegada(id)
		send Siguiente[idE] (id)
	}

}