//PMA

Process Empleado[id:0..99]{
	while(true){
		//procesar documento
		send pedidoEmpleado(documento,id);
		receive respuestaDocumento(res);
	}
}

Process impresora[id:0..4]{
	while(true){
		receive pedidoEmpleado(documento,id);
		res = imprimir(documento);
		send respuestaDocumento[id] (res);
	}
}