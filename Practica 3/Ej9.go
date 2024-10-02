

Monitor Parcial{
	cond espera;
	esperando = 0;
	cond preceptor;
	txt parcial;
	Cola colaParciales;

	int notasAlumnos[50] = ([50] -1)

	cond alumnosEsperan[50];

	int contadorParciales;
	
	cond profe;

	txt parcialMandado;

	Procedure llegada(copia_parcial: out txt){
		esperando++;
		if(esperando < 45){ wait(espera)}
		else {
			signal(preceptor)
			wait(espera)
		}
		copia_parcial = parcial;
	}

	Procedure preceptorEntrega(){
		if  (esperando < 45) { wait(preceptor) }
		signal_all(espera)
	}

	Procedure esperarNota(idAlumno: in int,ParcialResuelto: in txt,nota: out int){
		contadorParciales++;
		colaParciales.push(idAlumno,ParcialResuelto);
		signal(profe);
		wait(alumnosEsperan[idAlumno]);
		nota = notasAlumnos[idAlumno];
	}

	Procedure esperarParcial(parcial: out txt,idAlumno: out int){
		if(contadorParciales == 0){ wait(profe)}
		contadorParciales--;
		colaParciales.pop(parcialMandado,idAlumno);

	}

	Procedure darNota(idAlumno: in int,nota: in int){
		notasAlumnos[idAlumno] = nota;
		signal(alumnosEsperan[idAlumno]);
	}

}




Process Alumno[id:0..44]{
	txt copia_parcial; int nota;
	Parcial.llegada(copia_parcial);
	//hace parcial
	Parcial.esperarNota(id,copia_parcial,nota);

}

Process preceptor{
	Parcial.preceptorEntrega();
}

Process Profesora{
	txt parcial, int idAlumnol,int nota;
	for i:0..44{
		Parcial.esperarParcial(parcial,idAlumno);
		nota = corregirExamen(parcial);
		Parcial.darNota(idAlumno,nota)
	}
}