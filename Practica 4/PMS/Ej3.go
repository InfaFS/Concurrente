//a

Process alumnos[id:0..N-1]{
	text examen;
	int nota;
	examen = resolverExamen()
	admin!recibirExamen(examen,id);
	profesor?recibirNota(nota)
}

Process admin{
	cola Fila;
	do alumnos[*]?recibirExamen(examen) -> push(Fila,examen,id);
	[] !empty(Fila);profesor?pedirExamen() -> profesorEnviarExamen!(Fila.pop())
	od
}

Process profesor{
	int idAlumno;
	text examen;
	for i:0..C-1{
		admin!pedirExamen()
		admin?recibirExamen(examen,idAlumno)
		nota = corregir(examen)
		alumnos[idAlumno]! (nota)
	}
}

//b


Process alumnos[id:0..N-1]{
	text examen;
	int nota;
	examen = resolverExamen()
	admin!recibirExamen(examen,id);
	profesor[*]?recibirNota(nota)
}

Process admin{
	cola Fila;
	int idAlumno,idProf;
	int contador = N
	do alumnos[*]?recibirExamen(examen) -> push(Fila,examen,id);
	[] !empty(Fila) and contador > 0;profesor[*]?pedirExamen(idProf) -> contador--; profesorEnviarExamen[idProf]!(Fila.pop());
	[] contador == 0;profesor[*]?pedirExamen(idProf) -> profesorEnviarExamen[idProf]!("FINALIZADO",-1);
	od
}

Process profesor[id:0..P-1]{
	int idAlumno;
	text examen;
	while(examen <> "FINALIZADO"){
		admin!pedirExamen(id)
		admin?recibirExamen(examen,idAlumno)
		if(examen <> "FINALIZADO"){
			nota = corregir(examen)
			alumnos[idAlumno]! (nota)
		}
	}
}

//c


Process alumnos[id:0..N-1]{
	text examen;
	int nota;
	admin!llegue()
	//IMPORTANTE
	admin!comenzar() //Aca al ser sincronico va a esperar hasta que se reciba el mensaje, por ende se maximiza la concurrencia!
	examen = resolverExamen()
	admin!recibirExamen(examen,id);
	profesor?recibirNota(nota)
}

Process admin{
	cola Fila;
	int idAlumno,idProf;
	int contador = N

	for i = 0..N-1 -> alumnos[*]?llegue();

	for i = 0..N-1 -> alumnos[*]?comenzar(); //No mando una senial, sino simplemente recibo

	do alumnos[*]?recibirExamen(examen) -> push(Fila,examen,id);
	[] !empty(Fila) and contador > 0;profesor[*]?pedirExamen(idProf) -> contador--; profesorEnviarExamen[idProf]!(Fila.pop());
	[] contador == 0;profesor[*]?pedirExamen(idProf) -> profesorEnviarExamen[idProf]!("FINALIZADO");
	od
}

Process profesor[id:0..P-1]{
	int idAlumno;
	text examen;
	while(examen <> "FINALIZADO"){
		admin!pedirExamen(id)
		admin?recibirExamen(examen,idAlumno)
		if(examen <> "FINALIZADO"){
			nota = corregir(examen)
			alumnos[idAlumno]! (nota)
		}
	}
}