
chan reportes(texto);
chan Siguiente[3](texto);
chan pedido(int);

Process Cliente[id:0..P-1]{
	while(true){
		//trabaja
		send reportes(reporteGenerado)
	}
}

Process Admin{
	int idP;
	texto res;
	while(true){
		receive pedido(idP);
		if (!empty(reportes)){
			receive reportes(res);
		} else {
			res = "VACIO";
		}
		send Siguiente[idP](res);
	}
}

Process Empleado[id:0..2]{
	text res;

	while (true){
		send pedido(id)
		receive Siguiente[id] (res)
		if (res == "VACIO"){
			Programar();
		} else {
			ResolverErorr(res);
		}
	}
}