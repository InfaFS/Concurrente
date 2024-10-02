

Monitor Aula{
	int cant = 0;
	cond esperaAlumnos[50];
	int vectorIdAlumnos[50] = ([50] -1);
	int nota = 25;
	int vectorGruposPuntaje[25] = ([25] 0);
	cond vectorGruposEspera[25];
	Cola ColaTareas;

	Procedure llegada(idAlumno: in int ,numGrupo: out int){
		cant++
		if (cant < 50) { wait (esperaAlumnos[idAlumno]) }
		else {
			signal(esperaProfesor) 
			wait(esperaAlumnos[idAlumno])
		}
		numGrupo = vectorIdAlumnos[idAlumno]
	}
	
	Procedure iniciar(){
		if (cant < 50) { wait (esperaProfesor) }
		for i:0..49{
			vectorIdAlumnos[i] = AsignarNumeroGrupo()
			signal(esperaAlumnos[i])
		}
	}

	Procedure terminar(nota: out int,grupoId: in int){
		ColaTareas.push(grupoId);
		signal(esperaProfesor)	
		wait(vectorGruposEspera[grupoId])
		nota = vectorGruposPuntaje[grupoId]
	}

	Procedure esperarTarea(){
		wait(esperaProfesor)
		GrupoId = ColaTareas.pop()
		vectorGruposPuntaje[GrupoId]++
		if (vectorGruposPuntaje[GrupoId] == 2){
			vectorGruposPuntaje[GrupoId] = nota
			nota --
			signal_all(vectorGruposEspera[grupoId])
		}
	}
}





Process Alumno[id:0..49]{
	Aula.llegada(numGrupo);
	//REALIZA lA TAREA
	Aula.terminar();

}

Process Profesor{
	Aula.iniciar();
	for i:0..49{
		Aula.esperarTarea();
	}

}