/*



*/

Process primerEmpleado{
	text muestra;
	while(true){
		muestra = prepararMuestra()
		PrimerAdmin!enviarMuestra(muestra)
	}
}

Process PrimerAdmin{
	cola Buffer;
	text muestra;

	do primerEmpleado?enviarMuestra(muestra) -> push(Buffer,muestra);
	[] !empty(Buffer); segundoEmpleado?pedirMuestra() -> segundoEmpleado!enviarMuestra(pop(Buffer));
	od
}

Process segundoEmpleado{
	text muestra;
	text set;
	while(true){
		PrimerAdmin!pedirMuestra()
		PrimerAdmin?enviarMuestra(muestra)
		set = armarSet(muestra)
		tercerEmpleado!enviarSet(set)
		tercerEmpleado?respuesta(resultado)
		archivar(resultado)

	}
}


Process tercerEmpleado{
	text set,resultado
	while(true){
		segundoEmpleado?enviarSet(set)
		resultado = analizar(set)
		segundoEmpleado!respuesta(resultado)
	}
}