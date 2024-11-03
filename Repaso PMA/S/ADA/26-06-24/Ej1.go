
Process Persona[id:0..N-1]{
	solicitud string;
	res string;
	solicitud = generarSolicitud();
	send pedidoTramite(id,solicitud);
	receive recibirRespuesta[id] (res);
}

Process Empleado{
	contador = N;
	solicitud string;
	idC int;
	while(contador <> 0){
		if (not empty(pedidoTramite)){
			receive pedidoTramite (idC,solicitud);
			contador --;
			res = resolverTramite(solicitud);
			send recibirRespuesta[idC] (res);
		} else {
			leer(300);
		}
	}
}