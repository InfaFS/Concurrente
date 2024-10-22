//No chequearia en el a por las otras condiciones, es decir no habria prioridad viendo si ni esta empty la de los que se van
//Por ejemplo en el a si podria usar el if ni determinisitico

//a

Process Cliente[id:0..N-1]{
	send pedirCabina(id)
	receive obtenerCabina[id] (idCabina)
	usarCabina(idCabina);

	send pedirComprobante(id,,idCabina)
	receive enviarComprobante[id] (comprobante)

}

Process Empleado{
	bool cabinasOcupadas[10] ([10] false);
	while(true) { //usar una cola si no quiero hacer busy waiting, aunque convendria mas hacerlo en este caso con buys waiting 

		if (!empty(obtenerCabina) and !(todasCabinasOcupadas())){
			receive pedirCabina(idCliente)
			elegirCabina(idCabina);
			cabinasOcupadas[idCabina] = true;
			send obtenerCabina[idCliente] (idCabina)
		} 
		[] (!empty(pedirComprobante)){
			receive pedirComprobante(idCliente,idCabina);
			cabinasOcupadas[idCabina] = false;
			send enviarComprobante[idCliente] (Cobrar())
		}
		fi	
	}
}


//b
//En este caso no uso un receiv pedido y hago busy waiting porque no puedo acumular
//a las personas que piden obtener una cabina y estan todas llenas
Process Cliente[id:0..N-1]{
	send pedirCabina(id)
	receive obtenerCabina[id] (idCabina)
	usarCabina(idCabina);

	send pedirComprobante(id,,idCabina)
	receive enviarComprobante[id] (comprobante)

}

Process Empleado{
	bool cabinasOcupadas[10] ([10] false);
	while(true) { //usar una cola si no quiero hacer busy waiting, aunque convendria mas hacerlo en este caso con buys waiting 
		//En este caso, al no encolarse los pedidos de alguna manera, si por ejemplo se pide una cabina
		//No hay alguien que libere entonces se va a perder el pedido!, asi que aca hay que hacer busy waiting si o si
		if (!empty(pedirCabina) and empty(pedirComprobante) and !(todasCabinasOcupadas())){
			receive pedirCabina(idCliente)
			elegirCabina(idCabina);
			cabinasOcupadas[idCabina] = true;
			send obtenerCabina[idCliente] (idCabina)

		} 
		[] (!empty(pedirComprobante)){
			receive pedirComprobante(idCliente,idCabina);
			cabinasOcupadas[idCabina] = false;
			send enviarComprobante[idCliente] (Cobrar())
		}	
	}
}

