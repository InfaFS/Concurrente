
Process Turista[id:0..19]{
	empleado!llegada(); //El turista se queda esperando a que le reciban el llegada
	empleado!terminaCharla(); //Luego espera a que le reciban que ya termino la charla

	admin!pedidoTirolesa(id); //Le pide al admin tirarse 
	guia?concederPermiso(); //El guia le concede el permiso

	UsarTirolesa(); //La usa

	guia!tirolesaCompletada(); //Le avisa que ya la completo
 
}

Process Empleado{
	for i=0..19{ //Recibe cualquiera de los llegada
		Turista[*]?llegada();
	}
	DarCharla();
	for i=0..19{ //Recibe el terminaCharla de cualquiera para que ya puedan empezar a pedirel al guia para tirarse
		Turista[*]?terminaCharla();
	}

}

Process Admin{
	int contadorP = 20;
	Cola cola;
	do (contador > 20); Turista[*]?pedidoTirolesa(idT) -> cola.push(idT) //Recibe pedidos de tirolesa, hasta 20 veces
		contadorP--;
	[] !empty(cola); Guia?siguienteTurista() -> Guia!recibirTurista(cola.pop()); //Recibe pedidos del admin para mandarle el turista
	od
}

Process Guia{
	int idT;
	for i=0..19{
		Admin!siguienteTurista(); //Pide el siguiente turista
		Admin?recibirTurista(idT); //Espera que le den el id del turista

		Turista[idT]!concederPermiso(); //Le concede el permiso al turista

		Turista[idT]?tirolesaCompletada(); //Se queda bloqueado hasta que le manden que termino de tirarse el turist

	}
}