

Monitor Carrera{
	int contador = 0;
	cond esperaCorredores;
	int botellas_actuales = 20;
	bool libre = true;
	Cola ColaBotellas;
	cond Repositor;
	cond BotellasRepuestas;

	Procedure Iniciar(){
		contador++;
		if (contador < C){
			wait(esperaCorredores);
		} else {
			contador = 0;
			signal_all(esperaCorredores);
		}
	
	}

	Procedure TerminarCarrera(){
		if (not libre){
			contador++;
			wait(esperaCorredores);
		} else {
			libre = false;
		}

	}



	Procedure SacarBotella(botella: out int){
		if (botellas_actuales == 0){
			signal(Repositor);
			await(BotellasRepuestas);
		}
		botellas_actuales --;
		botella = ColaBotellas.pop();

		if(contador > 0){
			contador--;
			signal(esperaCorredores);

		} else {
			libre = true;
		}

	}

	Procedure Reponer(){
		if (botellas_actuales > 0){
			wait(Repositor);
		}

		for i:0..19{
			ColaBotellas.push(botella)
		}
		botellas_actuales = 20;

		signal(BotellasRepuestas);

	}

}



Process corredor[id:0..C-1]{
	Carrera.Iniciar()
	//corre...
	Carrera.TerminarCarrera();
	Maquina.SacarBotella(botella);
	Carrera.dejarMaquina();
}

Process repositor(){
	while(true){
		Maquina.Reponer();
	}
}


//alterno



Monitor Carrera{
	int contador = 0;
	cond esperaCorredores;
	bool libre = true;


	Procedure Iniciar(){
		contador++;
		if (contador < C){
			wait(esperaCorredores);
		} else {
			contador = 0;
			signal_all(esperaCorredores);
		}
	
	}

	Procedure TerminarCarrera(){
		if (not libre){
			contador++;
			wait(esperaCorredores);
		} else {
			libre = false;
		}

	}


	Procedure dejarMaquina(){
		if(contador > 0){
			contador--;
			signal(esperaCorredores);

		} else {
			libre = true;
		}
	}

}

Monitor Maquina{
	Cola ColaBotellas;
	cond Repositor;
	cond BotellasRepuestas;
	int botellas_actuales = 20;

	Procedure SacarBotella(botella: out int){
		if (botellas_actuales == 0){
			signal(Repositor);
			await(BotellasRepuestas);
		}
		botellas_actuales --;
		botella = ColaBotellas.pop();

	}

	Procedure Reponer(){
		if (botellas_actuales > 0){
			wait(Repositor);
		}

		for i:0..19{
			ColaBotellas.push(botella)
		}
		botellas_actuales = 20;

		signal(BotellasRepuestas);

	}

}

	





Process corredor[id:0..C-1]{
	Carrera.Iniciar()
	//corre...
	Carrera.TerminarCarrera();
	Carrera.SacarBotella(botella);
}

Process repositor(){
	while(true){
		Carrera.Reponer();
	}
}


