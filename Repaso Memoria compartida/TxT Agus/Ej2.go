
Monitor ComprasYVentas{
	cond cv;
	int contador = 0
	cond Boletero
	int respuestaBoletero

	Procedure Llegada(numeroEntrada: out int){
		contador++
		signal(Boletero)
		wait(cv)
		numeroEntrada = respuestaBoletero
	}

	Procedure espera(){
		if (contador == 0){wait(Boletero)}
		contador --
	}

	Procedure darRespuesta(numeroEntrada: in int){
		respuestaBoletero = numeroEntrada
		signal(cv)

	}

}


Process Personas[id=0..P-1]{
	int entrada;
	ComprasYVentas.Llegada(entrada)

}


Process Boletero{
	int cantidadEntradas = E
	int numeroEntrada;
	while(true){
		ComprasYVentas.espera()
		if (cantidadEntradas > 0){
			cantidadEntradas--
			numeroEntrada = vender()
			ComprasYVentas.darRespuesta(numeroEntrada)
		} else {
			ComprasYVentas.darRespuesta(-1)
		}
	}

}