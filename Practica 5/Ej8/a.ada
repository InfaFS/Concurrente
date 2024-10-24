--NO sabria como hacer para que hayan entries personalizadas dentro del admin ,PREGUNTAR
procedure Ej8 is

    --Workers le avisan que tienen que realizar un calculo
    Task type Cliente is
        Entry Atendido();
        Entry RecibirID(IDCliente: IN integer);
    End Cliente;

    --Definicion del admin que avisa
    Task Admin is
        Entry PedidoCliente();
        Entry PedidoCamion(IDCliente: OUT integer);
        Entry LlegadaCamion(IDCliente: IN integer);
    End Admin;

    Task Type Camion;
 

    arrCamiones: array(1..3) of Camion;

    arrCliente: array(1..P) of Cliente;

    
    --Clientes
    Task Body Cliente is
        MID: integer := 0;
    begin
        --voy a recibir el id del main
        accept RecibirID(IDCliente: IN Integer) do
            MID:=IDCliente;
        end RecibirID;

        
        while not atendido loop
        admin.hacerReclamo(MID);
            SELECT
                accept Atendido() do
                    atendido := true;
                end Atendido;
            OR delay 900.0;
                null;
            END SELECT
        end loop

    end Cliente;

    --Admin
    Task Body Admin is
    arrReclamos: array (1..P) of integer;
    begin
        loop
            SELECT
                when arrReclamos.notEmpty() =>
                    accept PedidoCamion(IDCliente : OUT integer) do
                        --se fija cual es el cliente que mas pedidos tiene, lo haria con un count de entries pero no se como hacerlo en este caso
                        IDCliente_Temp := max(arrReclamos);
                        arrReclamos[IDCliente_Temp] := 0;
                        IDCliente := IDCliente_Temp;
                    end PedidoCamion;
            OR
                accept hacerReclamo(IDCliente : IN integer) do
                    arrReclamos[IDCliente]++;
                end hacerReclamo;
            END SELECT;
        end loop 
      
    End Admin;

    Task Body Camion is

    begin
        loop
        admin.PedidoCamion(IDCliente);
        --voy hasta el lugar
        cliente[IDCliente].Atendido();
        end loop;
    end Camion;

begin
    for I in 1..P loop
        arrCliente.RecibirID(I);
    end loop;
end Ej8;

