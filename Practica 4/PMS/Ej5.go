
Process Espectador[id:0..E-1]{
	Admin!solicitarAcceso(id);
	Admin?accesoConcedido();
	//usa maquina
	Admin!maquinaUsada();
}

Process Admin{
	bool libre = true;
	cola fila;
	int idE;
	do	Espectador[*]?solicitarAcceso(idE) -> if (not libre){
		fila.push(idE);
	} else {
		libre = false;
		Espectador[idE]!accesoConcedido();
	} 
	[] Espectador[*]?maquinaUsada() -> if (empty(fila)){
		libre = true;
	} else {
		fila.pop(idE);
		Espectador[idE]! accesoConcedido()
	}

	od
}