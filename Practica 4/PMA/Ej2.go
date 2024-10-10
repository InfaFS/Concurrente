chan pedido[5](text, int);
chan comprobante[P](text);
chan buscarCaja(int);
chan obtenerCaja[P](int);
chan liberarCaja(int);
chan hayPedido(bool);

/*
Parece ser que agus usa un canal pedidos para avisar tanto si alguien quiere liberar la caja como si quieren pedir una caja
Intuyo que se haria para poder generar mayor concurrencia, ya que si usase solo un receive de los canales correspondientes
a buscar caja y liberar caja, se trabaria en uno de esos canales y no liberaria o asignaria una caja hasta recibir algo por el
receive
*/

Process Caja [id: 0..4] {
    int idAux;
    text pago;
    text comprobante;
    while(true) {
        receive pedido[id](pago, idAux); //la caja solo se encarga de recibir pedidos en el canal pedido
        generarComprobante(pago, comprobante); //genera el comprobante
        send comprobante[idAux](comprobante); //envia el comprobante al channel privado de cada persona
    }
}

Process Cliente [id: 0..P-1] {
    int idCaja;
    text pago;
    text comprobante;

    send buscarCaja(id); //enviar en el channel buscar caja su id
    send hayPedido(true); //Maneja el tema de pedir buscar caja con otro channel de hay pedidosj
    receive obtenerCaja[id](idCaja); //Obtiene la caja en su vector privado de cada persona
    send pedido[idCaja](pago, id); //Le envia a la caja el pedido y se van encolando
    receive comprobante[id](comprobante); //recibe el comprobante en su channel privado
    send liberarCaja(idCaja); //encola el id de la caja que hay que restarle gente esperando
    send hayPedido(true); //le avisa al admin que hay un pedido

}


Process Admin {
    int cantEspera[5] = ([5] 0);
    int min;
    int idAux;
    bool pedido;

    while (true) {
        receive hayPedido(pedido); //el admin recibe un pedido
        if (!empty (buscarCaja) && empty(liberarCaja)) { //pregunta si hay alguien queriendo buscar caja y no hay nadie que quiera liberar la caja
            receive buscarCaja(idAux); //obtiene lo que esta en el channel buscarCaja
            min = cajaMasVacia(cantEspera); //obtiene el minimo del vector de ints
            cantEspera[min]++; //suma la cantida de espera en esa caja
            send obtenerCaja[idAux](min); //le manda a la persona l caja que tiene que ir
        } else {
            if (!empty (liberarCaja)) { //si hay alguien que quiere liberar la caja, entonces lo recibe y resta la cantidad de espera
                receive liberarCaja(idAux);
                cantEspera[idAux]--;
            }
        }
    }
}