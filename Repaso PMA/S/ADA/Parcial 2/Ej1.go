//Es improtante que en sincronico y cuando no queramos hacer que sea deterministico la cosa, haya sistema de pedidos
//Com oes el caso del organizador o de los supervisores

Process Organizador{
	desafio texto;
	for i = 0..N-1{
		desafio = elegirDesafio();
		Competidor[*]?pedidoOrganizador(idC)
		Competidor[idC]!recibirDesafio(desafio);
	}	
}

Process Competidor[id:0..N-1]{
	desafio texto;
	resolucion texto;
	correcto bool = false;
	Organizador!pedidoOrganizador(id);
	Organizador?recibirDesafio(desafio);
	//realiza el desafio
	while(!correcto){
		Admin!EnviarDesafio(desafio,id);
		Supervisor[*]?RecibirCorreccion(correcto);
	}

	
}

Process Admin{
	while(true){
		if Competidor[*]?EnviarDesafio(desafio,id) -> cola.push(desafio,id);
		[] (!empty(cola));Supervisor[*]?PedirDesafio(idS) -> Supervisor[idS]!EnviarDesafio(cola.pop());
		fi
	}
}

Process Supervisor[id:0..4]{
	correcion bool;
	while(true){
		Admin!PedirDesafio(id);
		Admin?EnviarDesafio(desafio,idC);
		//corregir el desafrio
		Competidor[idC]!RecibirCorreccion(correcion);
	}
}
