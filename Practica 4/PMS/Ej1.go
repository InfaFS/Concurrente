//PREGUNTAR COMO FUNCIONA LO DEL ORDEN Y SI ESTAN BIEN HECHOS
//a, b -> para maximizar la concurrencia
//no debemos dejar al examinizador esperando a que le acepte el tesster el reporte
//por ende creamos un admin


Process admin{
	cola Buffer;
	text reporte;

	do examinizador[*]?reporte(direccion_sitio) -> push(Buffer,direccion_sitio)
	not empty(Buffer); tester?pedido() -> tester!reporte (pop(Buffer))
	od 
}

Process examinizador[id:0..R-1]{
	while(true){
		//buscando sitio infectado
		admin!reporte(direccion_sitio)
	}
}

Process tester{
	while(true){
		admin!pedido()
		admin?reporte(direccion_sitio)
		resolver(direccion_sitio)
	}
}

//c

Process examinizador[id:0..R-1]{
	while(true){
		//buscando sitio infectado
		admin!reporte(direccion_sitio)
	}
}

Process tester{
	while(true){
		admin!pedido()
		admin?reporte(direccion_sitio)
		resolver(direccion_sitio)
	}
}

Process admin{
	do examinizador[*]?reporte(direccion_sitio) -> push(Fila,(direccion_sitio))
	   not empty(Fila);tester?pedido() -> pop(Fila,(direccion_sitio)) tesster!reporte(direccion_sitio)
	od
}
