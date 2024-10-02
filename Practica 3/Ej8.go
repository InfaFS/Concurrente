Process Jugador[id:0..19]{
    int nroEquipo;
    int canchaId;

    Jugar.elegirEquipo(nroEquipo);
    equipo[nroEquipo].listo(canchaId);
    cancha[canchaId].llegada();
}

Process Partido[id: 0..1] {
    Cancha[id].iniciar();
    //Delay 50
    Cancha[id].terminar();
}

Monitor Jugar {

    Procedure elegirEquipo(nro: out int){
        nro = DarEquipo();
    }
}

Monitor Equipo[id: 1..4] {
    int equipoCant = 0;
    cond espera;
    int cancha;

    Procedure Listo (canchaId: out int) {
        equipoCant++;
        if (equipoCant < 5){
            wait(espera);
        } 
        else {
            Administrador.calcularCancha(cancha);
            signal_all(espera);
        }
        canchaId = cancha;
    }
}

Monitor Administrador {
    int totalFormados = 0;
    
    Procedure calcularCancha(canchaId: out int) {
        totalFormados++;
        if (totalFormados <= 2){
            canchaId = 1;
        }
        else {
            canchaId = 2;
        }
    }
}

Monitor Cancha[id: 0..1] {
    int cant = 0;
    cond espera;
    cond inicio;

    Procedure llegada() {
        cant++;
        if (cant == 10){
            signal(inicio)
        }
        wait(espera);
    }

    Procedure iniciar(){
        if (cant < 10){
            wait(inicio);
        }
    }

    Procedure terminar(){
        signal_all(espera);
    }

} 