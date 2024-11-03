chan enviarPedido[3](texto,int);
chan respuesta[N] (texto);
chan enviarTramite(texto,int);
chan pedidoEmpleado(int);

Process persona[id:0..N-1]{
	tramite text;
	tramite = generarTramite();

	send enviarTramite(tramite,id);
	receive respuesta[id] (respuesta);

}

Process Admin {
	int contador = N;
	while(contador <> 0){
		receive pedidoEmpleado(idE);
		if (not empty(enviarTramite)){
			receive enviarTramite(tramite,idC);
			contador--;
		} else {
			tramite = "no hay tramite";
			idC = -1;
		}
		send enviarPedido[idE] (tramite,idC);
	}
	while (not empty(pedidoEmpleado)){
		receive pedidoEmpleado(idE);
		tramite = "FIN";
		idC = -1;
		send enviarPedido[id] (tramite,idC);
	}
}

Process Empleado[id:0..2]{
	int idC;
	Tramite tramite;
	while(tramite != "FIN"){
		send pedidoEmpleado(id);
		receive enviarPedido[id](tramite,idC);
		if (tramite == "no hay tramite"){
			leer(300);
		} else {
			tramiteResuelto = resolverTramite(tramite);
			send respuesta[idC] (tramiteResuelto);
		}
	}
}