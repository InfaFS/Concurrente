//PMS

Process Empleado[id:0..99]{
	while(true){
		//procesar documento
		Admin!pedidoEmpleado(documento,id);
		impresora[*]?respuestaDocumento(res);
	}
}

Process Admin{
	documento texto;
	id int;
	idI int;
	while(true){
		do Empleado[*]?pedidoEmpleado(documento,id) -> colaDocumentos.push(documento,id);
		[] (!empty(colaDocumentos)); impresora[*]?PedidoImpresora(idI) -> 
			impresora[idI]!enviarDocumento(colaDocumentos.pop());
		od
	}
}

Process impresora[id:0..4]{
	idE int;
	documento texto;
	while(true){
		Admin!PedidoImpresora(id);
		Admin?enviarDocumento(idE,documento);
		res = imprimir(documento);
		Empleado[idE]!respuestaDocumento(res);
	}
}