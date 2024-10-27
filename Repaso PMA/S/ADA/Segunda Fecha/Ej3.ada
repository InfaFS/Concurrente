Procedure Ej3 is

--Task de admin que va recibiendo los pedidos y se lo da al server
--El server que va resolviendo los pedidos
--Los clientes que van mandando

Task type Cliente is
    Entry RecibirResultado (Respuesta: IN texto);
End Cliente;

Task type Admin is
    Entry RecibirPedidoServer(Cadena: OUT texto);
    Entry RecibirPedidoCliente(Cadena: IN texto);
End Admin;


Task type Server;


arrClientes: array (1..P) of Cliente;
arrServers: array (1..3) of Server;

Task Body Admin is
    colaCadenas: Cola;
Begin
    loop
        SELECT
            accept RecibirPedidoCliente(Cadena: IN texto,ID: IN int) do
                colaCadenas.push(Cadena,ID);
            end RecibirPedidoCliente;
        OR
            when (!empty(colaCadenas))=>
                accept RecibirPedidoServer(Cadena: OUT texto,ID: OUT int) do
                    Cadena,ID := colaCadenas.pop();
                end RecibirPedidoServer;
        END SELECT;
    end loop;


End Admin;


Task body Cliente is
    Cadena: texto;
    IDC: int;
begin
    accept recibirId(ID: IN int) do
        IDC:= ID;
    end recibirId;

    loop
        Cadena:= generarCadena();
        Admin.RecibirPedidoCliente(Cadena);
        accept RecibirResultado(Respuesta: IN texto) do
            --hace algo con el resultado
        end RecibirResultado;

    end loop;
end Cliente;


Task body Server is
    Cadena,Respuesta: texto;
    ID: int;
begin
    loop
        Admin.RecibirPedidoServer(Cadena,ID);
        Respuesta := ResolverAnalsis(Cadena);
        arrClientes[ID].RecibirResultado(Respuesta);
    end loop;
end Server;


begin
    loop for I = 0..P-1 do
        arrClientes[I].recibirId(I);
    end loop;
end Ej3;
