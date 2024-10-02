
//a

Monitor Corralon {
	bool libre = true;
	cond empleado;
	cond cola;

	Cola colaProductos;
	Cola colaComprobantes;
	Procedure llegada(comprobante: out txt,productos: in txt){
		colaProductos.push(productos)

		signal(empleado)
		wait(cola)
		comprobante = colaComprobantes.pop()
	}

	Procedure facturar(){
		if (colaProductos.empty()){
			wait(empleado)
		}

		productos = colaProductos.pop();
		comprobante = facturar(productos);
		colaComprobantes.push(comprobante);
		signal(cola)
	}

}


Process Cliente[id:0..N-1]{
	Corralon.llegada(comprobante,productos)
}

Process Empleado{
	for i:0..N-1 {
		Corralon.facturar()
	}
}

//a - 2

Monitor Corralon {
	bool libre = true;
	cond empleado;
	cond cola;
	cond recibirComprobante;
	cond recibirProductos;
	cond meVoy;

	Cola colaProductos;
	Cola colaComprobantes;

	int esperando = 0;

	Procedure llegada(comprobante: out txt,productos: in txt){
		esperando ++; 
		signal(empleado)
		wait(cola)
		colaProductos.push(productos)
		//deberia hacer 2 esperando, uno para las personas que esperan el comprobante y otra para las que esperan que les agarren los productos?
		signal(recibirProductos)
		wait(recibirComprobante)
		comprobante = colaComprobantes.pop()
		esperando --;
		signal(meVoy)
	}

	Procedure obtenerProductos(productos: out txt){
		if (esperando == 0){
			wait(empleado)
		}
		signal(cola)
		wait(recibirProductos)
		productos = colaProductos.pop();
	}

	Procedure darComprobante(comprobante: in txt){
		colaComprobantes.push(comprobante)
		signal(recibirComprobante)
		wait(meVoy)
	}

}




Process Cliente[id:0..N-1]{
	Corralon.llegada(comprobante,productos)
}

Process Empleado{
	for i:0..N-1 {
		Corralon.obtenerProductos(productos)
		//generar comprobante
		Corralon.darComprobante(comprobante)
	}
}

//b




Monitor Corralon{
    int cantLibres = 0;
    int esperando = 0; 
    cola eLibres;
    cond esperaC;

    Procedure llegada (idE: out int){
        if (cantLibres == 0){
            esperando++;
            wait(esperaC);
        }
        else {
            cantLibres--;
        }
        idE = eLibres.pop();
    }

    Procedure proximo(idE: in int){
        eLibres.push(idE);
        if (esperando > 0){
            esperando--;
            signal(esperaC);
        }
        else {
            cantLibres++;
        }
    }
}

Monitor Escritorio[id:0..E-1]{
    text listaC;
    text compE;
    bool hayDatos = false;
    cond datos;
    cond atencionE;

    Procedure atencion(list: in text; comp: out text){
        listaC = list;
        hayDatos = true;
        signal(datos);
        wait(atencionE);
        comp = compE;
        signal(datos);
    }

    Procedure obtenerLista(list: out text){
        if (!hayDatos){
            wait(datos);
        }
        list = listaC;
    }

    Procedure dejarComprobante(comp: in text){
        compE = comp;
        signal(atencionE);
        wait(datos);
        hayDatos = false;
    }

}


Process Cliente[id:0..N-1]{
    int idE;
    text comprobante;
    text lista;
    
    Corralon.llegada(idE);
    Escritorio[idE].atencion(lista,comprobante);
}

Process Empleado[id:0..E-1]{
    
    while(true){
        Corralon.proximo(id);
        Escritorio[id].obtenerLista(lista);
        comprobante = comprobar(lista);
        Escritorio[id].dejarComprobante(comprobante);
    }
}