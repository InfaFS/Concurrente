Process Estudiante[id:0..E-1]{
	Admin!PedidoHorno(id);
	Comedor?AccesoHorno();
	//Usa el horno
	Comedor!TerminadoAcceso();
}

Process Admin{
	idE int;
	cola ColaPedidos;
	while(true){
		do Estudiante[*]?PedidoHorno(idE) -> ColaPedidos.push(idE);
		[] (!Empty(ColaPedidos)); Comedor?PedidoAcceso() -> Comedor!RecibirId(ColaPedidos.pop());
		od
	}
}

Process Comedor{
	idE int;
	while(true){
		Admin!PedidoAcceso();
		Admin?RecibirId(idE);

		Estudiante[idE]!AccesoHorno();
		Estudiante?TerminadoAcceso():
	}
}

//Solucion catedra

Process Estudiante[id:0..E-1]{
	Admin!PedidoHorno(id);
	Admin?AccesoHorno();
	//Usa el horno
	Admin!TerminadoAcceso();
}

Process Admin{
	idE int;
	cola ColaPedidos;
	while(true){
		If (libre = true); Estudiante[*]?PedidoHorno(idE) ->
			libre = false;
			Estudiante[idE]!AccesoHorno();
		[] (libre = false);Estudiante[*]?PedidoHorno(idE) ->
			ColaPedidos.push(idE);
		[] Estudiante[*]?TerminadoAcceso();
			if (empty(ColaPedidos)){
				libre = true;
			} else {
				Estudiante[ColaPedidos.pop()]!AccesoHorno();
			}
		fi
	}
}