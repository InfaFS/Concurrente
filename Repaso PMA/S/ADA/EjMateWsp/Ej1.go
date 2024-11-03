
Process Cliente[id:0..N-1]{
	pasaje text;
	respuesta text;
	solicitud = TipoSolicitud();
	if (solicitud == "Compra"){
		send PedidoDevolucion(id,pasaje);
		receive respuestaDevolucion[id] (respuesta);
	} else {
		send PedidoCompra(id);
		receive respuestaCompra[id] (pasaje);
	}

}


Process Admin{
	idS int;
	while(true){
		receive PedidoServer(idS);
		if (not empty(PedidoDevolucion) and (empty (PedidoCompra))){
			receive PedidoDevolucion(idC,pasaje);
			send recibirPedido (idC,pasaje);
		} else if (not empty(PedidoCompra)){
			receive PedidoCompra(idC);
			send recibirPedido(idC,"COMPRA");
		} else if (empty(PedidoCompra) and empty(PedidoDevolucion)){
			send recibirPedido(-1,"NADA");
		}
	}
}

Process Server[id:0..2]{
	idC int;
	pasaje texto;
	while(true){
		send PedidoServer(idS);
		receive recibirPedido(idC,pasaje);
		if (pasaje <> "NADA"){
			if (pasaje == "COMPRA"){
				pasaje = generarPasaje();
				send respuestaCompra[idC] (pasaje);
			} else {
				respuesta = generarRespuesta();
				send respuestaDevolucion[idC] (respuesta);
			}
		}
	}
}