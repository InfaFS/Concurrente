Process Cliente[id:0..N-1]{
	send pedirCabina(id)
	send hayPedido()
	receive obtenerCabina[id] (idCabina)
	send enviarLlamada[idCabina] (id)
	receive llamadaRealizada[id] (llamadaHecha)

	send pedirComprobante(id)
	send hayPedido()
	receive enviarComprobante[id] (comprobante)

}

Process Empleado{
while(true) {
	receive hayPedido()
	if (!empty(obtenerCabina) and empty(pedirComprobante)){
		receive pedirCabina(idCliente)
		send obtenerCabina[idCliente] (elegirCabina())

	} else if (!empty(pedirComprobante)){
		receive pedirComprobante(idCliente)
		send enviarComprobante[idCliente] (Cobrar())
	}	
}
}

Process Cabina[id:0..9]{
while (true){
	receive enviarLlamada[id] (idCliente)
	//hacer llamada con el id del cliente
	send llamadaRealizada[idCliente] (true)
}
}