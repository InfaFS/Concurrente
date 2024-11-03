
Process Cliente[id:0..N-1]{
	Tramite tramite;
	Tramite respuesta;
	tramite = generarTramite();
	Admin!enviarTramite(tramite,idC);
	Empleado[*]?recibirRespuesta(res);
}

Process Admin{
	Tramite tramite;
	int contador = N;
	cola colaTramites;
	while(contador <> 0){
		if (contador <> 0); Cliente[*]?enviarTramite(tramite); colaTramites.push(tramite,idC);
		[] (not empty(colaTramites)); Empleado[*]?pedidoEmpleado(idE);{
			Empleado[idE]!recibirTramite(colaTramites.pop());
			contador--;
		} 
		fi
	}
	
	for i:0..2{
		Empleado[*]?pedidoEmpleado(idE);
		Empleado[idE]!recibirTramite("FIN",-1);
	}
}

Process Empleado[id:0..2]{
	Tramite tramite;
	int idC;

	while (tramite <> "FIN"){
		Admin!pedidoEmpleado(id);
		Admin?recibirTramite(Tramite,idC);
		if (tramite <> "FIN"){
			tramiteRes = resolverTramite(Tramite);
			Cliente[idC]!recibirRespuesta(tramiteRes);
		}
	}	
}