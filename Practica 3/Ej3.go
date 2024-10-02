//a)

Monitor Acceso{
  cond espera;
  usando = 0;

  Procedure acceder () {
	while (usando > 0 ) { wait (espera); }
	usando ++;
  }

  Procedure terminar () {
	usando --;
	signal(espera)
  }

}

Process persona[id:1..N]{
	Documento documento;
	Fotocopia fotocopia;

	Acceso.acceder();
	fotocopia = Fotocopiar(documento);
	Acceso.terminar()
}


//a-alterno) //--> Checkquear

Monitor Acceso{
  
	Procedure acceder (documento: in text,fotocopia out text) { //Las variables que me ingresan tienen que ser de algun tipo? onda text y eso
	  fotocopia = Fotocopiar(documento); // --> estaria bien o generaria bussy waiting?
	}
}	


  Process persona[id:1..N]{
	Documento documento;
	Fotocopia fotocopia;
	Acceso.acceder(documento,fotocopia);
}


//b)



Monitor Acceso{
	cond cv;
	esperando = 0;
	libre = true;
  
	Procedure acceder () {
	  if (not libre){
			esperando++;
			wait(cv)
	  } else {
			libre = false;
		}
	}
  
	Procedure terminar () {
		if (esperando > 0){
			esperando--;
			signal(cv);
		} else {
			libre = true;
		}
	}
  
  }
  
Process persona[id:1..N]{
	Documento documento;
	Fotocopia fotocopia;

	Acceso.acceder();
	fotocopia = Fotocopiar(documento);
	Acceso.terminar()
}


//c)

Monitor Acceso{
	cond cv[N];
	idAux,esperando = 0;
	libre = true;
	colaOrdenada fila;

  
	Procedure acceder (id: in int,edad: in int) {
	  if (not libre){
			insertar(fila,id,edad);
			esperando++;
			wait(cv[id])
	  } else {
			libre = false;
		}
	}
  
	Procedure terminar () {
		if (esperando > 0){
			esperando--;
			sacar(fila,idAux);
			signal(cv[idAux]);
		} else {
			libre = true;
		}
	}
  
  }
  
Process persona[id:1..N]{
	Documento documento;
	Fotocopia fotocopia;

	Acceso.acceder(id,edad);
	fotocopia = Fotocopiar(documento);
	Acceso.terminar()
}

//d)

Monitor Acceso{
  cond cv[N];
	proximo = 0;

  Procedure acceder (id: in int) {
		if (id != proximo) { wait(cv[id]) }
  }

  Procedure terminar () {
		proximo++;
		signal(cv[proximo])
  }

}

Process persona[id:1..N]{
	Documento documento;
	Fotocopia fotocopia;

	Acceso.acceder(id);
	fotocopia = Fotocopiar(documento);
	Acceso.terminar()
}

//e)



Monitor Fotocopiadora {
	int esperando = 0;
	cond empleado;
	cond persona;
	cond fin;

	Procedure usar(){
		esperando++; //primero no iria el esperando para que no cause problemas en asginar?
		signal(empleado);
		wait(persona);
	}

	Procedure dejar(){
		signal(fin);
	}

	Procedure asignar(){
		if (esperando == 0){
			wait (empleado);
		}
		esperando--;
		signal(persona);
		wait(fin);
	}
}

Process Persona [id:0..N-1]{
	Documento documento; 
	Fotocopia fotocopia;

	Fotocopiadora.usar()
	Fotocopia = scan(documento);
	Fotocopiadora.dejar();
}

Process Empleado{
	int i;
	for i = 0..N-1 {
		Fotocopiadora.asignar()
	}
}

//f

Process Persona [id:0..N-1]{
	Documento documento; 
	Fotocopia fotocopia;
	Fotocopiadora fotocopiadoraAsig;

	Fotocopiadora.usar(fotocopiadoraAsig);
	Fotocopia = fotocopiadoraAsig.fotocopiar(documento);
	Fotocopiadora.dejar(fotocopiadoraAsig);
}

Process Empleado{
	int i;
	for i = 0..N-1 {
			Fotocopiadora.asignar()
	}
}

Monitor Fotocopiadora {
	cola cola;
	cola fotocopiadoras = {1,2,3,4,5,6,7,8,9,10};
	cond empleado;
	cond persona[n] = ([n] = 0);
	cond fotocop;
	int asignada[n] = ([n] = 0);

	Procedure usar(fotocopiadoraAsig: out int, idP: in int){
		 signal(empleado); // --> deberia ir despues?
		 cola.push(idP);
		 wait(persona[idP]);
		 fotocopiadoraAsig = asignada[idP];
	}

	Procedure dejar(fotocopiadoraAsig: in int){
			fotocopiadoras.push(fotocopiadoraAsig);
			signal(fotocop);
	}

	Procesudre asignar(){
			if (cola.empty()){
					wait(empleado);
			}
			idAux = cola.pop();
			if (fotocopiadoras.empty()){
					wait(fotocop);
			}
			asignada[idAux] = fotocopiadoras.pop();
			signal(persona[idAux]);

	}

}