
Monitor Equipo[id:0..4]{
	int equipoArribados = 0;
	int totalVendido = 0;
	cond espera;

	Procedure listo(){
		equipoArribados ++;
		if (equipoArribados < 4){
			wait(espera)
		} else {
			signal_all(espera)
			equipoArribados = 0
		}
	}

	Procedure terminar(sumaVendido: in int,vendido: out int){
		equipoArribados++;
		totalVendido += sumaVendido; //hay que poner los valores que se van sumando a medida que llegan, sino va a devolver 0
		if (equipoArribados < 4){
			wait(espera)
		} else {
			signal_all(espera)
		}
		vendido = totalVendido;
	}

}

Process vendedor[id:0..19]{
	int id_grupo,vendido;
	Equipo[id_grupo].listo()
	//venden
	Equipo[id].terminar(vendido)
}