Process Administrativos[id:0..N-1]{
	while(true){
		//trabajan
		send enviarImpresora (documento)
	}

}

Process Impresoras[id:0..2]{
	while(true){
		receive enviarImpresora(documento)
		imprimir(documento)
	}
}


//b

Process Administrativos[id:0..N-1]{
	while(true){
		//trabajan
		send pedidoAdministrativo(id)
		send hayPedido()
		receive	idImpresora[id] (idImp) 
		send enviarImpresora[idImpresora] (documento)
	}

}

Process Director{
	while(true){
		//trabaja
		send pedidoDirector()
		send hayPedido()
		receive idImpresoraDirector (idIm)
		send enviarImpresora[idIm] (documento)
	}
}

Process Impresoras[id:0..2]{
	while(true){
		receive enviarImpresora[id] (documento)
		imprimir(documento)
	}
}

Process Administrador{
	while(true){
		receive hayPedido()
		if (!empty(pedidoAdministrativo) and empty(pedidoDirector)){
			receive pedidoAdministrativo(idAdmin)
			elegirImpresora(idImp)
			send idImpresora[idAdmin] (idImp)
		} else if (!empty(pedidoDirector)){
			receive pedidoDirector()
			elegirImpresora(idImp)
			send idImpresoraDirector (idImp)
		}

	}
}


// version agus 
chan imprimirDocU(text);
chan imprimirDocD(text);
chan hayDoc(bool);


Process Usuario [id: 0..N-1] {
    text doc;

    while (true) {
        doc = generarDoc();
        send imprimirDocU(doc);
        send hayDoc(true);
    }
}

Process Director {
    text doc;

    while (true) {
        doc = generarDoc();
        send imprimirDocD(doc);
        send hayDoc(true);
    }
}

Process Impresora [id: 0..2] {
    text doc;
    bool hay;

    while (true) {
        receive hayDoc(hay);

        if (!empty (imprimirDocD)) {
            receive imprimirDocD(doc);
        } else {
            receive imprimirDocU(doc);
        }
        imprimir(doc);
    }
}