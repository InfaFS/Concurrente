
Monitor Oficina {
	txt resultados[N-1]
	cond cv
	int esperando = 0
	cond Empleado
	cola colaSolicitudes;

	Procedure llegada(solicitud: in txt,resultadoTramite: out txt,idPersona: in int){
		esperando++
		colaSolicitudes.push(solicitud,idPersona)
		wait(cv)
		resultadoTramite = resultados[idPersona]
	}

	Procedure chequearPersonas(libre: out boolean){
		if (esperando == 0){
			libre = true
		} else {
			libre = false
		}

	}

	Procedure atenderPersona(solicitud: out txt,idPersona: out int){
		solicitud,id = colaSolicitudes.pop()
		esperando --
	}

	Procedure entregarResultado(resultado: in txt,idPersona: in int){
		resultados[idPersona] = resultado
		signal(cv)
	}

}


Process Personas[id=0..N-1]{
	txt resultadoTramite,solicitud
	Oficina.llegada(solicitud,resultadoTramite)

}

Process Empleado{
	boolean libre;
	txt solicitud,resultado
	int id;
	for id=0..N-1{
		Oficina.chequearPersonas(libre)
		while(libre = true){
			delay(600000)
			Oficina.chequearPersonas(libre)
		}
		Oficina.atenderPersona(solicitud,id)
		resultado = resolverTramite(solicitud)
		Oficina.entregarResultado(resultado,id)

	}

}